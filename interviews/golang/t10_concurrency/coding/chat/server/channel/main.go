package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
)

func main() {
	cs := NewChatServer(":8085")

	go func() {
		if err := cs.Run(); err != nil {
			fmt.Println("chat server run:", err)
		}
		fmt.Println("chat server stopped")
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	<-ctx.Done()
	cs.Shutdown()
}

type client struct {
	addr  string
	write chan string
	read  chan string
}

type message struct {
	sender string
	data   string
}

type ChatServer struct {
	address  string
	listener net.Listener

	connect    chan client
	disconnect chan client

	quit          chan struct{}
	quitBroadcast chan struct{}
	runDone       chan struct{}
}

func NewChatServer(address string) *ChatServer {
	return &ChatServer{
		address:       address,
		quit:          make(chan struct{}),
		connect:       make(chan client),
		disconnect:    make(chan client),
		quitBroadcast: make(chan struct{}),
		runDone:       make(chan struct{}),
	}
}

func (cs *ChatServer) Run() error {
	defer close(cs.runDone)

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	cs.listener = lis
	fmt.Println("server listening on", cs.listener.Addr())

	var wg sync.WaitGroup

	go cs.broadcast()

	for {
		var rawConn net.Conn
		rawConn, err = cs.listener.Accept()
		if err != nil {
			// listener closed
			break
		}

		conn := Connection{Conn: rawConn}
		cl := client{
			addr:  conn.RemoteAddr().String(),
			write: make(chan string),
			read:  make(chan string),
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			conn.writeLoop(&cl)
			_ = conn.Close()
			cs.disconnect <- cl
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			conn.readLoop(&cl)
			_ = conn.Close()
			cs.disconnect <- cl
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			<-cs.quit
			_ = conn.Close()
			cs.disconnect <- cl
		}()

		cs.connect <- cl
	}

	wg.Wait()

	close(cs.quitBroadcast)

	return nil
}

func (cs *ChatServer) broadcast() {
	defer fmt.Println("broadcast stopped")

	clients := make(map[string]client)
	messages := make(chan message)

	broadcastMessage := func(msg message) {
		sendMsg := fmt.Sprintf(
			"%s: %s", msg.sender, msg.data,
		)
		fmt.Println("broadcast message:", sendMsg)
		for _, cl := range clients {
			if cl.addr == msg.sender {
				continue
			}
			cl.write <- sendMsg
		}
	}

	for {
		select {
		case <-cs.quitBroadcast:
			return
		case cl := <-cs.connect:
			clients[cl.addr] = cl
			go func(cl *client) {
				for {
					msg, ok := <-cl.read
					if !ok {
						return
					}
					messages <- message{
						sender: cl.addr,
						data:   msg,
					}
				}
			}(&cl)
			fmt.Println(cl.addr, "client connected")
			cl.write <- "you are " + cl.addr
			broadcastMessage(message{
				sender: cl.addr,
				data:   "connected",
			})
		case cl := <-cs.disconnect:
			if _, ok := clients[cl.addr]; ok {
				delete(clients, cl.addr)
				close(cl.write)
				fmt.Println(cl.addr, "client disconnected")
			}
		case msg := <-messages:
			broadcastMessage(msg)
		}
	}
}

func (cs *ChatServer) Shutdown() {
	fmt.Println("server shutdown started")
	defer fmt.Println("server shutdown finished")
	if err := cs.listener.Close(); err != nil {
		fmt.Println("listener close:", err)
	}
	fmt.Println("listener closed")
	close(cs.quit)
	<-cs.runDone
	fmt.Println("run done")
}

type Connection struct {
	net.Conn
	closeOnce sync.Once
}

func (c *Connection) readLoop(cl *client) {
	fmt.Println(cl.addr, "readLoop started")

	defer func() {
		close(cl.read)
		fmt.Println(cl.addr, "readLoop stopped")
	}()

	for {
		buf := make([]byte, 100)

		n, err := c.Read(buf)
		if err != nil {
			// connection closed
			return
		}
		if n == 0 {
			fmt.Println(cl.addr, "zero bytes read")
			return
		}

		cl.read <- string(buf[:n])
	}
}

func (c *Connection) writeLoop(cl *client) {
	fmt.Println(cl.addr, "writeLoop started")
	defer fmt.Println(cl.addr, "writeLoop stopped")

	for {
		msg, ok := <-cl.write
		if !ok {
			return
		}
		n, err := fmt.Fprintln(c, msg)
		if err != nil {
			return
		}
		if n == 0 {
			fmt.Println(cl.addr, "zero bytes written")
			return
		}
	}
}

func (c *Connection) Close() error {
	c.closeOnce.Do(func() {
		if err := c.Conn.Close(); err != nil {
			fmt.Println(c.RemoteAddr(), "connection close:", err)
		}
		fmt.Println(c.RemoteAddr(), "connection closed")
	})

	return nil
}

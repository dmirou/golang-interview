package main

import (
	"context"
	"fmt"
	"io"
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

type ChatServer struct {
	address  string
	listener net.Listener

	ctx    context.Context
	cancel context.CancelFunc

	clients sync.Map
	runDone chan struct{}
}

func NewChatServer(address string) *ChatServer {
	ctx, cancel := context.WithCancel(context.Background())

	return &ChatServer{
		ctx:     ctx,
		cancel:  cancel,
		address: address,
		runDone: make(chan struct{}),
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

	defer wg.Wait()

	for {
		var rawConn net.Conn
		rawConn, err = cs.listener.Accept()
		if err != nil {
			// listener closed
			return nil
		}

		conn := Connection{Conn: rawConn}

		wg.Add(1)
		go func() {
			defer wg.Done()
			cs.handleConn(&conn)
			_ = conn.Close()
			cs.clients.Delete(conn.RemoteAddr())
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			<-cs.ctx.Done()
			_ = conn.Close()
		}()

		cs.clients.Store(conn.RemoteAddr(), &conn)
	}
}

func (cs *ChatServer) handleConn(conn *Connection) {
	fmt.Println(conn.RemoteAddr(), "handleConn started")

	for {
		buf := make([]byte, 100)

		n, err := conn.Read(buf)
		if err != nil {
			// connection closed
			return
		}
		if n == 0 {
			fmt.Println(conn.RemoteAddr(), "client disconnected")
			return
		}
		cs.broadcast(conn.RemoteAddr(), buf[:n])
	}
}

func (cs *ChatServer) broadcast(sender net.Addr, message []byte) {
	msg := fmt.Sprintf("%s: %s\n", sender.String(), string(message))
	fmt.Println("broadcast message: ", msg)

	cs.clients.Range(func(addr, conn any) bool {
		if sender == addr {
			return true
		}
		if w, ok := conn.(io.Writer); ok {
			_, _ = w.Write([]byte(msg))
		}
		return true
	})
}

func (cs *ChatServer) Shutdown() {
	fmt.Println("server shutdown started")
	defer fmt.Println("server shutdown finished")
	if err := cs.listener.Close(); err != nil {
		fmt.Println("listener close:", err)
	}
	fmt.Println("listener closed")
	cs.cancel()
	<-cs.runDone
	fmt.Println("all connections closed")
}

type Connection struct {
	net.Conn
	closeOnce sync.Once
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

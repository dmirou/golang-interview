// clockwall connects to a few time servers and prints received time
// Run example:
// go run main.go localhost:8080 localhost:8081 localhost:8082
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type ClientBuffer struct {
	addr string
	buf  []byte
}

func NewClientBuffer(addr string) *ClientBuffer {
	return &ClientBuffer{
		addr: addr,
		buf:  make([]byte, 25),
	}
}

var cbs []*ClientBuffer

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("connect addresses are not specified: %v\n", os.Args)
	}

	cbs = make([]*ClientBuffer, len(os.Args)-1)

	var done = make(chan struct{})
	for i, addr := range os.Args[1:] {
		cbs[i] = NewClientBuffer(addr)

		go func(addr string, done chan struct{}, buf []byte) {
			readConn(addr, buf)
			done <- struct{}{}
		}(addr, done, cbs[i].buf)
	}

	go showTime()

	for range os.Args[1:] {
		<-done
	}

	fmt.Println("main: done")
}

func readConn(addr string, buf []byte) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("can't connect to %s: %v\n", addr, err)
		return
	}

	defer func(addr string, conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("can't close connection %s: %v\n", addr, err)
		}
	}(addr, conn)

	for {
		_, err = conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("can't read from %s: %v", addr, err)
			return
		}
	}
}

func showTime() {
	var b strings.Builder
	for {
		b.Reset()
		time.Sleep(500 * time.Millisecond)

		for _, cb := range cbs {
			b.WriteString(cb.addr + "->")
			b.Write(cb.buf)
			b.WriteString(" ")
		}
		b.WriteString("\n")
		fmt.Print(b.String())
	}
}

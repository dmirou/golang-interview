package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("connect addresses are not specified: %v\n", os.Args)
	}

	var done = make(chan struct{})

	for _, addr := range os.Args[1:] {
		go func(done chan struct{}) {
			readConn(addr)
			done <- struct{}{}
		}(done)
	}

	for range os.Args[1:] {
		<-done
	}

	fmt.Println("main: done")
}

func readConn(addr string) {
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

	var b = make([]byte, 50)

	for {
		_, err := conn.Read(b)
		if err != nil {
			fmt.Printf("can't read from %s: %v\n", addr, err)
			break
		}
		fmt.Printf("from %s: %s", addr, b)
		time.Sleep(100 * time.Millisecond)
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	const port = 8181

	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("clock1: can't start listener at %s: %v", addr, err)
	}
	defer listener.Close()

	fmt.Println("clock1: wait for connections")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("clock1: can't accept connection: %v", err)
			return
		}
		go handleCon(conn)
	}

	fmt.Println("clock1: done")
}

func handleCon(conn net.Conn) {
	fmt.Println("handleCon: client connected")

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("handleCon: can't close connection: %v\n", conn)
		}
		fmt.Println("handleCon: done")
	}(conn)

	for {
		msg := time.Now().Format("15:04:05\n")
		_, err := io.WriteString(conn, msg)
		if err != nil {
			fmt.Printf("handleCon: can't send to the client: %v", err)
			return
		}
		fmt.Printf("handleCon: sent to the client: %s", msg)

		time.Sleep(1 * time.Second)
	}
}

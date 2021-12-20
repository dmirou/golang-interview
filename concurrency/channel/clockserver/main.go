package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("can't get port number from args: %v", os.Args)
	}

	port, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		log.Fatalf("can't parse input port %s: %v", os.Args[1], err)
	}

	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("clock1: can't start listener at %s: %v", addr, err)
	}
	defer listener.Close()

	fmt.Printf("clock1: wait for connections on %s\n", listener.Addr().String())

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
			fmt.Printf("handleCon: can't send to the client: %v\n", err)
			return
		}
		fmt.Printf("handleCon: sent to the client: %s", msg)

		time.Sleep(1 * time.Second)
	}
}

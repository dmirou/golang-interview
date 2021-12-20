// main sends current time to each client every second through tcp connection.
// To start the server specify timezone and port.
// TIMEZONE=Local go run main.go 8081
// TIMEZONE=UTC go run main.go 8082
// TIMEZONE=Europe/Moscow go run main.go 8081
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

	tz := os.Getenv("TIMEZONE")
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatalf("can't load location for TIMEZONE env var %s: %v", tz, err)
	}

	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("clockserver: can't start listener at %s: %v", addr, err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			fmt.Printf("clockserver: can't close listener: %v\n", err)
		}
		fmt.Println("clockserver: done")
	}()

	fmt.Printf("clockserver: wait for connections on %s\n", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("clockserver: can't accept connection: %v", err)
			return
		}
		go handleCon(conn, loc)
	}
}

func handleCon(conn net.Conn, location *time.Location) {
	fmt.Println("handleCon: client connected")

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("handleCon: can't close connection: %v\n", conn)
		}
		fmt.Println("handleCon: done")
	}(conn)

	for {
		msg := time.Now().In(location).Format("15:04:05")
		_, err := io.WriteString(conn, msg)
		if err != nil {
			fmt.Printf("handleCon: can't send to the client: %v\n", err)
			return
		}
		fmt.Printf("handleCon: sent to the client: %s\n", msg)

		time.Sleep(1 * time.Second)
	}
}

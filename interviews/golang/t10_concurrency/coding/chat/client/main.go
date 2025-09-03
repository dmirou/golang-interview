package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", ":8085")
	if err != nil {
		log.Fatal("net dial:", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		sc := bufio.NewScanner(conn)

		for sc.Scan() {
			fmt.Println("got:", sc.Text())
		}
	}()

	go func() {
		defer wg.Done()
		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {
			if sc.Text() == "" {
				continue
			}
			_, err = conn.Write(sc.Bytes())
			if err != nil {
				fmt.Println("conn write:", err)
				return
			}
			fmt.Println("sent:", sc.Text())
		}
	}()

	wg.Wait()
}

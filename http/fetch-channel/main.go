// Reads urls from command line args and fetch them. Use output channel.
// example:
// go run main.go https://ya.ru https://google.com t2 https://gmail.com https://vk.com https://github.com https://test.com
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := os.Args[1:]

	var ch = make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Printf("%s\n", <-ch)
	}
}

func fetch(url string, ch chan<- string) {
	client := &http.Client{Timeout: 2 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Errorf("can not fetch url %s: %v\n", url, err).Error()
		return
	}
	defer resp.Body.Close()

	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Errorf("can not read response body: %v\n", err).Error()
		return
	}
	ch <- fmt.Sprintf("read %d bytes from %s", n, url)
}

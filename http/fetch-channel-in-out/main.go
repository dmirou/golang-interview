// Reads urls from command line args and fetch them. Use input and output channel and wait group.
// example:
// go run main.go https://ya.ru https://google.com t2 https://gmail.com https://vk.com https://github.com https://test.com
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	urls := os.Args[1:]

	var chIn = make(chan string)

	go func(chIn chan string, urls []string) {
		for _, url := range urls {
			chIn <- url
		}
		close(chIn)
	}(chIn, urls)

	chOut := make(chan string)

	go runFetching(chIn, chOut)

	for msg := range chOut {
		fmt.Printf("%s\n", msg)
	}
}

func runFetching(ch <-chan string, chOut chan<- string) {
	var wg sync.WaitGroup

	for url := range ch {
		wg.Add(1)
		go func(url string) {
			fetch(url, chOut)
			wg.Done()
		}(url)
	}

	wg.Wait()

	close(chOut)
}

func fetch(url string, ch chan<- string) {
	fmt.Printf("parsing url started: %s\n", url)
	client := &http.Client{Timeout: 2 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Errorf("can not fetch url %s: %v", url, err).Error()
		return
	}
	defer func(resp *http.Response) {
		_ = resp.Body.Close()
	}(resp)

	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Errorf("can not read response body: %v", err).Error()
		return
	}
	ch <- fmt.Sprintf("read %d bytes from %s", n, url)
}

// Reads urls from command line args and fetch them. Use sync.WaitGroup.
// example:
// go run main.go https://ya.ru https://google.com t2 https://gmail.com https://vk.com https://github.com https://test.com
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	urls := os.Args[1:]

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			n, err := fetch(url)
			if err != nil {
				log.Printf("can not fetch url %s: %v", url, err)
				return
			}

			log.Printf("read %d bytes from url %s", n, url)
		}(url)
	}

	wg.Wait()
}

func fetch(url string) (int64, error) {
	client := &http.Client{Timeout: 2 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("can not fetch url %s: %v\n", url, err)
	}
	defer resp.Body.Close()

	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return 0, fmt.Errorf("can not read response body: %v\n", err)
	}

	return n, nil
}

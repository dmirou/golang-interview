package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	urls := make([]string, 1000000) // тут ссылки

	workersCount := 10
	urlsChan := make(chan string, workersCount)

	go func() {
		defer fmt.Println("urls chan closed")
		for i, _ := range urls {
			urlsChan <- strconv.Itoa(i)
		}
		close(urlsChan)
	}()

	codesChan := make(chan int, workersCount)
	var wg sync.WaitGroup
	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func() {
			defer wg.Done()
			for url := range urlsChan {
				codesChan <- sendRequest(url)
			}
		}()
	}

	go func() {
		wg.Wait()
		fmt.Println("workers finished")
		close(codesChan)
	}()

	respCodes := make(map[int]int, 3)
	for code := range codesChan {
		respCodes[code]++
	}

	for code, count := range respCodes {
		fmt.Printf("%d: %d\n", code, count)
	}
}

func sendRequest(url string) (code int) {
	// тут писать не нужно
	codes := []int{
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusRequestTimeout,
	}
	idx := rand.IntN(len(codes))

	return codes[idx]
}

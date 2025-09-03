package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

// Написать функцию, которая запрашивает  URL  из списка и в случае положительного кода 200 выводит
// в stdout в отдельной строке url: <url>, code: <statusCode>

// В случае ошибки выводит в отдельной строке url: <url>, code: <statusCode>

// Функция должна завершаться при отмене контекста.
// Доп задание: реализовать ограничение количества одновременно запущенных горутин.
func fetchParallel(ctx context.Context, urls []string) {
	const concurrentLimit = 10

	httpClient := &http.Client{}

	sem := make(chan struct{}, concurrentLimit)
	wg := sync.WaitGroup{}

	mch := make(chan string)
	defer close(mch)

	go func() {
		for msg := range mch {
			fmt.Println(msg)
		}
	}()

	wg.Add(len(urls))
	for _, u := range urls {
		select {
		case <-ctx.Done():
			return
		case sem <- struct{}{}:
		}

		u := u

		go func() {
			defer wg.Add(-1)
			defer func() { <-sem }()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				fmt.Printf("http.NewRequestWithContext: %v\n", err)
				return
			}

			resp, err := httpClient.Do(req)
			if err != nil {
				fmt.Printf("http client Do: %v\n", err)
				return
			}

			defer resp.Body.Close()

			select {
			case mch <- fmt.Sprintf("url: %s, code: %d", u, resp.StatusCode):
			case <-ctx.Done():
				return
			}
		}()
	}

	wg.Wait()
}

func main() {

}

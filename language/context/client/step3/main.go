package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func sendRequest(ctx context.Context, duration time.Duration) (*http.Response, error) {
	fmt.Println("sendRequest started")
	defer fmt.Println("sendRequest finished")

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	q := req.URL.Query()
	q.Add("duration", duration.String())
	req.URL.RawQuery = q.Encode()
	req = req.WithContext(ctx)

	fmt.Println(fmt.Sprintf("request:\n%s\n", req.URL.String()))

	return client.Do(req)
}

func main() {
	fmt.Println("main started")
	defer fmt.Println("main finished")

	duration := flag.Duration("duration", 1*time.Second, "duration to send to server")
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	go func() {
		if _, err := fmt.Scanln(); err != nil {
			fmt.Printf("scanln: %v", err)
			return
		}
		fmt.Println("new line scanned")
		cancel()
	}()

	resp, err := sendRequest(ctx, *duration)
	if err != nil {
		fmt.Printf("send request: %v\n", err)
		return
	}

	fmt.Println("response:")
	fmt.Println(fmt.Sprintf("%s", resp.Status))

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read all: %v\n", err)
		return
	}

	fmt.Printf("body:\n%s", b)
}

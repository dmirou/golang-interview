package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func sendRequest(duration time.Duration) (*http.Response, error) {
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

	fmt.Println(fmt.Sprintf("request:\n%s\n", req.URL.String()))

	return client.Do(req)
}

func main() {
	fmt.Println("main started")
	defer fmt.Println("main finished")

	duration := flag.Duration("duration", 1*time.Second, "duration to send to server")
	flag.Parse()

	resp, err := sendRequest(*duration)
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

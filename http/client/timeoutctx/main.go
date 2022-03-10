package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// response received, len is 16916 bytes
const requestTimeout = 2500 * time.Millisecond

// can not get response: Get "https://www.google.com/": context deadline exceeded
// const requestTimeout = 500 * time.Millisecond

func main() {
	cl := http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, "https://google.com", bytes.NewReader(nil),
	)
	if err != nil {
		fmt.Printf("can not create request: %v\n", err)
		return
	}

	resp, err := cl.Do(req)
	if err != nil {
		fmt.Printf("can not get response: %v\n", err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("can not get response body: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("response received, len is %d bytes", len(b))
}

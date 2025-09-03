package main

import (
	"fmt"
	"sync"
)

const count = 10

func ping(in <-chan struct{}, out chan<- struct{}) {
	defer close(out)
	for i := 0; i < count; i++ {
		<-in
		fmt.Println("ping")
		out <- struct{}{}
	}
}

func pong(in <-chan struct{}, out chan<- struct{}) {
	defer close(out)
	for i := 0; i < count; i++ {
		<-in
		fmt.Println("pong")
		if i != count-1 {
			out <- struct{}{}
		}
	}
}

func main() {
	pingTrigger := make(chan struct{})
	pongTrigger := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ping(pingTrigger, pongTrigger)
	}()

	go func() {
		defer wg.Done()
		pong(pongTrigger, pingTrigger)
	}()

	pingTrigger <- struct{}{}

	wg.Wait()
}

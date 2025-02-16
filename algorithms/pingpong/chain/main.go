package main

import (
	"fmt"
	"sync"
)

func print(in chan struct{}, wg *sync.WaitGroup, message string) chan struct{} {
	out := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-in:
				if ok {
					fmt.Println(message)
					out <- struct{}{}
					continue
				}
				close(out)
				return
			}
		}
	}()

	return out
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan struct{})

	out := print(ch, &wg, "-> ping")
	out = print(out, &wg, "<- pong")

	for i := 0; i < 20; i++ {
		ch <- struct{}{}
		<-out
	}
	close(ch)

	wg.Wait()
}

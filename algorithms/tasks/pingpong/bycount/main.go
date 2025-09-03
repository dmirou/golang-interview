package main

import (
	"fmt"
	"strconv"
	"sync"
)

func print(in, out chan struct{}, start bool, wg *sync.WaitGroup, message string) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		do := func(i int) {
			fmt.Println(message + strconv.Itoa(i))
			out <- struct{}{}
		}

		for i := 0; i < 20; i++ {
			if i == 0 && start {
				do(i + 1)
				continue
			}

			select {
			case <-in:
				do(i + 1)
				continue
			}
		}
	}()
}

func main() {
	var wg sync.WaitGroup

	in := make(chan struct{}, 1)
	out := make(chan struct{})

	print(in, out, true, &wg, "-> ping")
	print(out, in, false, &wg, "<- pong")

	wg.Wait()
}

// fan-in redirects input from a few channels into one channel
// https://go.dev/blog/pipelines
package main

import (
	"fmt"
	"sync"
)

func merge(chs ...chan int) chan int {
	var wg sync.WaitGroup
	var out = make(chan int)

	output := func(ch chan int) {
		for item := range ch {
			out <- item
		}
		wg.Done()
	}

	wg.Add(len(chs))
	for _, ch := range chs {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen(n int) chan int {
	res := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			res <- i
		}
		close(res)
	}()

	return res
}

func main() {
	ch1 := gen(10)
	ch2 := gen(10)
	ch3 := gen(5)

	res := merge(ch1, ch2, ch3)

	for i := range res {
		fmt.Println(i)
	}
}

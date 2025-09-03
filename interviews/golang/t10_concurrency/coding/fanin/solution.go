package main

import (
	"fmt"
	"sync"
)

func gen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func main() {
	ch1 := gen(1)
	ch2 := gen(2)
	ch3 := gen(3)

	res := merge(ch1, ch2, ch3)

	for i := range res {
		fmt.Println(i)
	}
}

func merge(chs ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

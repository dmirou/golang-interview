// fan-out send data from one channel to a few workers
package main

import (
	"fmt"
	"sync"
)

func gen(n int) chan string {
	res := make(chan string)
	go func() {
		for i := 0; i < n; i++ {
			res <- fmt.Sprintf("%d", i)
		}
		close(res)
	}()

	return res
}

func main() {
	in := gen(100)
	out := workInParallel(in, 5)
	for v := range out {
		fmt.Println(v)
	}
}

func workInParallel(in chan string, workersNum int) chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	wg.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- work(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func work(in string) string {
	return in + in
}

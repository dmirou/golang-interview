// fan-out send data from one channel to a few workers
package main

import (
	"fmt"
	"sync"
)

func work(name string, in, out chan string) {
	for v := range in {
		out <- fmt.Sprintf("worker %q, value: %s", name, v)
	}
}

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

const workersNum = 5

func main() {
	in := gen(100)

	var out = make(chan string)

	var wg sync.WaitGroup

	wg.Add(workersNum)
	for i := 0; i < workersNum; i++ {
		go func(i int) {
			work(fmt.Sprintf("%d", i+1), in, out)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for v := range out {
		fmt.Println(v)
	}
}

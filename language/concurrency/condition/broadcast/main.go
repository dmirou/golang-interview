package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	for i := range 10 {
		go func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()

			cond.Wait()
			fmt.Println(i)
		}(i)
	}

	time.Sleep(100 * time.Millisecond) // wait for goroutines to be ready
	cond.Broadcast()
	time.Sleep(100 * time.Millisecond) // wait for goroutines to be woken up
}

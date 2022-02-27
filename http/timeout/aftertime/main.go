package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const timeout = 350 * time.Millisecond

func init() {
	rand.Seed(time.Now().UnixNano())
}

func slowOperation() int {
	delay := rand.Intn(500) + 200
	time.Sleep(time.Millisecond * time.Duration(delay))
	return delay
}

func withTimeout(timeout time.Duration) (int, error) {
	// add buffer to unblock slow operations to send result to channel
	// after complete (after timeout)
	ch := make(chan int, 1)
	go func() {
		fmt.Println("goroutine started")
		ch <- slowOperation()
		fmt.Println("goroutine finished")
	}()

	select {
	case v := <-ch:
		return v, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("operation stopped by timeout: %v", timeout)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			start := time.Now()
			v, err := withTimeout(timeout)
			fmt.Printf("task %d completed with value: %v and error: %v for: %s\n",
				i, v, err, time.Since(start))
		}()
	}

	wg.Wait()
	fmt.Println("waitGroup unlocked")
	time.Sleep(1 * time.Second)
}

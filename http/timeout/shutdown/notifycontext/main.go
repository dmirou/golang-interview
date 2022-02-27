package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const timeout = 1500 * time.Millisecond
const minLatency = 1000
const maxLatency = 2000

func init() {
	rand.Seed(time.Now().UnixNano())
}

func slowOperation() int {
	latency := minLatency + rand.Intn(maxLatency-minLatency)
	time.Sleep(time.Millisecond * time.Duration(latency))
	return latency
}

func operation(ctx context.Context) (int, error) {
	ch := make(chan int, 1)
	go func() {
		fmt.Println("goroutine started")
		ch <- slowOperation()
		close(ch)
		fmt.Println("goroutine finished")
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case v := <-ch:
		return v, nil
	}
}

func withTimeout(ctx context.Context, timeout time.Duration) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return operation(ctx)
}

func main() {
	rootCtx := context.Background()
	ctx, stop := signal.NotifyContext(rootCtx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			start := time.Now()
			v, err := withTimeout(ctx, timeout)
			fmt.Printf("task %d completed for: %s with value: %v and error: %v \n",
				i, time.Since(start), v, err)
		}()
	}

	wg.Wait()
	fmt.Println("shutdown...")
	time.Sleep(maxLatency * time.Millisecond)
	fmt.Println("finished")
}

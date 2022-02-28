package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
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
	g, gCtx := errgroup.WithContext(rootCtx)
	for i := 0; i < 10; i++ {
		i := i
		g.Go(func() error {
			start := time.Now()
			v, err := withTimeout(gCtx, timeout)
			fmt.Printf("task %d completed for: %s with value: %v and error: %v \n",
				i, time.Since(start), v, err)

			return err
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("errgroup unblocked with error: %s\n", err)
	} else {
		fmt.Println("errgroup unlocked")
	}
	time.Sleep(1 * time.Second)
}

package main

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

/**
127.0.0.1
127.0.0.2
127.0.0.3
*/

const retries = 3

var ErrNotFound = errors.New("not found")

func Get(ctx context.Context, address, key string) (string, error) {
	return "", nil // already implemented
}

func DistributedGet(ctx context.Context, addresses []string, key string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	out := make(chan string)
	var found atomic.Bool

	var wg sync.WaitGroup
	wg.Add(len(addresses))

	for _, address := range addresses {
		go func() {
			defer wg.Done()

			for i := 0; i < retries && !found.Load(); i++ {
				val, err := Get(ctx, address, key)
				if err != nil {
					if errors.Is(err, ErrNotFound) {
						return
					}
					continue
				}

				select {
				case out <- val:
					found.Store(true)
				default:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	select {
	case val, ok := <-out:
		if !ok {
			return "", errors.New("all failed")
		}
		return val, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	DistributedGet(ctx, []string{"addr1", "addr2", "addr3"}, "some-key")
}

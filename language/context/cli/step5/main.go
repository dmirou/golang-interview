package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"
)

type Job struct {
	name string
	args []interface{}
}

func getPendingJobs(ctx context.Context) chan Job {
	ch := make(chan Job)
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			n := rand.Int63n(999)
			ch <- Job{
				name: fmt.Sprintf("job#%d", n),
				args: []interface{}{n, n + 1, n + 2},
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	return ch
}

func executeJob(job Job) {
	fmt.Printf("job: %s started with args: %v\n", job.name, job.args)
	defer fmt.Printf("job: %s finished with args: %v\n\n", job.name, job.args)
	time.Sleep(time.Duration(100+rand.Int63n(500)) * time.Millisecond)
}

func main() {
	rand.Seed(time.Now().Unix())

	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for job := range getPendingJobs(ctx) {
		executeJob(job)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	name string
	args []interface{}
}

func getPendingJobs() chan Job {
	ch := make(chan Job)
	go func() {
		for {
			n := rand.Int63n(999)
			ch <- Job{
				name: fmt.Sprintf("job#%d", n),
				args: []interface{}{n, n + 1, n + 2},
			}
			time.Sleep(2 * time.Second)
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

	for job := range getPendingJobs() {
		executeJob(job)
	}
}

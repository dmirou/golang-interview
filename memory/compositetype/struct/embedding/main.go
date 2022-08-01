package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Job struct {
	name string
	*log.Logger
}

func NewJob(name string) *Job {
	return &Job{
		name,
		log.New(os.Stdout,
			fmt.Sprintf("Job: %s : ", name),
			log.Ldate|log.Ltime,
		)}
}

func (j *Job) Panic(v ...interface{}) {
	j.Println("panic catched: ", v)
}

func main() {
	j := NewJob("makePizza")
	j.Println("started")
	time.Sleep(1 * time.Second)
	j.Panic("oh my God")
	j.Println("finished")
}

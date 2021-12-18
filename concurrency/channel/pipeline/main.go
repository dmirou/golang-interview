package main

import (
	"fmt"
)

type item float64

type myWaitGroup struct {
	sema  chan struct{}
	count int
}

func NewWg() *myWaitGroup {
	return &myWaitGroup{
		sema: make(chan struct{}, 1),
	}
}

func (wg *myWaitGroup) Inc() {
	wg.sema <- struct{}{}
	wg.count++
	<-wg.sema
}

func (wg *myWaitGroup) Done() {
	wg.sema <- struct{}{}
	wg.count--
	<-wg.sema
}

func (wg *myWaitGroup) Wait() {
	done := false
	for !done {
		wg.sema <- struct{}{}
		if wg.count == 0 {
			done = true
		}
		<-wg.sema
	}
}

type work func(in <-chan item, out chan<- item)

func generate(_ <-chan item, out chan<- item) {
	defer close(out)
	for i := 0; i < 100; i++ {
		fmt.Printf("generate: %d\n", i)
		out <- item(i)
	}
}

func squaring(in <-chan item, out chan<- item) {
	defer close(out)
	for num := range in {
		square := num * num
		fmt.Printf("squaring: %.f -> %.f\n", num, square)
		out <- square
	}
}

func inc(in <-chan item, out chan<- item) {
	defer close(out)
	for num := range in {
		added := num + 1
		fmt.Printf("inc: %.f -> %.f\n", num, added)
		out <- added
	}
}

func output(in <-chan item, _ chan<- item) {
	for i := range in {
		fmt.Printf("output: %.f\n", i)
	}
}

func main() {
	var workers = []work{
		generate,
		squaring,
		inc,
		output,
	}

	in := make(chan item)
	var wg = NewWg()

	for _, w := range workers {
		out := make(chan item)
		wg.Inc()
		go func(w work, in, out chan item, wg *myWaitGroup) {
			w(in, out)
			wg.Done()
		}(w, in, out, wg)
		in = out
	}

	wg.Wait()
}

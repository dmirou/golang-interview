package main

import (
	"fmt"
	"sync"
)

type item float64

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
	var wg = &sync.WaitGroup{}

	for _, w := range workers {
		out := make(chan item)
		wg.Add(1)
		go func(w work, in, out chan item, wg *sync.WaitGroup) {
			w(in, out)
			wg.Done()
		}(w, in, out, wg)
		in = out
	}

	wg.Wait()
}

package main

import (
	"fmt"
)

type item float64

type work func(in <-chan item) chan item

func generate(_ <-chan item) chan item {
	out := make(chan item)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("generate: %d\n", i)
			out <- item(i)
		}
		close(out)
	}()

	return out
}

func squaring(in <-chan item) chan item {
	out := make(chan item)
	go func() {
		for num := range in {
			square := num * num
			fmt.Printf("squaring: %.f -> %.f\n", num, square)
			out <- square
		}
		close(out)
	}()

	return out
}

func inc(in <-chan item) chan item {
	out := make(chan item)
	go func() {
		for num := range in {
			added := num + 1
			fmt.Printf("inc: %.f -> %.f\n", num, added)
			out <- added
		}
		close(out)
	}()

	return out
}

func main() {
	var workers = []work{
		generate,
		squaring,
		inc,
	}

	in := make(chan item)

	for _, w := range workers {
		in = w(in)
	}

	for res := range in {
		fmt.Printf("output: %.f\n", res)
	}
}

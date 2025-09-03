package main

import "fmt"

func gen(start, finish, inc int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := start; i <= finish; i += inc {
			out <- i
		}
	}()

	return out
}
func main() {
	for i := range gen(1, 5, 1) {
		fmt.Println(i)
	}
}

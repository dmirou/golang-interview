package main

import "fmt"

func sum(a ...int) int {
	fmt.Printf("a is %#v\n", a)

	s := 0
	for _, v := range a {
		s += v
	}

	return s
}

func main() {
	a := []int{1, 2, 5}
	fmt.Printf("sum(a...) is %d", sum(a...))
}

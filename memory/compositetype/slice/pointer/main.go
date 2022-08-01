package main

import (
	"fmt"
)

func printElems(elems []*int) {
	fmt.Println("Print elems")
	for _, v := range elems {
		fmt.Printf("%d ", *v)
	}
	fmt.Println()
}

func double(elems []*int) {
	fmt.Printf("elems pointer %p inside double\n", elems)

	sum := 0
	for _, v := range elems {
		sum += *v
	}

	elems = append(elems, &sum)
}

func main() {
	elems := make([]*int, 5, 10)

	for i, v := range []int{1, 2, 3, 4, 5} {
		v := v // comment to have the same address in all elems items
		fmt.Printf("v pointer is %p, value is %d\n", &v, v)
		elems[i] = &v
	}
	printElems(elems)

	fmt.Printf("elems pointer %p outside double\n", elems)
	double(elems)
	printElems(elems)
}

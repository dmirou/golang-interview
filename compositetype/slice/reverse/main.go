// main reverses slice elements using the same base array
package main

import "fmt"

func reverse(s []int32) {
	if s == nil || len(s) == 1 {
		return
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func printReverse(s []int32) {
	fmt.Printf("source: %v, ", s)
	reverse(s)
	fmt.Printf("reverse: %v\n", s)
}

func main() {
	printReverse([]int32{})
	printReverse([]int32{1})
	printReverse([]int32{1, 2})
	printReverse([]int32{1, 2, 3})
	printReverse([]int32{1, 2, 3, 4})
}

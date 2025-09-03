// after init 	[0] ([0], _, _)
// after append [0] ([0], 1, _)
// after copy	[1] ([1], 1, _)
//
// on mutate: panic: runtime error: index out of range [1] with length 1
// it can be fixing by extending len:
// nums = nums[:2]
// after extending ([1, 1], _), len 2, cap 3
// after mutate ([1, 2], _)
package main

import "fmt"

func main() {
	nums := make([]int, 1, 3)
	fmt.Println(nums) // ?

	appendSlice(nums, 1)
	fmt.Println(nums) // ?

	copySlice(nums, []int{1, 2})
	fmt.Println(nums) // ?

	fmt.Println("len", len(nums), "cap", cap(nums))

	mutateSlice(nums, 1, 4)
	fmt.Println(nums) // ?
}

func appendSlice(s []int, value int) {
	s = append(s, value)
}

func copySlice(s1 []int, s2 []int) {
	copy(s1, s2)
}

func mutateSlice(s []int, index, value int) {
	s[index] = value
}

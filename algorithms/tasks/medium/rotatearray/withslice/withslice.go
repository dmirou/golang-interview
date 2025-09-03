package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[(i+k)%n] = nums[i]
	}
	copy(nums, res)
}

func main() {
	src := []int{1, 2, 3, 4, 5}

	dst1 := make([]int, len(src))
	copy(dst1, src)
	k := 4
	rotate(dst1, k)
	fmt.Printf("src=%v, k=%v, dst=%v\n", src, k, dst1)
}

package main

// [1, 2], k = 2, [1, 2]
// [1, 2, 3], k = 2, [2, 3, 1]

// [1, 2, 3, 4, 5], k = 3, [3, 4, 5, 1, 2]
// [5, 4, 3, 2, 1]
// [0, k) [3, 4, 5, 2, 1]
// [k, len(nums)-1] [3, 4, 5, 1, 2]

import (
	"fmt"

	"slices"
)

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func main() {
	src := []int{1, 2, 3, 4, 5}

	dst1 := make([]int, len(src))
	copy(dst1, src)
	k := 4
	rotate(dst1, k)
	fmt.Printf("src=%v, k=%v, dst=%v\n", src, k, dst1)
}

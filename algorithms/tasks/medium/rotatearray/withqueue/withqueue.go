package main

import "fmt"

// [1, 2], k = 2, [1, 2]
// [1, 2, 3], k = 2, [2, 3, 1]

// [1, 2, 3, 4, 5] k = 2 => [4, 5, 1, 2, 3]
// i=0 [4, 2, 3, 4, 5], buff[1]
// i=1 [4, 5, 3, 4, 5], buff[1, 2]
// i=2, i-k >= 0, buff = append(buff, a[i-k]), a[i] = buff[0], remove buff[0] [4, 5, 1, 4, 5], buff[2, 3]
// i=3 [4, 5, 1, 2, 5], buff[3, 4]
// i=4 [4, 5, 1, 2, 3], end

// a[i] = a[i-k], a[3] = 1, a[4] = 2, a[5] = 3
// if i - k < 0; a[i] = a[i - k + len(a)]; a[0] = 4 a[3], a[1] = 5 a[4]

// [1, 2, 3, 4, 5] k = 3 => [3, 4, 5, 1, 2]

// [1, 2, 3, 4, 5, 6, 7], k = 3
// queue [1], nums[0] = 5
// queue [1 2], nums[1] = 6
// queue [1 2 3], nums[2] = 7
// nums[3] = 1, queue [2 3 4]
// nums[4] = 2, queue [3 4 5]
func rotate(nums []int, k int) {
	k = k % len(nums)

	if len(nums) == 1 || k == 0 || len(nums) == k {
		return
	}

	queue := make([]int, 0, k)

	n := len(nums)

	for i := range nums {
		from := i - k
		if from < 0 {
			from = from + n
			queue = append(queue, nums[i])
			nums[i] = nums[from]
			continue
		}
		src := nums[i]
		nums[i] = queue[0]
		queue = append(queue[1:], src)
	}
}

func main() {
	src := []int{1, 2, 3, 4, 5}

	dst1 := make([]int, len(src))
	copy(dst1, src)
	k := 1
	rotate(dst1, k)
	fmt.Printf("src=%v, k=%v, dst=%v\n", src, k, dst1)
}

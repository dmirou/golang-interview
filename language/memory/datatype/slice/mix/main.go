package main

import "fmt"

func main() {
	// task1()
	// task2()
	task3()
}

func task1() {
	s := make([]int, 1, 3)
	fmt.Println("make([]int, 1, 3): ", s) // [0]

	appendSlice(s, 2)
	fmt.Println("after append: ", s) // [0]

	copySlice(s, []int{4, 5})
	fmt.Println("after copy: ", s) // [4]

	mutateSlice(s, 1, 4)
	fmt.Println("after mutate: ", s) // panic
}

func task2() {
	s := make([]int, 3, 3)
	fmt.Println("make([]int, 3, 3): ", s) // [0 0 0]

	appendSlice(s, 2)
	fmt.Println("after append: ", s) // [0 0 0]

	copySlice(s, []int{4, 5})
	fmt.Println("after copy: ", s) // [4 5 0]

	mutateSlice(s, 1, 4)
	fmt.Println("after mutate: ", s) // [4 4 0]
}

func appendSlice(s []int, value int) {
	s = append(s, value)
}

func copySlice(s1, s2 []int) {
	copy(s1, s2)
}

func mutateSlice(s []int, ind, value int) {
	s[ind] = value
}

func task3() {
	nums := []int{1, 2, 3}

	addNum(nums[0:2])
	fmt.Println("after addNum: ", nums) // [1 2 4]

	addNums(nums[0:2])
	fmt.Println("after addNums: ", nums) // [1 2 4]
}

func addNum(nums []int) {
	nums = append(nums, 4)
}

func addNums(nums []int) {
	nums = append(nums, 5, 6)
}

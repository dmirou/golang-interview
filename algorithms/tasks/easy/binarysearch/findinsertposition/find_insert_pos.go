package main

// [1], 1
// left = 0, right = 0, return 1
//
// [2], 1
// left = 0, right = 0, return 0
//
// [1 3], 2
// left = 0, right = 1, middle = 0
// left = 0, right = 1, middle = 0
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		if left == right {
			if target <= nums[left] {
				return left
			}
			return left + 1
		}

		if target <= nums[left] {
			return left
		}
		if target == nums[right] {
			return right
		}
		if target > nums[right] {
			return right + 1
		}

		middle := left + (right-left)/2

		if middle == left {
			return middle + 1
		}

		if target <= nums[middle] {
			right = middle
			continue
		}
		left = middle + 1
	}
}

func main() {

}

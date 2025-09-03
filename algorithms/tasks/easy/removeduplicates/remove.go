package removeduplicates

// 1 1 1 2 3 3 3 4 5 => 1 2 3 4 5 x x x x, 5

// 1 1 2
// slow = 1, slowVal=1, i = 1
// 1 == 1, i++, i = 2
// 2 != 1, nums[slow] = nums[i], slowVal = nums[i], slow++
// return slow

func removeDuplicates(nums []int) int {
	slow := 1
	slowVal := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != slowVal {
			nums[slow] = nums[i]
			slowVal = nums[i]
			slow++
		}
	}

	return slow
}

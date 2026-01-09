package leftrightsumdiff

func leftRightDifference(nums []int) []int {
	leftSums := make([]int, len(nums))
	rightSums := make([]int, len(nums))

	// [1 2 3 4 5] len 5
	// ls [0 1 3 6 10]
	// rs [14 12 9 5 0]

	// [10,4,8,3]
	// ls [0, 10, 14, 22]
	// rs [15 11 3 0]
	// diff [15,1,11,22]

	for i := 1; i < len(nums); i++ {
		leftSums[i] = leftSums[i-1] + nums[i-1]

		// 4-1 = 3, 2, 1, 0
		j := len(nums) - 1 - i
		rightSums[j] = rightSums[j+1] + nums[j+1]
	}

	diff := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		diff[i] = leftSums[i] - rightSums[i]
		if diff[i] < 0 {
			diff[i] *= -1
		}
	}

	return diff
}

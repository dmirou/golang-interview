package minstartvalue

// Input: nums = [-3,2,-3,4,2]
// Output: 5

// Input: nums = [1,2]
// Output: 1

// Input: nums = [1,-2,-3]
// Output: 5

// x + sum >= 1
// x >= 1 - sum
func minStartValue(nums []int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	minsum := sums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
		minsum = min(minsum, sums[i])
	}

	start := 1 - minsum
	// to return at least minimum positive value
	if start < 1 {
		start = 1
	}
	return start
}

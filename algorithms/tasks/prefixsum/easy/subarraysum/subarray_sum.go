package subarraysum

func subarraySum(nums []int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		start := max(0, i-nums[i])
		leftSum := 0
		if start > 0 {
			leftSum = sums[start-1]
		}
		sum += sums[i] - leftSum
	}
	return sum
}

package productexceptself

func productExceptSelf(nums []int) []int {
	prefixes := make([]int, len(nums))
	suffixes := make([]int, len(nums))

	prefixes[0] = 1
	suffixes[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		prefixes[i] = prefixes[i-1] * nums[i-1]
		suffixes[len(nums)-1-i] = suffixes[len(nums)-i] * nums[len(nums)-i]
	}

	results := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		results[i] = prefixes[i] * suffixes[i]
	}

	return results
}

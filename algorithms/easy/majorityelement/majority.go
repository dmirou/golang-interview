package majorityelement

func majorityElement(nums []int) int {
	count := make(map[int]int)

	max := nums[0]
	maxCount := 1
	count[nums[0]]++

	for i := 1; i < len(nums); i++ {
		count[nums[i]]++
		if maxCount < count[nums[i]] {
			max = nums[i]
			maxCount = count[nums[i]]
		}
	}

	return max
}

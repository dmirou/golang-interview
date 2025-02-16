package containsduplicates

func containsNearbyDuplicate(nums []int, k int) bool {
	numi := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		idx, ok := numi[nums[i]]
		if !ok {
			numi[nums[i]] = i
			continue
		}
		if i-idx <= k {
			return true
		}
		numi[nums[i]] = i
	}

	return false
}

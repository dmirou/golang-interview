package countvalidselections

/*
nums [1 0 1]
sums [1 1 2]
i = 1, ls 1, rs = 2 - 1 = 1

nums [2 1 5 0 8 1]
sums [2 3 8 8 16 17]
i = 3 ls 8, rs 9

nums [16,13,10, 0, 0, 0,10, 6, 7, 8, 7]
sums [16 29 39 39 39 39 49 55 62 70 77]
expected 3
*/
func countValidSelections(nums []int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			continue
		}

		ls := 0
		if i-1 >= 0 {
			ls = sums[i-1]
		}
		rs := 0
		if i < len(nums)-1 {
			rs = sums[len(nums)-1] - sums[i]
		}

		if ls == rs {
			count += 2
		}
		if ls-rs == 1 || ls-rs == -1 {
			count += 1
		}
	}

	return count
}

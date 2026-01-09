package findpivotindex

// Input: nums = [1,7,3,6,5,6]
// Output: 3
// ls = 1 + 7 + 3 = 11
// rs = 5 + 6 = 11

// Input: nums = [1,2,3]
// Output: -1

// Input: nums = [2,1,-1]
// Output: 0

// Input: nums = [0]
// Output: 0

// Input: nums = [1]
// Output: -1

func pivotIndex(nums []int) int {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	leftsum := 0
	for i := 0; i < len(nums); i++ {
		if leftsum == sums[len(sums)-1]-sums[i] {
			return i
		}
		leftsum = sums[i]
	}

	return -1
}

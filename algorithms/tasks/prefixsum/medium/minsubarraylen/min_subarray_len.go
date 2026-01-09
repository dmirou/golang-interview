package minsubarraylen

/*
Given an array of positive integers nums and a positive integer target, return
the minimal length of a subarray whose sum is greater than or equal to target.
If there is no such subarray, return 0 instead.

Example 1:
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:
Input: target = 4, nums = [1,4,4]
Output: 1

Example 3:
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0

Constraints:
1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 104
*/
func minSubArrayLen(target int, nums []int) int {
	// idxs = [0, 1, 2, 3, 4,  5,  6]
	// nums = [2, 3, 1, 2, 4,  3], target = 7
	// sums = [0, 2, 5, 6, 8, 12, 15]
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	// sum(subarr(2,4)) = 1 + 2 + 4 = 7 = 12 - 5 = sums(5)-sums(2)
	// sum(subarr(0,0)) = 2 - 0 = 2 = sums(1)-sums(0)
	// sum(subarr(i,j)) = sums[j+1] - sums[i]
	// sum(subarr(i,j)) = sums[j+1] - sums[i] >= target
	// sums[j+1] >= sums[i] + target
	minLen := 0
	// iterate to len(sums)-1 since we need j > i
	for i := 0; i < len(sums)-1; i++ {
		atLeastVal := sums[i] + target

		// Find the leftmost j where sums[j] >= atLeastVal
		j := -1
		left := i + 1 // j must be > i (since we need a subarray of length >= 1)
		right := len(sums) - 1

		for left <= right {
			mid := left + (right-left)/2
			if sums[mid] >= atLeastVal {
				j = mid
				right = mid - 1 // Try to find a smaller valid j
			} else {
				left = mid + 1
			}
		}
		if j == -1 {
			continue
		}

		// Use binary search to find the smallest j
		// where sums[j] >= atLeastVal
		// j := sort.Search(len(sums), func(k int) bool {
		// 	return sums[k] >= atLeastVal
		// })

		if j < len(sums) {
			length := j - i
			if minLen == 0 || length < minLen {
				minLen = length
			}
		}
	}

	return minLen
}

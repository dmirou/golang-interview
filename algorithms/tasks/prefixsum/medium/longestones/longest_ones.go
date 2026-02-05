package longestones

/*
1004. Max Consecutive Ones III
Medium
Topics
premium lock icon
Companies
Hint

Given a binary array nums and an integer k, return the maximum number of consecutive
1's in the array if you can flip at most k 0's.

Example 1:
Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Example 2:
Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
0 <= k <= nums.length
*/
func longestOnes(nums []int, k int) int {
	zeroSums := make([]int, len(nums)+1)
	for i := 1; i < len(zeroSums); i++ {
		delta := 0
		if nums[i-1] == 0 {
			delta = 1
		}
		zeroSums[i] = zeroSums[i-1] + delta
	}

	/*
		index 0  1  2  3  4  5  6  7  8  9  10
		nums {1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0} k = 2
		zs   {0  1  2  3  3  3  3  4  5  6  7  7}
		i0 target2 possible2 maxLen2
		i1 target3 possible6 maxLen5
		i2 target4 possible7 maxLen5
		i3 target5 possible8 maxLen5
		i4 target5 possible8 maxLen5
		i5 target5 possible8 maxLen5
		i6 target5 possible8 maxLen5
		i7 target6 possible9 maxLen5
		i8 target7 possible11 maxLen5
	*/
	maxLen := 0
	for i := 0; i < len(zeroSums); i++ {
		target := k + zeroSums[i]
		l := i
		r := len(zeroSums) - 1
		possible := -1
		for l <= r {
			mid := l + (r-l)/2
			if zeroSums[mid] <= target {
				possible = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if possible != -1 {
			//j = possible-1, len = j-i+1
			maxLen = max(maxLen, possible-i)
		}
	}

	return maxLen
}

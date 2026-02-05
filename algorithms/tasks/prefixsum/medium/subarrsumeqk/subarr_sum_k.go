package subarrsumeqk

/*
560. Subarray Sum Equals K

https://leetcode.com/problems/subarray-sum-equals-k/description/?envType=problem-list-v2&envId=prefix-sum

Medium
Topics
premium lock icon
Companies
Hint

Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
A subarray is a contiguous non-empty sequence of elements within an array.


Example 1:

Input: nums = [1,1,1], k = 2
Output: 2

Example 2:

Input: nums = [1,2,3], k = 3
Output: 2


Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107

For any subarray nums[i:j], its sum equals prefix[j] - prefix[i-1].

Your target is k.
So you're looking for pairs (i, j) where prefix[j] - prefix[i-1] = k.
Rearrange: prefix[j] - k = prefix[i-1].

As you build the prefix sum while scanning left-to-right, you can check if you've already seen
current_prefix - k earlier in the array. If yes, those earlier positions are valid starting
points i-1 for subarrays ending at current index j.

One small implementation detail: initialize your map with {0: 1} to handle subarrays that start from the beginning.
idx	  0 1 2 4
nums [1 2 3], k = 3
sums [0 1 3 6]
sum(1,2) = 2+3 = 5
sus(0,1) = 1+2 = 3 - 0

sum(nums[i:j]) = sums[j]-sums[i-1]
sums[j]-sums[i-1] == k
sums[j] - k == sums[i-1]
*/

/*
nums = [1,1,1], k = 2
i 0, sum 1, needSum -1, count 0, prevSum{0:1, 1:1}
i 1, sum 2, needSum 0,  count 1, prevSum{0:2, 1:1}
i 2, sum 3, needSum 1, count 2, prevSum{0:2, 1:1}

nums = [1,2,3], k = 3
i0, sum 1, needSum -2, prevSum {0:1, 1:1},
i1, sum 3, needSum 0, count 1, prevSum {0:2, 3:1},
i2, sum 6, needSum 3, count 2
*/
func subarraySum(nums []int, k int) int {
	count := 0
	prevSums := make(map[int]int)
	prevSums[0] = 1 // Initialize with {0: 1} to handle subarrays starting from index 0

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		needSum := sum - k
		if cnt, ok := prevSums[needSum]; ok {
			count += cnt
		}
		prevSums[sum]++
	}

	return count
}

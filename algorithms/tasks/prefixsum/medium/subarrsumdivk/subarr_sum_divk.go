package subarrsumdivk

/*
https://leetcode.com/problems/subarray-sums-divisible-by-k/?envType=problem-list-v2&envId=prefix-sum

974. Subarray Sums Divisible by K
Medium
Topics
premium lock icon
Companies

Given an integer array nums and an integer k, return the number of non-empty subarrays
that have a sum divisible by k. A subarray is a contiguous part of an array.

Example 1:
Input: nums = [4,5,0,-2,-3,1], k = 5
Output: 7
Explanation: There are 7 subarrays with a sum divisible by k = 5:
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

Example 2:
Input: nums = [5], k = 9
Output: 0

(sums(i) - sums(j)) % k == 0, j < i
sums(i) % k = sums(j) % k

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
2 <= k <= 104
*/
func subarraysDivByK(nums []int, k int) int {
	reminders := make(map[int]int)
	reminders[0] = 1

	count := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		target := sum % k
		if target < 0 {
			target += k
		}
		count += reminders[target]
		reminders[target]++
	}
	return count
}

/*
[4,5,0,-2,-3,1], k = 5
sum4, target4, count0, reminders [0:1, 4:1]
sum9, target4, count1, reminders [0:1, 4:2]
sum9, target4, count3, reminders [0:1, 4:3]
sum7, target2, count3, reminders [0:1, 2:1, 4:3]
sum4, target4, count6, reminders [0:1, 2:1, 4:4]
sum5, target0, count7, reminders [0:1, 2:1, 4:4, 5:1]

-1 % 2 = -1 but we need 1, reminders [1:1]
-1 + 2 = 1 % 2 = 1, count = 1, reminders [1:2]
10 % 2 = 0, reminders[0] = 1; count2
[-1,2,9], k = 2
Use Testcase
Output 1 Expected 2

Why This Breaks Our Algorithm
Our algorithm relies on:
If (prefix[i] - prefix[j]) % k = 0
Then prefix[i] % k = prefix[j] % k
But in Go, with negative numbers:

prefix[i] = 4, prefix[j] = -1, k = 5
(4 - (-1)) % 5 = 5 % 5 = 0 ✓

But:
4 % 5 = 4
-1 % 5 = -1
4 ≠ -1 ✗
So we'd miss this valid subarray!

To fix it we need to normalize reminders:
if reminder < 0 {
reminder += k
}
so -1 + 5 = 4 will be the same as 4%5
*/

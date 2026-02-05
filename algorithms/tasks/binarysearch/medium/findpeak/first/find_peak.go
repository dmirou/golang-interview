package findpeak

/*
162. Find Peak Element

https://leetcode.com/problems/find-peak-element/description/?envType=problem-list-v2&envId=binary-search

Medium
Topics
premium lock icon
Companies

A peak element is an element that is strictly greater than its neighbors.

Given a 0-indexed integer array nums, find a peak element, and return its index.
If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always
considered to be strictly greater than a neighbor that is outside the array.

You must write an algorithm that runs in O(log n) time.

Example 1:
Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.

Example 2:
Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2,
or index number 5 where the peak element is 6.

Constraints:

1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
nums[i] != nums[i + 1] for all valid i.

1, l0, r1, c0

1 	2
l 	r
c

	l
	r
	c

2 1

1 2 3
l c r

	lrc

3 2 1
l c r
lrc
*/
func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		c := l + (r-l)/2

		cgprev := c == 0 || nums[c] > nums[c-1]
		cgnext := c == len(nums)-1 || nums[c] > nums[c+1]

		if cgprev && cgnext {
			return c
		}

		if cgprev {
			l = c + 1
			continue
		}
		r = c - 1
	}
	return r
}

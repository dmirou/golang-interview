package searchrange

/*
34. Find First and Last Position of Element in Sorted Array

https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=problem-list-v2&envId=binary-search

Medium
Topics
premium lock icon
Companies

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]

Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109

0 1 2 3 4 5
5,7,7,8,8,10 t=8
l 		  r

	m
	  l   r
	  	m
*/
func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			start := m
			end := m

			i := m - 1
			for i >= 0 && nums[i] == target {
				start = i
				i--
			}

			i = m + 1
			for i < len(nums) && nums[i] == target {
				end = i
				i++
			}

			return []int{start, end}
		}

		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return []int{-1, -1}
}

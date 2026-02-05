package targetindices

import "slices"

/*
2089. Find Target Indices After Sorting Array

https://leetcode.com/problems/find-target-indices-after-sorting-array/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies
Hint

You are given a 0-indexed integer array nums and a target element target.
A target index is an index i such that nums[i] == target.
Return a list of the target indices of nums after sorting nums in non-decreasing order.
If there are no target indices, return an empty list. The returned list must be sorted in increasing order.

Example 1:
Input: nums = [1,2,5,2,3], target = 2
Output: [1,2]
Explanation: After sorting, nums is [1,2,2,3,5].
The indices where nums[i] == 2 are 1 and 2.

Example 2:
Input: nums = [1,2,5,2,3], target = 3
Output: [3]
Explanation: After sorting, nums is [1,2,2,3,5].
The index where nums[i] == 3 is 3.

Example 3:
Input: nums = [1,2,5,2,3], target = 5
Output: [4]
Explanation: After sorting, nums is [1,2,2,3,5].
The index where nums[i] == 5 is 4.

Constraints:

1 <= nums.length <= 100
1 <= nums[i], target <= 100

	0 1 2 3 4

[1,2,2,3,5] t2
l r c nums[c] 	f
0 4 2 	2	  	2
0 1 0	1		2
1 1 1	2		1
*/
func targetIndices(nums []int, target int) []int {
	slices.Sort(nums)
	l := 0
	r := len(nums) - 1
	first := -1
	for l <= r {
		c := l + (r-l)/2
		if nums[c] >= target {
			if nums[c] == target {
				first = c
			}
			r = c - 1
		} else {
			l = c + 1
		}
	}
	if first == -1 {
		return nil
	}
	indices := make([]int, 0)
	for i := first; i < len(nums); i++ {
		if nums[i] == target {
			indices = append(indices, i)
			continue
		}
		break
	}
	return indices
}

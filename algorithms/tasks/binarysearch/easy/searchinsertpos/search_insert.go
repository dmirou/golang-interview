package searchinsertpos

/*
35. Search Insert Position

https://leetcode.com/problems/search-insert-position/description/?envType=problem-list-v2&envId=binary-search

Solved
Easy
Topics
premium lock icon
Companies

Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4


Constraints:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
*/

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	c := -1
	for l < r {
		c = l + (r-l)/2
		if nums[c] == target {
			return c
		}
		if nums[c] < target {
			l = c + 1
		} else {
			r = c - 1
		}
	}
	if l == r {
		if nums[l] < target {
			return l + 1
		}
		return l
	}
	if l > r && nums[l] > target {
		return l
	}

	return c
}

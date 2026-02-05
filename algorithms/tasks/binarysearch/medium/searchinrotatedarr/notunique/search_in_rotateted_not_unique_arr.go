package notunique

/*
81. Search in Rotated Sorted Array II

https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/

Medium
Topics
premium lock icon
Companies

There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or
false if it is not in nums.

You must decrease the overall operation steps as much as possible.

Example 1:
Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true

Example 2:
Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums is guaranteed to be rotated at some pivot.
-104 <= target <= 104

Follow up: This problem is similar to Search in Rotated Sorted Array, but nums may contain duplicates.
Would this affect the runtime complexity? How and why?

0 1 2 3 4 5 6
2,5,6,0,0,1,2 target = 3
l			r

	c

l	r

	c

l
r
c

5 2 targ 3
l r
c

	l
	r
	c

0 1 2 3 4
1,0,1,1,1  targ 0
l       r

	  c
	l     r
	  c
	l
	r
	c

[1] target 0

	l
	r
	c

0 1 2
1,3,5 target 1
l   r

	c
*/
func search(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		c := l + (r-l)/2
		cv := nums[c]

		if cv == target {
			return true
		}
		lv := nums[l]
		rv := nums[r]

		if cv == lv && cv == rv {
			for l <= r && nums[l] == cv {
				l = l + 1
			}
			continue
		}

		if cv < lv {
			if target > cv && target <= rv {
				l = c + 1
				continue
			}
			r = c - 1
			continue
		}
		if target >= lv && target < cv {
			r = c - 1
			continue
		}
		l = c + 1
	}
	return false
}

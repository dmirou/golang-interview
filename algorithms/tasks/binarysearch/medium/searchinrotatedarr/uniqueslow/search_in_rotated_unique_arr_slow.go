package unique

/*
33. Search in Rotated Sorted Array

https://leetcode.com/problems/search-in-rotated-sorted-array/description/?envType=problem-list-v2&envId=binary-search

Medium
Topics
premium lock icon
Companies

There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly left rotated at an unknown index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).

For example, [0,1,2,4,5,6,7] might be left rotated by 3 indices and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target
if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1

Constraints:

1 <= nums.length <= 5000
-104 <= nums[i] <= 104
All values of nums are unique.
nums is an ascending array that is possibly rotated.
-104 <= target <= 104
*/
func search(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	k := 0
	for i := 1; i < n; i++ {
		if nums[i-1] > nums[i] {
			k = n - i
			break
		}
	}

	/*
		     0 1 2 3 4 5 6
			[0,1,2,4,5,6,7]
			 l
			     r
			   m

			 0 1 2 3 4 5 6
			[4,5,6,7,0,1,2], target = 0, n =7
			k0, i1, nums[0] > nums[1]
			i4, 7 > 0, k = 7 - 4 = 3
			k = 3

			target0
			n7, k3, n-k = 7 - 3 = 4
			l r mido midn 	nums[midn]
			0 6 3	 0		4
			0 2 1    5		1
			0 1 0
	*/

	l := 0
	r := n - 1
	for l <= r {
		mido := l + (r-l)/2
		midn := oldToNew(mido, k, n)

		if nums[midn] == target {
			return midn
		}
		if nums[midn] < target {
			l = mido + 1
		} else {
			r = mido - 1
		}
	}
	return -1
}

func oldToNew(oldi, k, n int) int {
	if oldi >= k {
		return oldi - k
	}
	return oldi + n - k
}

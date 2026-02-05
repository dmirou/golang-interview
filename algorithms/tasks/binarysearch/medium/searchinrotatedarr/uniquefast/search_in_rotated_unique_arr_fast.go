package uniquefast

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
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2

		midv := nums[mid]
		if midv == target {
			return mid
		}

		rv := nums[r]
		lv := nums[l]
		// mid < r && mid > l
		// case 1 2 3 4 5
		// lv < rv
		if midv < rv && midv > lv {
			if midv < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
			continue
		}

		// mid < r && mid < l
		// case 5 1 2 3 4
		// lv > rv
		// mid < r && mid < l
		// case 4 5 1 2 3
		// lv > rv
		if midv < rv && midv < lv {
			if midv > target {
				r = mid - 1
				continue
			}
			if rv >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
			continue
		}

		// mid > r && mid > l
		// case 2 3 4 5 1
		if midv < target {
			l = mid + 1
			continue
		}
		// midv > target
		if lv <= target {
			r = mid - 1
			continue
		}
		l = mid + 1
	}
	return -1
}

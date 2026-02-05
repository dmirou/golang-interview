package intersection

import "slices"

/*
349. Intersection of Two Arrays

https://leetcode.com/problems/intersection-of-two-arrays/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.

Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
8
*/
func intersection(nums1 []int, nums2 []int) []int {
	shortSlice := nums1
	longSlice := nums2
	if len(nums1) > len(nums2) {
		shortSlice = nums2
		longSlice = nums1
	}

	slices.Sort(shortSlice)
	intMap := make(map[int]struct{}, len(shortSlice))
	for i := 0; i < len(longSlice); i++ {
		target := longSlice[i]
		if _, ok := intMap[target]; ok {
			continue
		}

		l := 0
		r := len(shortSlice) - 1
		for l <= r {
			c := l + (r-l)/2
			if shortSlice[c] == target {
				intMap[target] = struct{}{}
				break
			}
			if shortSlice[c] < target {
				l = c + 1
			} else {
				r = c - 1
			}
		}
	}

	intSlice := make([]int, len(intMap))
	i := 0
	for key := range intMap {
		intSlice[i] = key
		i++
	}

	return intSlice
}

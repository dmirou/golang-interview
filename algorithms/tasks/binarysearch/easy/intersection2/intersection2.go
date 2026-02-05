package intersection2

import "slices"

/*
350. Intersection of Two Arrays II

https://leetcode.com/problems/intersection-of-two-arrays-ii/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies

Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must appear as many times as it shows in both arrays
and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.

Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000

Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/
func intersect(nums1 []int, nums2 []int) []int {
	shortSlice := nums1
	longSlice := nums2
	if len(nums1) > len(nums2) {
		shortSlice = nums2
		longSlice = nums1
	}

	slices.Sort(shortSlice)
	intMap := make(map[int]int, len(shortSlice))
	used := make(map[int]struct{}, len(shortSlice))
	for i := 0; i < len(longSlice); i++ {
		target := longSlice[i]
		upper := longSlice[i] + 1
		possible := -1

		l := 0
		r := len(shortSlice) - 1
		for l <= r {
			c := l + (r-l)/2
			if shortSlice[c] < upper {
				possible = c
				l = c + 1
			} else {
				r = c - 1
			}
		}
		if possible != -1 {
			for j := possible; j >= 0; j-- {
				if shortSlice[j] < target {
					break
				}
				if shortSlice[j] != target {
					continue
				}
				if _, ok := used[j]; !ok {
					used[j] = struct{}{}
					intMap[target]++
					break
				}
			}
		}
	}

	intSlice := make([]int, 0, len(shortSlice))
	i := 0
	for key, count := range intMap {
		for j := 0; j < count; j++ {
			intSlice = append(intSlice, key)
			i++
		}
	}

	return intSlice
}

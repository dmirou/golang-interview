package missingnumber

import (
	"cmp"
	"slices"
)

/*
268. Missing Number

https://leetcode.com/problems/missing-number/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies

Given an array nums containing n distinct numbers in the range [0, n],
return the only number in the range that is missing from the array.

Example 1:
Input: nums = [3,0,1]
Output: 2
Explanation:
n = 3 since there are 3 numbers, so all numbers are in the range [0,3].
2 is the missing number in the range since it does not appear in nums.

Example 2:
Input: nums = [0,1]
Output: 2
Explanation:
n = 2 since there are 2 numbers, so all numbers are in the range [0,2].
2 is the missing number in the range since it does not appear in nums.

Example 3:
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation:

n = 9 since there are 9 numbers, so all numbers are in the range [0,9].
8 is the missing number in the range since it does not appear in nums.

Constraints:

n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
All the numbers of nums are unique.

nums [0]
l r c
0 0 0
nums[0] == c
l = c+1 = 1
l > r return l 1

nums [1]
l r c
0 0 0
c < nums[c]
probable = 0
r = -1
return probable 0

nums [0, 1]
l r c
0 1 0
1 1 1
2 1
return 2

nums [1, 2]
l r c
0 1 0
probable = 0
0 -1
return 0

nums [0, 2]
l r c
0 1 0
1 1 1
prob = 1
1 0
*/
func missingNumber(nums []int) int {
	slices.SortFunc(nums, cmp.Compare)
	l := 0
	r := len(nums) - 1
	probable := -1
	for l <= r {
		c := l + (r-l)/2
		if c < nums[c] {
			probable = c
			r = c - 1
		} else {
			l = c + 1
		}
	}
	if l > r {
		return l
	}
	return probable
}

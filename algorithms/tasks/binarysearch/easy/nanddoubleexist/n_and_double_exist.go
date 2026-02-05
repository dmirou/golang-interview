package nanddoubleexist

import "slices"

/*
1346. Check If N and Its Double Exist
https://leetcode.com/problems/check-if-n-and-its-double-exist/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies
Hint

Given an array arr of integers, check if there exist two indices i and j such that :
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]

Example 1:
Input: arr = [10,2,5,3]
Output: true
Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

Example 2:
Input: arr = [3,1,7,11]
Output: false
Explanation: There is no i and j that satisfy the conditions.

Constraints:
2 <= arr.length <= 500
-103 <= arr[i] <= 103

Input
arr = [-10,12,-20,-8,15]
exp true
sort = [-20, -10, -8, 12, 15]
*/
func checkIfExist(arr []int) bool {
	slices.Sort(arr)
	for i := 0; i < len(arr)-1; i++ {
		target := 2 * arr[i]
		if arr[i] < 0 {
			if arr[i]%2 != 0 {
				continue
			}
			target = arr[i] / 2
		}
		l := i + 1
		r := len(arr) - 1
		mid := -1
		for l <= r {
			mid = l + (r-l)/2
			if arr[mid] == target {
				return true
			}
			if arr[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return false
}

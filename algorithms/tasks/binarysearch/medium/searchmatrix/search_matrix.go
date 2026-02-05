package searchmatrix

/*
74. Search a 2D Matrix

https://leetcode.com/problems/search-a-2d-matrix/

Medium
Topics
premium lock icon
Companies

You are given an m x n integer matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Example 1:

Input: matrix = [
[1,3,5,7],
[10,11,16,20],
[23,30,34,60]
], target = 3

Output: true

Example 2:

Input: matrix = [
[1,3,5,7],
[10,11,16,20],
[23,30,34,60]
], target = 13
Output: false

Constraints
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	l := 0
	r := rows - 1
	firstLess := -1
	for l <= r {
		mid := l + (r-l)/2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] < target {
			firstLess = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if firstLess == -1 {
		return false
	}

	l = 1
	r = cols - 1
	for l <= r {
		mid := l + (r-l)/2
		if matrix[firstLess][mid] == target {
			return true
		}
		if matrix[firstLess][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

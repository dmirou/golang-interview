package searchmatrix2

/*
240. Search a 2D Matrix II

https://leetcode.com/problems/search-a-2d-matrix-ii/description/?envType=problem-list-v2&envId=binary-search

Medium
Topics
premium lock icon
Companies

Write an efficient algorithm that searches for a value target
in an m x n integer matrix matrix.

This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.

Example 1:
Input: matrix =
[
[1,4,7,11,15],
[2,5,8,12,19],
[3,6,9,16,22],
[10,13,14,17,24],
[18,21,23,26,30]], target = 5
Output: true

Example 2:
Input: matrix = [
[1,4,7,11,15],
[2,5,8,12,19],
[3,6,9,16,22],
[10,13,14,17,24],
[18,21,23,26,30]], target = 20
Output: false

Constraints:
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matrix[i][j] <= 109
All the integers in each row are sorted in ascending order.
All the integers in each column are sorted in ascending order.
-109 <= target <= 109
*/
func searchMatrix(matrix [][]int, target int) bool {
	/*
		[1,4,7,11,15],
		[2,5,8,12,19],
		[3,6,9,16,22],
		[10,13,14,17,24],
		[18,21,23,26,30]], target = 5

		cr4, cc0
		findCol(cr4, startCol0) -> cr4, cc0
		findRow(maxRow3, cc0) -> cr2, cc0
		findCol(cr2, cc0) -> cr2, cc1
		findRow(cr1, cc1)


		matrix =[[-5]]
		target =-5

		matrix [
		[-1,3]
		] target = -1
		cr0 cc0

	*/
	var (
		cc    int
		cr    = len(matrix) - 1
		equal bool
	)

	for cr != 0 || cc != len(matrix[0])-1 {
		// go right
		cc, equal = findColEqOrFirstBigger(matrix, target, cr, cc)
		if equal {
			return true
		}
		if cc == -1 {
			return false
		}
		// go up
		cr, equal = findRowEqOrFirstLess(matrix, target, cr-1, cc)
		if equal {
			return true
		}
		if cr == -1 {
			return false
		}
	}

	return matrix[0][len(matrix[0])-1] == target
}

func findColEqOrFirstBigger(matrix [][]int, target int, row, startCol int) (ind int, equal bool) {
	l := startCol
	r := len(matrix[0]) - 1
	possible := -1
	for l <= r {
		m := l + (r-l)/2
		if matrix[row][m] == target {
			return m, true
		}
		if matrix[row][m] > target {
			possible = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return possible, false
}

func findRowEqOrFirstLess(matrix [][]int, target int, maxRow, column int) (ind int, equal bool) {
	l := 0
	r := maxRow
	possible := -1
	for l <= r {
		m := l + (r-l)/2
		if matrix[m][column] == target {
			return m, true
		}
		if matrix[m][column] < target {
			possible = m
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return possible, false
}

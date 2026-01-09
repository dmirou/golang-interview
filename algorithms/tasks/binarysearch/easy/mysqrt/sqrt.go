package mysqrt

/*
69. Sqrt(x)

https://leetcode.com/problems/sqrtx/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies
Hint
Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

Example 1:
Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.

Example 2:
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

Constraints:

0 <= x <= 231 - 1
*/
func mySqrt(x int) int {
	l := 0
	r := x
	c := -1
	possible := -1
	for l <= r {
		c = l + (r-l)/2
		if c*c == x {
			return c
		}
		if c*c < x {
			possible = c
			l = c + 1
		} else {
			r = c - 1
		}
	}

	return possible
}

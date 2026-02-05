package perfectsquare

/*
367. Valid Perfect Square

https://leetcode.com/problems/valid-perfect-square/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies
Given a positive integer num, return true if num is a perfect square or false otherwise.

A perfect square is an integer that is the square of an integer. In other words, it is the product of some integer with itself.

You must not use any built-in library function, such as sqrt.

Example 1:
Input: num = 16
Output: true
Explanation: We return true because 4 * 4 = 16 and 4 is an integer.

Example 2:
Input: num = 14
Output: false
Explanation: We return false because 3.742 * 3.742 = 14 and 3.742 is not an integer.

Constraints:
1 <= num <= 231 - 1
*/
func isPerfectSquare(num int) bool {
	l := 1
	r := num
	c := -1
	for l <= r {
		c = l + (r-l)/2
		sq := c * c
		if sq == num {
			return true
		}
		if sq < num {
			l = c + 1
		} else {
			r = c - 1
		}
	}

	return false
}

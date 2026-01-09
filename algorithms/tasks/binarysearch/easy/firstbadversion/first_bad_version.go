package firstbadversion

/* 278. First Bad Version

https://leetcode.com/problems/first-bad-version/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies

You are a product manager and currently leading a team to develop a new product.
Unfortunately, the latest version of your product fails the quality check. Since each version
is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which
causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a
function to find the first bad version. You should minimize the number of calls to the API.

Example 1:

Input: n = 5, bad = 4
Output: 4

Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.

Example 2:
Input: n = 1, bad = 1
Output: 1

Constraints:

1 <= bad <= n <= 231 - 1
*/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return false
}

/*
Input: n = 5, bad = 4
l r c bad
1 5 3 F
4 5 4 T
4 4

Input: n = 5, bad = 1
l r c bad   fb
1 5 3 T		3
1 2 1 T		1
1 1

Input: n = 3, bad = 1
l r c bad  fb
1 3 2  T   2
1 1 1 T 1
*/
func firstBadVersion(n int) int {
	l := 1
	r := n
	firstBad := n
	c := -1
	for l <= r {
		c = l + (r-l)/2
		if isBadVersion(c) {
			firstBad = c
			r = c - 1
		} else {
			l = c + 1
		}
	}
	return firstBad
}

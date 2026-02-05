package arrangecoins

/*
441. Arranging Coins
https://leetcode.com/problems/arranging-coins/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies

You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where
the ith row has exactly i coins. The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.

Example 1:

Input: n = 5
Output: 2
Explanation: Because the 3rd row is incomplete, we return 2.
1 *
2 * *
3 * * -

Example 2:
Input: n = 8
Output: 3
Explanation: Because the 4th row is incomplete, we return 3.
1 * 		S 1
2 * *		S 3
3 * * *		S 6
4 * * - -	S 10

an = a1 + (n-1)*d, d =1
an = a1 + n - 1
Sn = (a1 + an)n/2

Sn = n/2 * (2*a1 + n - 1)
Sn = n/2 * (2-n-1)
Sn = n/2 * (1+n) = (n + n*n) / 2
S2 = 3
S3 = 3/2 * 4 = 6
S4 = 2 * 5 = 10

Sn = (n + n*n) / 2

Constraints:
1 <= n <= 2^31 - 1

n = 5
l r rows needCoins possible
1 5 3	 6		   1
1 2 1	 1		   1
2 2 2	 3	       2
3

n = 8
l r rows needCoins possible
1 8 4	 10			1
1 3 2 	 3			2
3 3 3	 6			3
4
*/
func arrangeCoins(n int) int {
	l := 1
	r := n
	rows := -1
	possible := l
	needCoins := 0

	for l <= r {
		rows = l + (r-l)/2
		needCoins = (rows + rows*rows) / 2
		if needCoins == n {
			return rows
		}
		if needCoins < n {
			possible = rows
			l = rows + 1
		} else {
			r = rows - 1
		}
	}
	return possible
}

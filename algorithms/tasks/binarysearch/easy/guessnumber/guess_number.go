package guessnumber

/*
374. Guess Number Higher or Lower

https://leetcode.com/problems/guess-number-higher-or-lower/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked (the number I picked stays the same throughout the game).

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.

Example 1:
Input: n = 10, pick = 6
Output: 6

Example 2:
Input: n = 1, pick = 1
Output: 1

Example 3:
Input: n = 2, pick = 1
Output: 1


Constraints:

1 <= n <= 231 - 1
1 <= pick <= n
*/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *
 *               otherwise return 0
 */
func guess(num int) int { return 0 }

/*
-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.
*/
func guessNumber(n int) int {
	l := 1
	r := n
	c := -1
	for l <= r {
		c = l + (r-l)/2
		ans := guess(c)
		if ans == 0 {
			return c
		}
		if ans > 0 {
			l = c + 1
		} else {
			r = c - 1
		}
	}

	return -1
}

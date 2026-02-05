package nextgreatestletter

/*
744. Find Smallest Letter Greater Than Target

https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/?envType=problem-list-v2&envId=binary-search

Easy
Topics
premium lock icon
Companies
Hint

You are given an array of characters letters that is sorted in non-decreasing order,
and a character target. There are at least two different characters in letters.

Return the smallest character in letters that is lexicographically greater than target.
If such a character does not exist, return the first character in letters.

Example 1:
Input: letters = ["c","f","j"], target = "a"
Output: "c"
Explanation: The smallest character that is lexicographically greater than 'a' in letters is 'c'.

Example 2:
Input: letters = ["c","f","j"], target = "c"
Output: "f"
Explanation: The smallest character that is lexicographically greater than 'c' in letters is 'f'.

Example 3:
Input: letters = ["x","x","y","y"], target = "z"
Output: "x"
Explanation: There are no characters in letters that is lexicographically greater than 'z' so we return letters[0].

Constraints:
2 <= letters.length <= 104
letters[i] is a lowercase English letter.
letters is sorted in non-decreasing order.
letters contains at least two different characters.
target is a lowercase English letter.

Input: letters = ["c","f","j"], target = "a"
l r c  letters[c] 	possible
0 2 1	'f'		 	1
0 1 0   'c' 		0
0 0

Input: letters = ["c","f","j"], target = "c"
Output: "f"
l r c  letters[c] 	possible
0 2 1 		f			1
0 1 0		c
1 1 0
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters) - 1
	c := -1
	possible := -1
	for l <= r {
		c = l + (r-l)/2
		if letters[c] > target {
			possible = c
			r = c - 1
		} else {
			l = c + 1
		}
	}
	if possible != -1 {
		return letters[possible]
	}
	return letters[0]
}

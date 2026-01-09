package shiftingletters

/*
848. Shifting Letters

https://leetcode.com/problems/shifting-letters/description/?envType=problem-list-v2&envId=prefix-sum

Medium
Topics
premium lock icon
Companies

You are given a string s of lowercase English letters and an integer array shifts of the same length.

Call the shift() of a letter, the next letter in the alphabet, (wrapping around so that 'z' becomes 'a').

For example, shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'.
Now for each shifts[i] = x, we want to shift the first i + 1 letters of s, x times.

Return the final string after all such shifts to s are applied.

Example 1:

Input: s = "abc", shifts = [3,5,9]
Output: "rpl"
Explanation: We start with "abc".
After shifting the first 1 letters of s by 3, we have "dbc".
After shifting the first 2 letters of s by 5, we have "igc".
After shifting the first 3 letters of s by 9, we have "rpl", the answer.
Example 2:

Input: s = "aaa", shifts = [1,2,3]
Output: "gfd"

Constraints:

1 <= s.length <= 105
s consists of lowercase English letters.
shifts.length == s.length
0 <= shifts[i] <= 109
*/
func shiftingLetters(s string, shifts []int) string {
	n := len(shifts)
	sums := make([]int, n)

	sums[n-1] = shifts[n-1]
	for i := n - 2; i >= 0; i-- {
		sums[i] = sums[i+1] + shifts[i]
	}

	result := make([]rune, n)
	for i := 0; i < n; i++ {
		newRune := int(s[i]) + sums[i]
		newRune = 'a' + (newRune-'a')%('z'-'a'+1)

		result[i] = rune(newRune)
	}

	return string(result)
}

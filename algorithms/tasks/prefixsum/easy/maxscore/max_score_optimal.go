package maxscore

// The score after splitting a string is the
// zeros in the left substring + ones in the right substring.
//
// 2 <= s.length <= 500
// The string s consists of characters '0' and '1' only.

// Input: s = "011101"
// Output: 5

// "011101"
// ls0 111122
// rs1 112344
// "0.11101" 1 (ls0[0]) + 4(rs1[6-2-0])
// "01.1101" 1 (ls0[1]) + 3(rs1[6-2-1])
// "011.101" 1 (ls0[2]) + 2(rs1[6-2-2])
// "0111.01" 1 (ls0[3]) + 1(rs1[6-2-3])
// "01110.1" 2 (ls0[4]) + 1(rs1[6-2-4])

// Explanation:
// All possible ways of splitting s into two non-empty substrings are:
// left = "0" and right = "11101", score = 1 + 4 = 5
// left = "01" and right = "1101", score = 1 + 3 = 4
// left = "011" and right = "101", score = 1 + 2 = 3
// left = "0111" and right = "01", score = 1 + 1 = 2
// left = "01110" and right = "1", score = 2 + 1 = 3
// Example 2:

// Input: s = "00111"
// Output: 5
// Explanation: When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5
// Example 3:
//

// "00111"
// ls0 12222
// rs1 12333
// "0.0111" 1 (ls0[0]) + 3(rs1[5-2-0])
// "00.111" 2 (ls0[1]) + 3(rs1[5-2-1])
// "001.11" 2 (ls0[2]) + 2(rs1[5-2-2])
// "0011.1" 2 (ls0[3]) + 1(rs1[5-2-3])

// Input: s = "1111"
// Output: 3

// ls0 0000
// rs1 1234
// "1.111" 0 + 3 = ls0[0] + rs[n-1-1-0]
// "11.11" 0 + 2 = ls0[1] + rs[n-1-1-1]
// "111.1" 0 + 1 = ls0[2] + rs[n-1-1-2]

// return the maximum score after splitting the string into two non-empty substrings
// "011101"
// len = 6
// i   012345
// ls0 111122
// rs1 112344
// "0.11101" 1 (ls0[0]) + 4(rs1[6-2-0])
// "01.1101" 1 (ls0[1]) + 3(rs1[6-2-1])
// "011.101" 1 (ls0[2]) + 2(rs1[6-2-2])
// "0111.01" 1 (ls0[3]) + 1(rs1[6-2-3])
// "01110.1" 2 (ls0[4]) + 1(rs1[6-2-4])

// "01"
// i = 0
// lsc0 [1]
// i = 0
// rsc1 [1]
// maxsc = 2
// i := 1; i < 1

// 001
// i 0, l0 [1 0], r1 [1 0]
// i 1, i < 2, l0 [1 2], r1 [1 1]
// i 1, i < 2, max(2, 2 + 1) = 3

// 1111, n 4
// i 0, l0 [0 0 0 0], r1 [1 0 0 0]
// i 1, l0 [0 0 0 0], r1 [1 2 0 0]
// i 2, l0 [0 0 0 0], r1 [1 2 3 0]
// i 3 stop

func maxScoreOptimal(s string) int {
	r1 := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			r1++
		}
	}

	l0 := 0
	if s[0] == '0' {
		l0++
	} else {
		r1--
	}
	maxScore := l0 + r1

	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			l0++
		} else {
			r1--
		}
		maxScore = max(maxScore, l0+r1)
	}

	return maxScore
}

// "01", n 2
// r1 1
// i 0, l0 1, r1 1
// 1 < 1

// "111", n 3
// r1 3
// i 0, l0 0, r1 2
// maxScore = 2
// i 1 l0 0, r1 1
// maxScore = 2
// i 2, i < 2
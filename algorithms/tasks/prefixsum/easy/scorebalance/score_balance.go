package scorebalance

/*
'a' = 1, 'b' = 2, ..., 'z' = 26.

Example 1:

Input: s = "adcb"
Output: true

Explanation:
Split at index i = 1:
Left substring = s[0..1] = "ad" with score = 1 + 4 = 5
Right substring = s[2..3] = "cb" with score = 3 + 2 = 5
Both substrings have equal scores, so the output is true.

Example 2:
Input: s = "bace"
Output: false
Explanation: No split produces equal scores, so the output is false.

2 <= s.length <= 100
s consists of lowercase English letters.
*/
func scoreBalance(s string) bool {
	sums := make([]int, len(s))

	sums[0] = int(s[0]) - 'a' + 1
	for i := 1; i < len(s); i++ {
		// d 100 4
		// b 98 2
		// c 99 3
		// fmt.Println(string(s[i]), s[i], int(s[i])-'a'+1)
		sums[i] = sums[i-1] + int(s[i]) - 'a' + 1
	}

	for i := 0; i < len(sums); i++ {
		if sums[i] == sums[len(sums)-1]-sums[i] {
			return true
		}
	}

	return false
}

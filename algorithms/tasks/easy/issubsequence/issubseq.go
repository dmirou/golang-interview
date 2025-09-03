package issubsequence

import "strings"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	si := 0

	for ti := 0; ti < len(t); ti++ {
		if s[si] == t[ti] {
			si++
		}
		if si == len(s) {
			return true
		}
	}

	return false
}

func isSubsequenceAnotherSolution(sub string, of string) bool {
	var ofStart int
	for _, r := range sub {
		idx := strings.IndexRune(of[ofStart:], r)
		if idx == -1 {
			return false
		}
		ofStart += idx + 1
	}
	return true
}

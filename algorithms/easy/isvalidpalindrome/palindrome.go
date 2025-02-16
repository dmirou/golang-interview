package isvalidpalindrome

func isPalindrome(s string) bool {
	direct := make([]rune, 0, len(s))
	reverse := make([]rune, 0, len(s))
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		r := rune(s[i])
		if isAlphaNum(r) {
			direct = append(direct, toLower(r))
		}

		r = rune(s[sLen-1-i])
		if isAlphaNum(r) {
			reverse = append(reverse, toLower(r))
		}
	}

	for i := 0; i < len(direct); i++ {
		if direct[i] != reverse[i] {
			return false
		}
	}

	return true
}

func isAlphaNum(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}

	return false
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return 'a' + r - 'A'
	}

	return r
}

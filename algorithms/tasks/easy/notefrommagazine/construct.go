package notefrommagazine

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)

	for _, r := range magazine {
		letters[r]++
	}

	for _, r := range ransomNote {
		if count, ok := letters[r]; ok && count >= 1 {
			letters[r]--
			continue
		}
		return false
	}

	return true
}

package isanagrams

func isAnagram(s1 string, s2 string) bool {
	m1 := make(map[rune]int)

	for _, r := range s1 {
		m1[r]++
	}

	for _, r := range s2 {
		if _, ok := m1[r]; !ok {
			return false
		}

		m1[r]--
		if m1[r] == 0 {
			delete(m1, r)
		}
	}

	return len(m1) == 0
}

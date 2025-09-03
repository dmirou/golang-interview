package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	m := make(map[rune]rune)
	used := make(map[rune]struct{})

	for i := 0; i < len(s); i++ {
		r1 := rune(s[i])
		r2 := rune(t[i])

		want, ok := m[r1]
		if ok && want != r2 {
			return false
		}
		if ok && want == r2 {
			continue
		}
		if _, ok := used[r2]; ok {
			return false
		}

		m[r1] = r2
		used[r2] = struct{}{}
	}

	return true
}

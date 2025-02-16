package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")

	if len(strs) != len(pattern) {
		return false
	}

	m := make(map[string]byte)
	used := make(map[byte]struct{})

	for i := 0; i < len(strs); i++ {
		si := strs[i]
		bi := pattern[i]

		want, ok := m[si]
		if ok && want != bi {
			return false
		}
		if ok && want == bi {
			continue
		}
		if _, ok := used[bi]; ok {
			return false
		}

		m[si] = bi
		used[bi] = struct{}{}
	}

	return true
}

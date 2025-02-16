package longestcommonprefix

// flower flow flight
// rune = strs[0][0], rune = f
// if strs[1][0] != rune; return
// if strs[2][0] != rune; return
// prefix += rune; prefix = f
// rune = strs[0][1], rune = l
// if strs[1][1] != rune; return
// if strs[2][1] != rune; return
// prefix += rune; prefix = fl
// rune = strs[0][2], rune = o
// if strs[1][2] != rune; return
// if strs[2][2] != rune; return prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := ""

	var prefixRune rune

	for ci := 0; ci < len(strs[0]); ci++ {
		for si := 0; si < len(strs); si++ {
			if si == 0 {
				prefixRune = rune(strs[si][ci])
				continue
			}
			if ci == len(strs[si]) {
				return prefix
			}
			if prefixRune != rune(strs[si][ci]) {
				return prefix
			}
		}
		prefix = prefix + string(prefixRune)
	}

	return prefix
}

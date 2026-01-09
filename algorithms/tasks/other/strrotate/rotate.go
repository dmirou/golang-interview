package strrotate

import (
	"unicode/utf8"
)

func Rotate(s string) string {
	rc := utf8.RuneCountInString(s)
	res := make([]rune, rc)

	i := rc - 1
	for _, r := range s {
		res[i] = r
		i--
	}

	return string(res)
}

func Rotate2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

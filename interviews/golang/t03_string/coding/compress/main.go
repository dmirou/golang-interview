package main

import "fmt"

// "AAABBCDDD" -> "A3B2C1D3
func compress(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""

	prev := rune(-1)
	count := 0

	for _, r := range s {
		if r == prev {
			count++
			continue
		}
		if prev != rune(-1) {
			result += fmt.Sprintf("%c%d", prev, count)
		}
		prev = r
		count = 1
	}
	if count > 0 {
		result += fmt.Sprintf("%c%d", prev, count)
	}

	return result
}

func main() {
	fmt.Println(compress("AAABBCDDD"))
}

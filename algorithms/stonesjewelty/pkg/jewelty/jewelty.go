package jewelty

func Count(j, s string) int {
	count := 0
	for _, jR := range j {
		for _, sR := range s {
			if sR == jR {
				count++
			}
		}
	}

	return count
}

func CountFast(j, s string) int {
	jRunes := make(map[rune]struct{})

	count := 0
	for _, jR := range j {
		jRunes[jR] = struct{}{}
	}

	for _, sR := range s {
		if _, ok := jRunes[sR]; ok {
			count++
		}
	}

	return count
}

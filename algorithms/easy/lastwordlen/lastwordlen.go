package lastwordlen

func lengthOfLastWord(s string) int {
	lastWordLen := 0
	curLen := 0

	for _, r := range s {
		if r != ' ' {
			curLen++
			continue
		}
		if curLen != 0 {
			lastWordLen = curLen
		}
		curLen = 0
	}

	if curLen != 0 {
		lastWordLen = curLen
	}

	return lastWordLen
}

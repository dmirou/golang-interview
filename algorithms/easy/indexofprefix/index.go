package indexofprefix

// haystack = "asadf", needle = "sad", return 1
func strStr(haystack string, needle string) int {
	nlen := len(needle)

	for i := 0; i <= len(haystack)-nlen; i++ {
		if needle == haystack[i:i+nlen] {
			return i
		}
	}

	return -1
}

package validparentheses

func isValid(s string) bool {
	st := make([]rune, 0, len(s)/2)

	for _, r := range s {
		switch r {
		case '[', '{', '(':
			st = append(st, r)
			continue
		}

		if len(st) == 0 {
			return false
		}

		lasti := len(st) - 1
		top := st[lasti]
		st = st[:lasti]
		switch {
		case top == '[' && r == ']',
			top == '{' && r == '}',
			top == '(' && r == ')':
			continue
		}
		return false
	}

	return len(st) == 0
}

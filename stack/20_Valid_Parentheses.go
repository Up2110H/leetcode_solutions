package stack

func isValid(s string) bool {
	stack := make([]rune, 0)
	sLen := 0

	for _, r := range s {
		check := r == ')' && (sLen == 0 || stack[sLen-1] != '(')
		check = check || r == '}' && (sLen == 0 || stack[sLen-1] != '{')
		check = check || r == ']' && (sLen == 0 || stack[sLen-1] != '[')

		if r == ')' || r == '}' || r == ']' {
			if sLen == 0 || stack[sLen-1] != giveOpenP(r) {
				return false
			}
			sLen--
			stack = stack[:sLen]
		} else {
			stack = append(stack, r)
			sLen++
		}
	}

	return sLen == 0
}

func giveOpenP(r rune) rune {
	switch r {
	case ')':
		return '('
	case '}':
		return '{'
	default:
		return '['
	}
}

/* 18:02 11.02.2022 - 18:15 11.02.2022 */

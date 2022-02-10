package stack

func canBeValid(s string, locked string) bool {
	stack, rotA, rotB, rA, rB, goodPos := 0, 0, 0, 0, 0, -1

	for i, c := range s {
		if locked[i] == '0' {
			if c == '(' {
				rotA++
			} else {
				rotB++
			}
		}

		if c == '(' {
			if locked[i] == '0' && stack > 0 {
				rotA--
				rotB++
				stack--
			} else {
				stack++
			}
		} else {
			if stack == 0 {
				if rotB == 0 {
					return false
				} else {
					rotB--
					rotA++
					stack++
				}
			} else {
				stack--
			}

			if stack == 0 {
				goodPos = i
			}
		}
	}

	if stack == 0 {
		return true
	}

	stack = 0

	for i := len(s) - 1; i > goodPos; i-- {
		if locked[i] == '0' {
			if s[i] == '(' {
				rA++
			} else {
				rB++
			}
		}

		if s[i] == ')' {
			if locked[i] == '0' && stack > 0 {
				rB--
				rA++
				stack--
			} else {
				stack++
			}
		} else {
			if stack == 0 {
				if rA == 0 {
					return false
				} else {
					rA--
					rB++
					stack++
				}
			} else {
				stack--
			}
		}
	}

	if stack == 0 || stack%2 == 0 && stack/2 <= rotB {
		return true
	}

	return false
}

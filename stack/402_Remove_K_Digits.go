package stack

func removeKdigits(num string, k int) string {
	stack := ""

	for i := 0; i < len(num); i++ {
		j := len(stack) - 1
		for k > 0 && j >= 0 && stack[j] > num[i] {
			stack = stack[:j]
			k--
			j--
		}
		stack += string(num[i])
	}

	if k > 0 {
		if len(stack) <= k {
			stack = "0"
		} else {
			stack = stack[:len(stack)-k]
		}
	}

	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:]
	}

	return stack
}

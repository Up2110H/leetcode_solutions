package stack

func scoreOfParentheses(s string) int {
	stack := make([]int, len(s)/2+1)
	counter := 0

	for _, c := range s {
		if c == '(' {
			counter++
		} else {
			if stack[counter] == 0 {
				stack[counter-1] += 1
			} else {
				stack[counter-1] += 2 * stack[counter]
			}

			stack[counter] = 0
			counter--
		}
	}

	return stack[0]
}

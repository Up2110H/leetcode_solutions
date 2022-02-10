package stack

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, (len(tokens)+1)/2)
	for _, c := range tokens {
		last := len(stack) - 1
		switch c {
		case "+":
			stack[last-1] += stack[last]
			stack = stack[:last]
		case "-":
			stack[last-1] -= stack[last]
			stack = stack[:last]
		case "/":
			stack[last-1] /= stack[last]
			stack = stack[:last]
		case "*":
			stack[last-1] *= stack[last]
			stack = stack[:last]
		default:
			s, _ := strconv.Atoi(c)
			stack = append(stack, s)
		}
	}

	return stack[0]
}

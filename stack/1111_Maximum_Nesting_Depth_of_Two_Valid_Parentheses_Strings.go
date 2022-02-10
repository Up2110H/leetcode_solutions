package stack

func maxDepthAfterSplit(seq string) []int {
	answer := make([]int, len(seq))
	stack := make([]int, 0, len(seq)/2)

	for i, c := range seq {
		if c == '(' {
			stack = append(stack, i)
		} else {
			if len(stack)%2 == 0 {
				answer[i] = 1
				answer[stack[len(stack)-1]] = 1
			}
			stack = stack[:len(stack)-1]
		}
	}

	return answer
}

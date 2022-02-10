package stack

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	answer := make([]int, len(temperatures))

	for i := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			tail := stack[len(stack)-1]
			answer[tail] = i - tail
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return answer
}

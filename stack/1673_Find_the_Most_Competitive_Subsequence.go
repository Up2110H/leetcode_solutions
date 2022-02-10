package stack

func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0, k)

	for i, v := range nums {
		count := len(stack) + len(nums) - i - k
		for len(stack) > 0 && count > 0 && stack[len(stack)-1] > v {
			stack = stack[:len(stack)-1]
			count--
		}
		if len(stack) < k {
			stack = append(stack, v)
		}
	}

	return stack
}

package stack

type letters struct {
	letter rune
	count  int
}

func removeDuplicates(s string, k int) string {
	stack := make([]letters, 1, len(s)+1)
	for _, c := range s {
		tail := len(stack) - 1
		if len(stack) > 1 && stack[tail].letter == c && stack[tail].count == k-1 {
			for stack[tail].letter == c {
				stack = stack[:tail]
				tail--
			}
			continue
		}
		counter := 1
		if stack[tail].letter == c {
			counter = stack[tail].count + 1
		}
		stack = append(stack, letters{c, counter})
	}
	answer := ""
	for i := 1; i < len(stack); i++ {
		answer += string(stack[i].letter)
	}
	return answer
}

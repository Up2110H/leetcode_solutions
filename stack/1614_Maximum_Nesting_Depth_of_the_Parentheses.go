package stack

func maxDepth(s string) int {
	stack := 0
	depth := 0
	for _, c := range s {
		if c == '(' {
			stack++
			if depth < stack {
				depth = stack
			}
		} else if c == ')' {
			stack--
		}
	}

	return depth
}

package stack

import "strings"

func isValidSerialization(preorder string) bool {

	if len(preorder) == 1 && preorder[0] == '#' {
		return true
	}

	split := strings.Split(preorder, ",")

	stack := make([]bool, 0, len(split))

	for i, c := range split {
		if c == "#" {
			if len(stack) == 0 {
				return false
			}

			for len(stack) > 0 && stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}

			if len(stack) > 0 {
				stack[len(stack)-1] = true
			} else if i < len(split)-1 {
				return false
			}

			continue
		}

		stack = append(stack, false)
	}

	return len(stack) == 0
}

package stack

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	stack := make([]int, 0, n)
	mins := make([]int, len(nums))
	min := math.MaxInt32

	for i, v := range nums {
		if v < min {
			min = v
		}

		mins[i] = min

		j := len(stack) - 1
		for len(stack) > 0 && nums[stack[j]] <= v {
			stack = stack[:j]
			j--
		}

		if j >= 0 && mins[stack[j]] < v {
			return true
		}

		stack = append(stack, i)
	}

	return false
}

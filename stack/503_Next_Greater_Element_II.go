package stack

import "math"

func nextGreaterElements(nums []int) []int {
	answers := make([]int, len(nums))
	maxPos, maxValue := 0, math.MinInt32
	stack := make([]int, 0, len(nums))

	for i := range nums {
		if nums[i] >= maxValue {
			maxValue = nums[i]
			maxPos = i
		}
	}

	for i := 1; i <= len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[(maxPos+i)%len(nums)] {
			answers[stack[len(stack)-1]] = nums[(maxPos+i)%len(nums)]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, (maxPos+i)%len(nums))
	}

	for i := range stack {
		answers[stack[i]] = -1
	}

	return answers
}

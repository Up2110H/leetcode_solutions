package stack

func maxWidthRamp(nums []int) int {
	stack := []int{0}
	res := 0

	for i, v := range nums {
		if nums[stack[len(stack)-1]] > v {
			stack = append(stack, i)
		}
	}

	for i := len(nums) - 1; i >= 0 && len(stack) > 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			if i-stack[len(stack)-1] > res {
				res = i - stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}

	return res
}

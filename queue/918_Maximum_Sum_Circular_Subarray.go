package queue

import "math"

func maxSubarraySumCircular(nums []int) int {
	S := 0

	for _, v := range nums {
		S += v
	}

	return max(
		kadane(nums, 0, len(nums), 1),
		max(
			S+kadane(nums, 0, len(nums)-1, -1),
			S+kadane(nums, 1, len(nums), -1),
		),
	)
}

func kadane(nums []int, l, r, sign int) int {
	s := 0
	ans := math.MinInt32

	for i := l; i < r; i++ {
		s = sign*nums[i] + max(s, 0)
		ans = max(s, ans)
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

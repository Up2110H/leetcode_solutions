package stack

func maxSumMinProduct(nums []int) int {
	n := len(nums)
	prefixSum := make([]uint64, n+1)
	var sum, ans uint64
	left, right, stack := make([]int, n), make([]int, n), make([]int, 0, n)

	for i, v := range nums {
		sum += uint64(v)
		prefixSum[i+1] = sum
	}

	for i, v := range nums {
		last := len(stack) - 1
		for last >= 0 && nums[stack[last]] >= v {
			stack = stack[:last]
			last--
		}
		left[i] = i

		if last >= 0 {
			left[i] -= stack[last] + 1
		}

		stack = append(stack, i)
	}

	stack = make([]int, 0, n)

	for i := n - 1; i >= 0; i-- {
		last := len(stack) - 1
		for last >= 0 && nums[stack[last]] >= nums[i] {
			stack = stack[:last]
			last--
		}
		right[i] = n - 1 - i

		if last >= 0 {
			right[i] -= n - stack[last]
		}

		stack = append(stack, i)
	}

	for i, v := range nums {
		sum = (prefixSum[i+right[i]+1] - prefixSum[i-left[i]]) * uint64(v)
		if sum > ans {
			ans = sum
		}
	}

	return int(ans % 1000000007)
}

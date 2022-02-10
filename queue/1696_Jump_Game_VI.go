package queue

func maxResult(nums []int, k int) int {
	queue := make([]int, 0, k)
	ans := nums[0]
	curr := 0

	for i := 1; i < len(nums); i++ {
		j := len(queue)
		for j > 0 && nums[queue[j-1]] <= nums[i] {
			queue = queue[:j-1]
			j--
		}

		queue = append(queue, i)

		if i-curr == k || nums[i] >= 0 {
			ans += nums[queue[0]]
			curr = queue[0]
			queue = queue[1:]
		}
	}

	if len(queue) > 0 {
		ans += queue[len(queue)-1]
	}

	return ans
}

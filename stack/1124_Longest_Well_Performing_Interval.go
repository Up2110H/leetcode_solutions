package stack

func longestWPI(hours []int) int {
	answer := 0
	sum := 0
	mp := make(map[int]int)

	for i, v := range hours {
		sum += v/9 - (v/9+1)%2

		if sum > 0 {
			answer = max(answer, i+1)
		} else if _, ok := mp[sum]; !ok {
			mp[sum] = i
		}

		if mv, ok := mp[sum-1]; ok {
			answer = max(answer, i-mv)
		}
	}

	return answer
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

package array

func shuffle(nums []int, n int) []int {
	maxN := 1000
	for i := 0; i < 2*n; i++ {
		nums[i]--
	}
	for i := 0; i < n; i++ {
		nums[2*i] += maxN * (nums[i] % maxN)
		nums[2*i+1] += maxN * (nums[n+i] % maxN)
	}
	for i := 0; i < 2*n; i++ {
		nums[i] /= maxN
		nums[i]++
	}

	return nums
}

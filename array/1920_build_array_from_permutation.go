package array

func buildArray(nums []int) []int {
	maxN := 1000
	for i := range nums {
		nums[i] += maxN * (nums[nums[i]] % maxN)
	}
	for i := range nums {
		nums[i] /= maxN
	}

	return nums
}

package array

func getConcatenation(nums []int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		nums = append(nums, nums[i])
	}

	return nums
}

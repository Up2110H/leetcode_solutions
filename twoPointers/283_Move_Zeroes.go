package twoPointers

func moveZeroes(nums []int) {
	n := len(nums)
	k := 0

	for i, v := range nums {
		if v == 0 {
			k++
		} else {
			nums[i-k] = v
		}
	}

	for i := 0; i < k; i++ {
		nums[n-i-1] = 0
	}
}

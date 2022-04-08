package bitManipulation

func singleNumber(nums []int) []int {
	xor := 0
	ans := make([]int, 2)

	for _, v := range nums {
		xor ^= v
	}

	xor &= -xor

	for _, v := range nums {
		if xor&v == 0 {
			ans[0] ^= v
		} else {
			ans[1] ^= v
		}
	}

	return ans
}

package bitManipulation

func minFlips(a int, b int, c int) int {
	ans := 0

	for i := 0; i < 32; i++ {
		ans += ((c+1)&1)*(a&1+b&1) + (c&1)*(((a|b)+1)&1)
		a >>= 1
		b >>= 1
		c >>= 1
	}

	return ans
}

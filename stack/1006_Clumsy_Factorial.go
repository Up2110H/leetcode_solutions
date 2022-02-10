package stack

func clumsy(n int) int {
	a := n % 4
	b := a + (a-1)*(a-2)/2*(a-7*(a-3)/3)
	if n <= 4 {
		return b
	}
	return n + a + 2 - b/7 - b%7
}

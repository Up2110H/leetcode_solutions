package stack

func minSwaps(s string) int {
	counter := 0
	result := 0

	for i, j := 0, len(s)-1; i < j; i++ {
		if s[i] == '[' {
			counter++
		} else if counter != 0 {
			counter--
		} else {
			for s[j] != '[' {
				j--
			}
			counter++
			result++
		}
	}

	return result
}

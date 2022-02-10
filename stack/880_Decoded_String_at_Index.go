package stack

import "unicode"

func decodeAtIndex(s string, k int) string {
	count := 0
	var preC rune

	for _, c := range s {
		if unicode.IsDigit(c) {
			newCount := count * int(c-'0')

			if k%count == 0 && k <= newCount {
				return string(preC)
			} else if newCount < k {
				count = newCount
			} else {
				return decodeAtIndex(s, k%count)
			}
		} else {
			count++
			if count == k {
				return string(c)
			}
			preC = c
		}
	}

	return ""
}

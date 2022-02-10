package stack

func minAddToMakeValid(s string) int {
	counter, result := 0, 0
	for i := range s {
		if s[i] == '(' {
			counter++
		} else {
			if counter == 0 {
				result++
			} else {
				counter--
			}
		}
	}

	return counter + result
}

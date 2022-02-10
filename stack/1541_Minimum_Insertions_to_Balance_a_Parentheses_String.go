package stack

func minInsertions(s string) int {
	counter := 0
	extra := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			counter++
		} else {
			if i != len(s)-1 && s[i+1] == ')' {
				i++
			} else {
				extra++
			}
			if counter > 0 {
				counter--
			} else {
				extra++
			}
		}
	}

	return extra + counter*2
}

package string

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	ans := ""

	for i := 0; i < numRows; i++ {
		k := numRows - 1 - i
		for j := i; j < n; j, k = j+2*k, numRows-1-k {
			if k == 0 {
				continue
			}
			ans += string(s[j])
		}
	}

	return ans
}

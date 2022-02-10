package string

import "math"

func longestPalindrome(s string) string {
	s2 := "#"
	for _, v := range s {
		s2 += string(v) + "#"
	}

	p := make([]int, len(s2), len(s2))
	c, r, maxlength, index := 0, 0, 0, 0

	for i := 1; i < len(s2)-1; i++ {
		i_mir := 2*c - i

		if r > i {
			p[i] = int(math.Min(float64(p[i_mir]), float64(r-i)))
		}

		for i > p[i] && i+p[i] < len(s2)-1 && s2[i-p[i]-1] == s2[i+p[i]+1] {
			p[i]++
		}

		if r < p[i] {
			c = i
			r = i + p[i]
		}

		if maxlength < p[i] {
			maxlength = p[i]
			index = i
		}
	}

	s3 := ""

	for i := index - maxlength; i <= index+maxlength; i++ {
		if s2[i] != '#' {
			s3 += string(s2[i])
		}
	}

	return s3
}

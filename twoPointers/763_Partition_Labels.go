package twoPointers

func partitionLabels(s string) []int {
	counter := make([]int, 26)
	ans := make([]int, 0)

	for i, c := range s {
		counter[c-'a'] = i
	}

	for i := 0; i < len(s); i++ {
		l, r := i, counter[s[i]-'a']

		for j := l; j < r; j++ {
			r = max(r, counter[s[j]-'a'])
		}

		ans = append(ans, r-l+1)
		i = r
	}

	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

/* 20:26 11.02.2022 - 20:48 11.02.2022 */

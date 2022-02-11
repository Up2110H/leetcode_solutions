package list

func guessNumber(n int) int {
	l, r := 1, n+1

	for {
		k := (l + r) / 2
		if j := guess(k); j == 0 {
			return k
		} else if j == 1 {
			l = k + 1
		} else {
			r = k
		}
	}
}

func guess(num int) int {
	a := 1
	if num == a {
		return 0
	} else if num < a {
		return 1
	} else {
		return -1
	}
}

/* 17:53 11.02.2022 - 18:01 11.02.2022 */

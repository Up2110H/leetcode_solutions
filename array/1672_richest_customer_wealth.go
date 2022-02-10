package array

func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for _, c := range accounts {
		sumWealth := 0
		for _, w := range c {
			sumWealth += w
		}
		if maxWealth < sumWealth {
			maxWealth = sumWealth
		}
	}

	return maxWealth
}

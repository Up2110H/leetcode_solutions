package array

func mostWordsFound(sentences []string) int {
	maxJ := 0
	for i := range sentences {
		j := 1
		for _, c := range sentences[i] {
			if c == ' ' {
				j++
			}
		}

		if maxJ < j {
			maxJ = j
		}
	}

	return maxJ
}

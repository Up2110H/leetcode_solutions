package hashTable

func isAnagram(s string, t string) bool {
	hashTable := make(map[rune]int)
	c := 0

	for _, v := range s {
		if _, ok := hashTable[v]; ok {
			hashTable[v]++
		} else {
			c++
			hashTable[v] = 1
		}
	}

	for _, v := range t {
		if _, ok := hashTable[v]; ok {
			hashTable[v]--

			if hashTable[v] == 0 {
				delete(hashTable, v)
				c--
			}
		} else {
			return false
		}
	}

	return c == 0
}

/* 22:57 11.02.2022 - 23:03 11.02.2022 */

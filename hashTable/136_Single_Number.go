package hashTable

/*func singleNumber(nums []int) int {
	hashTable := make(map[int]bool)

	for _, v := range nums {
		if _, ok := hashTable[v]; ok {
			delete(hashTable, v)
		} else {
			hashTable[v] = true
		}
	}

	for ans := range hashTable {
		return ans
	}

	return 0
}*/

func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}

	return a
}

/* 22:22 11.02.2022 - 22:30 11.02.2022 */
/* 22:30 11.02.2022 - 22:34 11.02.2022 */

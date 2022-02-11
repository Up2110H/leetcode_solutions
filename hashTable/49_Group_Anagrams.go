package hashTable

import "sort"

type myStr []rune

func (m myStr) Len() int {
	return len(m)
}

func (m myStr) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m myStr) Less(i, j int) bool {
	return m[i] < m[j]
}

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	hashTable := make(map[string][]string)

	for i, v := range strs {
		r := []rune(v)
		sort.Sort(myStr(r))
		v = string(r)

		if _, ok := hashTable[v]; ok {
			hashTable[v] = append(hashTable[v], strs[i])
		} else {
			hashTable[v] = []string{strs[i]}
		}
	}

	for _, v := range hashTable {
		ans = append(ans, v)
	}

	return ans
}

/* 22:35 11.02.2022 - 22:55 11.02.2022 */

package string

import "sort"

type pair struct {
	word  string
	count int
}

type pairList []pair

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairList) Less(i, j int) bool {
	if p[i].count != p[j].count {
		return p[i].count > p[j].count
	} else {
		return p[i].word < p[j].word
	}
}

func topKFrequent(words []string, k int) []string {
	hashTable := make(map[string]int)

	for _, s := range words {
		if _, ok := hashTable[s]; ok {
			hashTable[s]++
		} else {
			hashTable[s] = 0
		}
	}

	uWords := make(pairList, 0)

	for s, v := range hashTable {
		uWords = append(uWords, pair{s, v})
	}

	sort.Sort(uWords)
	ans := make([]string, k)

	for i := 0; i < k; i++ {
		ans[i] = uWords[i].word
	}

	return ans
}

/* 19:10 11.02.2022 - 19:31 11.02.2022 (1 WA) */

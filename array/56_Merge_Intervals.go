package array

import "sort"

type pairList [][]int

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairList) Less(i, j int) bool {
	if p[i][0] != p[j][0] {
		return p[i][0] < p[j][0]
	} else {
		return p[i][1] < p[j][1]
	}
}

func merge(intervals [][]int) [][]int {
	sort.Sort(pairList(intervals))

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] >= intervals[i][0] {
			intervals[i-1][1] = max(intervals[i-1][1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
		}
	}

	return intervals
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	return i + j - min(i, j)
}

/* 18:42 11.02.2022 - 19:03 11.02.2022 */

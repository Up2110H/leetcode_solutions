package array

import "sort"

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (v PairList) Len() int {
	return len(v)
}

func (v PairList) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v PairList) Less(i, j int) bool {
	return v[i].Key < v[j].Key
}

func twoSum(nums []int, target int) []int {
	numsPair := make([]Pair, len(nums))

	for key, value := range nums {
		numsPair[key] = Pair{value, key}
	}

	sort.Sort(PairList(numsPair))

	i, j := 0, len(nums)-1

	for i < j {
		if numsPair[i].Key+numsPair[j].Key == target {
			break
		}

		if numsPair[i].Key+numsPair[j].Key > target {
			j--
		} else {
			i++
		}
	}

	return []int{numsPair[i].Value, numsPair[j].Value}
}

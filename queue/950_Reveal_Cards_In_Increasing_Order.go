package queue

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(deck)))
	n := len(deck)
	queue := make([]int, n)
	head := 0
	length := 0

	for _, v := range deck {
		if length > 1 {
			queue[(head+length)%n] = queue[head]
			head = (head + 1) % n
		}

		queue[(head+length)%n] = v
		length++
	}

	for i := 0; i < n; i++ {
		deck[n-1-i] = queue[(head+i)%n]
	}

	return deck
}

package stack

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	var stack *ListNode = nil
	result := 0
	it1, it2 := head, head

	for it2 != nil {
		stack = &ListNode{it1.Val, stack}
		it1 = it1.Next
		it2 = it2.Next.Next
	}

	for it1 != nil {
		if stack.Val += it1.Val; stack.Val > result {
			result = stack.Val
		}
		stack = stack.Next
		it1 = it1.Next
	}

	return result
}

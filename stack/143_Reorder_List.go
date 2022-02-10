package stack

func reorderList(head *ListNode) {
	stack := make([]*ListNode, 0)
	for i := head; i != nil; i = i.Next {
		stack = append(stack, i)
	}

	for i, j := 0, len(stack)-1; i < j-1; i, j = i+1, j-1 {
		stack[j].Next = stack[i].Next
		stack[i].Next = stack[j]
		stack[j-1].Next = nil
	}
}

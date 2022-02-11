package list

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		for i := 0; i < len(lists)/2; i++ {
			lists[i] = merge(lists[i], lists[i+1])
			lists = append(lists[:i+1], lists[i+2:]...)
		}
	}

	return lists[0]
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	var list *ListNode = nil
	curr := list
	i, j := l1, l2

	for i != nil || j != nil {
		if curr == nil {
			curr = &ListNode{0, nil}
			list = curr
		} else {
			curr.Next = &ListNode{0, nil}
			curr = curr.Next
		}
		if j == nil || i != nil && i.Val < j.Val {
			curr.Val = i.Val
			i = i.Next
		} else {
			curr.Val = j.Val
			j = j.Next
		}
	}

	return list
}

/* 16:20 11.02.2022 - 16:59 11.02.2022 */
/* 17:01 11.02.2022 - 17:22 11.02.2022 */

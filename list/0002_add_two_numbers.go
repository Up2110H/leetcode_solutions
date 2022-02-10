package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 ListNode
	it := &l3
	tmp := 0
	for {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		it.Val = (v1 + v2 + tmp) % 10
		tmp = (v1 + v2 + tmp) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil && tmp == 0 {
			break
		}

		it.Next = new(ListNode)
		it = it.Next
	}

	return &l3
}

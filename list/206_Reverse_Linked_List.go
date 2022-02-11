package list

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil
	for curr != nil {
		head = curr.Next
		curr.Next = prev
		prev = curr
		curr = head
	}
	return prev
}

/* 17:40 11.02.2022 - 17:47 11.02.2022 */
/* 17:47 11.02.2022 - 17:51 11.02.2022 */

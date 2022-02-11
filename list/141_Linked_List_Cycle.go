package list

func hasCycle(head *ListNode) bool {
	m := 1000000
	for head != nil {
		if head.Val == m {
			return true
		}
		head.Val = m
		head = head.Next
	}

	return false
}

/* 17:25 11.02.2022 - 17:37 11.02.2022 */

package list

/*func hasCycle(head *ListNode) bool {
	m := 1000000
	for head != nil {
		if head.Val == m {
			return true
		}
		head.Val = m
		head = head.Next
	}

	return false
}*/

/*func hasCycle(head *ListNode) bool {
	for i := 0; head != nil; i++ {
		if i == 10002 {
			return true
		}
		head = head.Next
	}

	return false
}*/

/*func hasCycle(head *ListNode) bool {
	for i := 0; head != nil; i++ {
		curr := head
		for j := 0; j < i; j++ {
			if curr.Next == nil {
				return false
			} else if curr.Next == head {
				return true
			}
			curr = curr.Next
		}

		head = head.Next
	}

	return false
}*/

func hasCycle(head *ListNode) bool {
	i := head
	var j *ListNode

	if i != nil {
		j = i.Next
	}

	for j != nil {
		if i == j {
			return true
		}
		i = i.Next
		j = j.Next
		if j != nil {
			j = j.Next
		}
	}

	return false
}

/* 17:25 11.02.2022 - 17:37 11.02.2022 */
/* 19:33 11.02.2022 - 19:37 11.02.2022 */
/* 19:39 11.02.2022 - 19:52 11.02.2022 */

package cycle

func hasCycle(head *ListNode) bool {
	result := false

	if head == nil {
		return result
	}

	if head.Next == nil {
		return result
	}

	var next *ListNode = head
	var afterNext *ListNode = head.Next
	for next != nil {
		if next == afterNext {
			result = true
			break
		}

		if next.Next == nil {
			return false
		}

		next = next.Next

		if afterNext.Next == nil {
			return false
		}

		if afterNext.Next.Next == nil {
			return false
		}

		afterNext = afterNext.Next.Next
	}

	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

package reverse

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var curr *ListNode
	var next *ListNode

	curr = head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

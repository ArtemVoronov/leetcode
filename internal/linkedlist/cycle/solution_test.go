package cycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasLoop1(t *testing.T) {
	var head *ListNode = &ListNode{Val: 3}
	(*head).Next = &ListNode{Val: 2}
	(*head).Next.Next = &ListNode{Val: 0}
	(*head).Next.Next.Next = &ListNode{Val: 4, Next: (*head).Next}

	assert.True(t, hasCycle(head))
}

func TestHasLoop2(t *testing.T) {
	var head *ListNode = &ListNode{Val: 1}
	(*head).Next = &ListNode{Val: 2, Next: head}

	assert.True(t, hasCycle(head))
}

func TestHasNotLoop1(t *testing.T) {
	var head *ListNode = &ListNode{Val: 1}

	assert.False(t, hasCycle(head))
}

func TestHasNotLoop2(t *testing.T) {
	var head *ListNode = &ListNode{Val: 1}
	(*head).Next = &ListNode{Val: 2}

	assert.False(t, hasCycle(head))
}

func TestHasNotLoop3(t *testing.T) {
	var head *ListNode = &ListNode{Val: 1}
	var next *ListNode = &ListNode{Val: 2}
	head.Next = next
	for i := 3; i < 10003; i++ {
		next.Next = &ListNode{Val: i}
		next = next.Next
	}

	assert.False(t, hasCycle(head))
}

package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	var input *ListNode = createSortedList(1, 10)
	var expected *ListNode = createReverseSortedList(1, 10)

	var actual *ListNode = reverseList(input)

	assert.True(t, assertEqualLists(expected, actual))
}

func TestCase2(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	expected := &ListNode{Val: 2, Next: &ListNode{Val: 1}}

	var actual *ListNode = reverseList(input)

	assert.True(t, assertEqualLists(expected, actual))
}

func TestEmpty(t *testing.T) {
	var input *ListNode
	var expected *ListNode

	var actual *ListNode = reverseList(input)

	assert.True(t, assertEqualLists(expected, actual))
}

func createSortedList(from int, n int) *ListNode {
	var head *ListNode = &ListNode{Val: from}
	var next *ListNode = &ListNode{Val: from + 1}
	head.Next = next
	for i := from + 2; i < from+n; i++ {
		next.Next = &ListNode{Val: i}
		next = next.Next
	}

	return head
}

func assertEqualLists(expected *ListNode, actual *ListNode) bool {
	curr1 := expected
	curr2 := actual

	for curr1 != nil || curr2 != nil {
		// fmt.Printf("curr1: %v, curr2: %v\n", curr1.Val, curr2.Val)
		if curr1.Val != curr2.Val {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return true
}

func createReverseSortedList(from int, n int) *ListNode {
	var head *ListNode = &ListNode{Val: from + n - 1}
	var next *ListNode = &ListNode{Val: from + n - 2}
	head.Next = next
	for i := from + n - 3; i >= 1; i-- {
		next.Next = &ListNode{Val: i}
		next = next.Next
	}

	return head
}

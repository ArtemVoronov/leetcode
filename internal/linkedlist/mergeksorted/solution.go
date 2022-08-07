package mergeksorted

import (
	"fmt"
	"sort"
)

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	var resultHead *ListNode

	if isNil(lists) || isEmpty(lists) {
		return result
	}

	currMinHead := lists[0]
	currMinIndex := 0
	listsNum := len(lists)
	count := 0
	for count < listsNum { // if all head all nil, then exit curr = nil
		count = 0
		for i, head := range lists {
			if head == nil {
				count += 1
				continue
			}
			if currMinHead == nil {
				currMinHead = head
				currMinIndex = i
			}
			if (*head).Val < (*currMinHead).Val {
				currMinHead = head
				currMinIndex = i
			}
		}

		if count == listsNum {
			break
		}
		// cut head from inner list, add it to result
		next := currMinHead.Next
		currMinHead.Next = nil
		lists[currMinIndex] = next
		if result == nil {
			resultHead = currMinHead
			result = currMinHead
		} else {
			result.Next = currMinHead
			result = result.Next
		}
		currMinHead = nil

	}

	return resultHead
}

func mergeKLists2(lists []*ListNode) *ListNode {
	var all []*ListNode = make([]*ListNode, 0)

	if isNil(lists) || isEmpty(lists) {
		return nil
	}

	for _, head := range lists {
		curr := head
		for curr != nil {
			next := curr.Next
			curr.Next = nil
			all = append(all, curr)
			curr = next
		}
	}

	sort.Slice(all, func(i, j int) bool {
		return (*all[i]).Val < (*all[j]).Val
	})

	result := convert(all)

	return result
}

func convert(lists []*ListNode) *ListNode {
	var head *ListNode
	var curr *ListNode
	for i, l := range lists {
		if i == 0 {
			head = l
			curr = l
		} else {
			curr.Next = l
			curr = curr.Next
		}
	}

	return head
}

func isNil(lists []*ListNode) bool {
	return len(lists) == 0
}

func isEmpty(lists []*ListNode) bool {
	return len(lists) == 1 && lists[0] == nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(vals []int) *ListNode {
	var head *ListNode
	var curr *ListNode
	for i, val := range vals {
		if i == 0 {
			head = &ListNode{
				Val: val,
			}
			curr = head
		} else {
			curr.Next = &ListNode{
				Val: val,
			}
			curr = curr.Next
		}
	}
	return head
}

func print(list *ListNode) {
	curr := list
	for curr != nil {
		// time.Sleep(100 * time.Millisecond)
		fmt.Printf("%v ", (*curr).Val)
		curr = curr.Next
	}
	fmt.Println()
}

func toInts(list *ListNode) []int {
	result := []int{}
	curr := list
	for curr != nil {
		result = append(result, (*curr).Val)
		curr = curr.Next
	}
	return result
}

package mergeksorted

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCase(t *testing.T) {
	lists := []*ListNode{}
	l1 := createList([]int{1, 4, 5})
	l2 := createList([]int{1, 3, 4})
	l3 := createList([]int{2, 6})

	lists = append(lists, l1, l2, l3)

	expected := createList([]int{1, 1, 2, 3, 4, 4, 5, 6})

	actual := mergeKLists(lists)

	fmt.Printf("expected:\n")
	fmt.Printf("%v\n", toInts(expected))
	fmt.Printf("actual:\n")
	fmt.Printf("%v\n", toInts(actual))

	assert.True(t, reflect.DeepEqual(toInts(expected), toInts(actual)))
}

func TestEmpty(t *testing.T) {
	lists := []*ListNode{}
	lists = append(lists, nil)

	var expected *ListNode

	actual := mergeKLists(lists)

	fmt.Printf("expected:\n")
	fmt.Printf("%v\n", toInts(expected))
	fmt.Printf("actual:\n")
	fmt.Printf("%v\n", toInts(actual))

	assert.True(t, reflect.DeepEqual(toInts(expected), toInts(actual)))
}
func TestNil(t *testing.T) {
	var lists []*ListNode

	var expected *ListNode

	actual := mergeKLists(lists)

	fmt.Printf("expected:\n")
	fmt.Printf("%v\n", toInts(expected))
	fmt.Printf("actual:\n")
	fmt.Printf("%v\n", toInts(actual))

	assert.True(t, reflect.DeepEqual(toInts(expected), toInts(actual)))
}

func TestBasicCase2(t *testing.T) {
	lists := []*ListNode{}
	l1 := createList([]int{1, 4, 5})
	l2 := createList([]int{1, 3, 4})
	l3 := createList([]int{2, 6})

	lists = append(lists, l1, l2, l3)

	expected := createList([]int{1, 1, 2, 3, 4, 4, 5, 6})

	actual := mergeKLists2(lists)

	fmt.Printf("expected:\n")
	fmt.Printf("%v\n", toInts(expected))
	fmt.Printf("actual:\n")
	fmt.Printf("%v\n", toInts(actual))

	assert.True(t, reflect.DeepEqual(toInts(expected), toInts(actual)))
}
func TestBig2(t *testing.T) {
	lists := []*ListNode{}
	lists = append(lists, createList([]int{-6, -4, 0, 0, 4}))
	lists = append(lists, createList([]int{}))
	lists = append(lists, createList([]int{-4, -2, -1, 1, 2, 3}))
	lists = append(lists, createList([]int{-9, -6, -5, -2, 4, 4}))
	lists = append(lists, createList([]int{-9, -6, -5, -2, -1, 3}))
	lists = append(lists, createList([]int{-2, -1, 0}))
	lists = append(lists, createList([]int{-6}))
	lists = append(lists, createList([]int{-8, -1, 0, 2}))

	expected := createList([]int{-9, -9, -8, -6, -6, -6, -6, -5, -5, -4, -4, -2, -2, -2, -2, -1, -1, -1, -1, 0, 0, 0, 0, 1, 2, 2, 3, 3, 4, 4, 4})

	actual := mergeKLists2(lists)

	fmt.Printf("expected:\n")
	fmt.Printf("%v\n", toInts(expected))
	fmt.Printf("actual:\n")
	fmt.Printf("%v\n", toInts(actual))

	assert.True(t, reflect.DeepEqual(toInts(expected), toInts(actual)))
}
func TestEmpty2(t *testing.T) {
	lists := []*ListNode{}
	lists = append(lists, nil)

	var expected *ListNode

	actual := mergeKLists2(lists)

	fmt.Printf("expected:\n")
	fmt.Printf("%v\n", toInts(expected))
	fmt.Printf("actual:\n")
	fmt.Printf("%v\n", toInts(actual))

	assert.True(t, reflect.DeepEqual(toInts(expected), toInts(actual)))
}
func TestNil2(t *testing.T) {
	var lists []*ListNode

	var expected *ListNode

	actual := mergeKLists2(lists)

	fmt.Printf("expected:\n")
	fmt.Printf("%v\n", toInts(expected))
	fmt.Printf("actual:\n")
	fmt.Printf("%v\n", toInts(actual))

	assert.True(t, reflect.DeepEqual(toInts(expected), toInts(actual)))
}

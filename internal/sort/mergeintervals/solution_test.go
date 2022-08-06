package mergeintervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCase1(t *testing.T) {
	input := [][]int{}
	input = append(input, []int{1, 3})
	input = append(input, []int{2, 6})
	input = append(input, []int{8, 10})
	input = append(input, []int{15, 18})

	expected := [][]int{}
	expected = append(expected, []int{1, 6})
	expected = append(expected, []int{8, 10})
	expected = append(expected, []int{15, 18})

	actual := merge(input)

	assert.True(t, assertEqualIntervals(expected, actual))
}
func TestBasicCase2(t *testing.T) {
	input := [][]int{}
	input = append(input, []int{1, 4})
	input = append(input, []int{4, 5})

	expected := [][]int{}
	expected = append(expected, []int{1, 5})

	actual := merge(input)

	assert.True(t, assertEqualIntervals(expected, actual))
}
func TestBasicCase3(t *testing.T) {
	input := [][]int{}
	input = append(input, []int{1, 4})
	input = append(input, []int{5, 6})

	expected := [][]int{}
	expected = append(expected, []int{1, 4})
	expected = append(expected, []int{5, 6})

	actual := merge(input)

	assert.True(t, assertEqualIntervals(expected, actual))
}

func assertEqualIntervals(expected [][]int, actual [][]int) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := range actual {
		actualInterval := actual[i]
		expectedInterval := expected[i]
		if len(actualInterval) != 2 {
			return false
		}
		if expectedInterval[0] != actualInterval[0] {
			return false
		}
		if expectedInterval[1] != actualInterval[1] {
			return false
		}
	}

	return true
}

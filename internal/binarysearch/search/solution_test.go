package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseFound(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCaseNotFound(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	expected := -1

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCaseSingleElementFound(t *testing.T) {
	nums := []int{5}
	target := 5
	expected := 0

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCaseSingleElementNotFound(t *testing.T) {
	nums := []int{5}
	target := 2
	expected := -1

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

func TestCaseOutOFMemory(t *testing.T) {
	nums := []int{2, 5}
	target := 0
	expected := -1

	actual := search(nums, target)

	assert.Equal(t, expected, actual)
}

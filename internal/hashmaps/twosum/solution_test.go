package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	actual := twoSum(nums, target)

	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6

	expected := []int{1, 2}
	actual := twoSum(nums, target)

	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	expected := []int{0, 1}
	actual := twoSum(nums, target)

	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected, actual)
}
func TestCase4(t *testing.T) {
	nums := []int{18, 19, 15, 2, 7, 11, 15}
	target := 9

	expected := []int{3, 4}
	actual := twoSum(nums, target)

	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected, actual)
}

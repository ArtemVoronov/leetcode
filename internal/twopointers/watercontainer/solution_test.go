package watercontainer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49

	actual := maxArea(height)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	height := []int{1, 1}
	expected := 1

	actual := maxArea(height)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	height := []int{2, 3, 4, 5, 18, 17, 6}
	expected := 17

	actual := maxArea(height)

	assert.Equal(t, expected, actual)
}

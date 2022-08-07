package mergesort

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCase(t *testing.T) {
	input := []int{5, 8, 4, 1, 2, 6, 3, 7}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	actual := mergesort(input)

	assert.Equal(t, len(expected), len(actual))
	assert.True(t, reflect.DeepEqual(expected, actual))
}

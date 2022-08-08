package stocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5

	actual := maxProfit(prices)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0

	actual := maxProfit(prices)

	assert.Equal(t, expected, actual)
}

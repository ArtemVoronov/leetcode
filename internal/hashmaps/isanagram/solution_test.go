package isanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"

	expected := true
	actual := isAnagram(s1, s2)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	s1 := "rat"
	s2 := "cat"

	expected := false
	actual := isAnagram(s1, s2)

	assert.Equal(t, expected, actual)
}

package groupanagrams

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{}
	g1 := []string{"bat"}
	g2 := []string{"nat", "tan"}
	g3 := []string{"ate", "eat", "tea"}
	expected = append(expected, g1, g2, g3)

	actual := groupAnagrams(input)

	assert.Equal(t, len(expected), len(actual))
	assert.True(t, reflect.DeepEqual(sortAnagramGroups(expected), sortAnagramGroups(actual)))
}

func TestCase2(t *testing.T) {
	input := []string{""}
	expected := [][]string{}
	g1 := []string{""}
	expected = append(expected, g1)

	actual := groupAnagrams(input)

	assert.Equal(t, len(expected), len(actual))
	assert.True(t, reflect.DeepEqual(sortAnagramGroups(expected), sortAnagramGroups(actual)))
}
func TestCase3(t *testing.T) {
	input := []string{"a"}
	expected := [][]string{}
	g1 := []string{"a"}
	expected = append(expected, g1)

	actual := groupAnagrams(input)

	assert.Equal(t, len(expected), len(actual))
	assert.True(t, reflect.DeepEqual(sortAnagramGroups(expected), sortAnagramGroups(actual)))
}

func toString(input []string) string {
	result := ""
	for _, s := range input {
		result += s
	}
	return result
}

func sortAnagramGroups(input [][]string) [][]string {
	for i := range input {
		sort.Strings(input[i])
	}
	sort.Slice(input, func(i, j int) bool {
		s1 := toString(input[i])
		s2 := toString(input[j])
		return s1 > s2
	})
	return input
}

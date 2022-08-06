package groupanagrams

import (
	"sort"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	var anagramsGroups map[string][]string = make(map[string][]string)

	for _, s := range strs {
		hash := getHash2(s)
		group, ok := anagramsGroups[hash]
		if !ok {
			group = make([]string, 0)
		}
		group = append(group, s)
		anagramsGroups[hash] = group
	}

	return convert(anagramsGroups)
}

func getFrequency(s string) map[string]int {
	result := make(map[string]int)

	for _, c := range s {
		result[string(c)] += 1
	}

	return result
}

func getHash1(s string) string {
	result := ""
	m := getFrequency(s)
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		result += k + strconv.Itoa(m[k])
	}
	return result
}

func convert(input map[string][]string) [][]string {
	var result [][]string = [][]string{}
	for _, v := range input {
		result = append(result, v)
	}

	return result
}

func getHash2(s string) string {
	var h [26]byte
	for _, c := range s {
		h[c-'a'] += 1
	}
	return string(h[:])
}

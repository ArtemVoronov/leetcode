package longestcommonprefix

import (
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Slice(strs, func(x, y int) bool {
		return len(strs[x]) < len(strs[y])
	})

	result := strs[0]

	for {
		if result == "" {
			break
		}

		isCommonPrefix := true
		for _, str := range strs {
			if !strings.HasPrefix(str, result) {
				isCommonPrefix = false
			}
		}
		if isCommonPrefix {
			return result
		}
		result = result[:len(result)-1]
	}

	return ""
}

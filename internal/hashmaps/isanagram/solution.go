package isanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return getHash(s) == getHash(t)
}

func getHash(s string) [26]int {
	var h [26]int
	for _, c := range s {
		h[c-'a'] += 1
	}
	return h
}

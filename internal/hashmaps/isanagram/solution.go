package isanagram

func isAnagram(s1 string, s2 string) bool {
	m := make(map[string]int)

	l1 := len(s1)
	l2 := len(s2)

	if l1 != l2 {
		return false
	}

	for _, c := range s1 {
		m[string(c)] += 1
	}

	for _, c := range s2 {
		m[string(c)] -= 1
	}

	for k := range m {
		if m[k] != 0 {
			return false
		}
	}

	return true
}

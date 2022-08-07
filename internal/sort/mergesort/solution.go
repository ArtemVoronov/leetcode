package mergesort

func mergesort(input []int) []int {
	// fmt.Printf("input: %v\n", input)
	l := len(input)

	if l == 0 {
		return []int{}
	}

	if l == 1 {
		return input
	}

	middle := l / 2

	left := input[:middle]
	right := input[middle:]
	sortedLeft := mergesort(left)
	sortedRight := mergesort(right)
	// fmt.Printf("sortedLeft: %v\n", sortedLeft)
	// fmt.Printf("sortedRight: %v\n", sortedRight)
	// fmt.Printf("l: %v\n", l)

	var result []int = make([]int, l, l)

	leftLen := len(sortedLeft)
	rightLen := len(sortedRight)
	i := 0
	j := 0
	for k := 0; k < l; k++ {
		if i >= leftLen {
			result[k] = sortedRight[j]
			j += 1
			// fmt.Printf("k: %v, result: %v\n", k, result)
			continue
		}

		if j >= rightLen {
			result[k] = sortedLeft[i]
			i += 1
			// fmt.Printf("k: %v, result: %v\n", k, result)
			continue
		}

		if sortedLeft[i] <= sortedRight[j] {
			result[k] = sortedLeft[i]
			i += 1
		} else {
			result[k] = sortedRight[j]
			j += 1
		}
		// fmt.Printf("k: %v, result: %v\n", k, result)
	}

	// fmt.Printf("result: %v\n", result)
	return result
}

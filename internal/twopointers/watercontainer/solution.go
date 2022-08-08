package watercontainer

func maxArea(height []int) int {
	result := 0

	r := len(height) - 1
	l := 0

	for l <= r {
		area := getArea(height, l, r)

		if area > result {
			result = area
		}

		if height[l] > height[r] {
			r -= 1
		} else {
			l += 1
		}
	}

	return result
}

func getArea(input []int, i, j int) int {
	width := abs(i - j)
	height := min(input[i], input[j])
	return width * height
}

func abs(number int) int {
	if number < 0 {
		return number * -1
	} else {
		return number
	}
}

func min(i int, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

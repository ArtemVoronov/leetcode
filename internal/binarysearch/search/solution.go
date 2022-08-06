package search

func search(nums []int, target int) int {
	// supposed nums is sorted already, of not then sort.Ints(nums) before

	return binarySearch(&nums, target, 0, len(nums)-1)
}

func binarySearch(nums *[]int, target int, from int, to int) int {
	if from == to {
		if from >= 0 && from < len(*nums) && (*nums)[from] == target {
			return from
		} else {
			return -1
		}
	}

	if from >= len(*nums) || to < 0 {
		return -1
	}

	middle := (from + to) / 2
	middleValue := (*nums)[middle]

	if middleValue == target {
		return middle
	} else if middleValue < target {
		// go to the right
		return binarySearch(nums, target, middle+1, to)
	} else {
		// go to the left
		return binarySearch(nums, target, from, middle-1)
	}
}

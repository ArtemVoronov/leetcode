package medianoftwosortedarray

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// TODO: calc avg of left and right, then compare it
	//cut the half of array where median is lesser/greater
	// TODO: try to split the task to several and solve it separately
	m := merge(nums1, nums2)
	return avg(m)
}

func avg(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		return 0
	}

	middle := l / 2

	if l%2 == 1 {
		return float64(nums[middle])
	}

	return float64(nums[middle-1]+nums[middle]) / 2.0
}

func merge(nums1 []int, nums2 []int) []int {
	i, j := 0, 0
	result := make([]int, 0, len(nums1)+len(nums2))
	for {
		// fmt.Printf("i: %v, j: %v, result: %v\n", i, j, result) // todo clean
		if i >= len(nums1) && j >= len(nums2) {
			break
		}
		if i >= len(nums1) {
			result = append(result, nums2[j])
			j++
			continue
		}
		if j >= len(nums2) {
			result = append(result, nums1[i])
			i++
			continue
		}
		l := nums1[i]
		r := nums2[j]
		if l < r {
			result = append(result, l)
			i++
		} else {
			result = append(result, r)
			j++
		}
	}

	return result
}

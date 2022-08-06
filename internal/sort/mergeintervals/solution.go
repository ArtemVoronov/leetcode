package mergeintervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int = [][]int{}
	var last []int
	for _, interval := range intervals {
		if last != nil {
			if isOverlapping(last, interval) {
				last = mergreIntervals(last, interval)
			} else {
				result = append(result, last)
				last = interval
			}

		} else {
			last = interval
		}
	}

	result = append(result, last)

	return result
}

func isOverlapping(interval1 []int, interval2 []int) bool {
	return interval1[1] >= interval2[0]
}

func mergreIntervals(interval1 []int, interval2 []int) []int {
	var max int
	var min int

	if interval1[0] < interval2[0] {
		min = interval1[0]
	} else {
		min = interval2[0]
	}
	if interval1[1] > interval2[1] {
		max = interval1[1]
	} else {
		max = interval2[1]
	}

	return []int{min, max}
}

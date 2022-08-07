package quicksort

import (
	"math/rand"
	"time"
)

func quicksort(input []int) []int {
	// fmt.Printf("input: %v\n", input)
	sort(input, 0, len(input)-1)
	// fmt.Printf("sorted: %v\n", input)

	return input
}

func sort(input []int, l int, r int) {
	// fmt.Printf("l: %v, r: %v, input: %v\n", l, r, input)
	if l >= r {
		return
	}

	i := choosePivot(input, l, r)
	// fmt.Printf("pivot index: %v, pivot value: %v\n", i, input[i])
	swap(input, l, i) // move pivot element to the start
	// fmt.Printf("after pivoting: %v\n", input)
	j := partition(input, l, r)
	// fmt.Printf("partition: %v\n", input)
	// fmt.Printf("j: %v\n", j)
	sort(input, l, j)
	sort(input, j+1, r)
}

func choosePivot(input []int, l int, r int) int {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(r-l) + l
	return result
}

func partition(input []int, l int, r int) int {
	p := input[l]
	i := l + 1
	for j := l + 1; j <= r; j++ {
		if p > input[j] {
			swap(input, i, j)
			i += 1
		}
	}
	// fmt.Printf("partition before pivo: %v\n", input)
	swap(input, l, i-1) // move pivot element to the correct position
	return i - 1        // index of pivot element
}

func swap(input []int, i int, j int) {
	temp := input[i]
	input[i] = input[j]
	input[j] = temp
}

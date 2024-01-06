package medianoftwosortedarray

import (
	"testing"
)

func TestCase1(t *testing.T) {
	input1 := []int{1, 3}
	input2 := []int{2}

	expected := 2.0
	actual := findMedianSortedArrays(input1, input2)

	if expected != actual {
		t.Errorf("Failed. Expected: %v. Actual: %v\n", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4}

	expected := 2.5
	actual := findMedianSortedArrays(input1, input2)

	if expected != actual {
		t.Errorf("Failed. Expected: %v. Actual: %v\n", expected, actual)
	}
}

func TestAvg1(t *testing.T) {
	input := []int{2}

	expected := 2.0
	actual := avg(input)

	if expected != actual {
		t.Errorf("Failed. Expected: %v. Actual: %v\n", expected, actual)
	}
}

func TestAvg2(t *testing.T) {
	input := []int{1, 2, 3}

	expected := 2.0
	actual := avg(input)

	if expected != actual {
		t.Errorf("Failed. Expected: %v. Actual: %v\n", expected, actual)
	}
}

func TestAvg3(t *testing.T) {
	input := []int{1, 2, 3, 4}

	expected := 2.5
	actual := avg(input)

	if expected != actual {
		t.Errorf("Failed. Expected: %v. Actual: %v\n", expected, actual)
	}
}

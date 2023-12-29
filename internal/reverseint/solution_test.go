package reverseint

import (
	"testing"
)

func TestPositive(t *testing.T) {
	input := 123
	expected := 321

	actual := reverse(input)

	if actual != expected {
		t.Errorf("Failed. Expected: %v. Actual: %v", expected, actual)
	}
}

func TestNegative(t *testing.T) {
	input := -123
	expected := -321

	actual := reverse(input)

	if actual != expected {
		t.Errorf("Failed. Expected: %v. Actual: %v", expected, actual)
	}
}

func TestWithZero(t *testing.T) {
	input := 120
	expected := 21

	actual := reverse(input)

	if actual != expected {
		t.Errorf("Failed. Expected: %v. Actual: %v", expected, actual)
	}
}

func TestZero(t *testing.T) {
	input := 0
	expected := 0

	actual := reverse(input)

	if actual != expected {
		t.Errorf("Failed. Expected: %v. Actual: %v", expected, actual)
	}
}

func TestBig(t *testing.T) {
	input := 1534236469
	expected := 0

	actual := reverse(input)

	if actual != expected {
		t.Errorf("Failed. Expected: %v. Actual: %v", expected, actual)
	}
}

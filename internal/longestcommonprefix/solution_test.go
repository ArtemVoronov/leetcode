package longestcommonprefix

import "testing"

func TestBasic1(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	expected := "fl"

	actual := longestCommonPrefix(input)

	if actual != expected {
		t.Errorf("Failed. Expected: %v. Actual: %v", expected, actual)
	}
}

func TestBasic2(t *testing.T) {
	input := []string{"dog", "racecat", "car"}
	expected := ""

	actual := longestCommonPrefix(input)

	if actual != expected {
		t.Errorf("Failed. Expected: %v. Actual: %v", expected, actual)
	}
}

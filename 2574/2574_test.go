package leetcode2574

import (
	"slices"
	"testing"
)

func TestLeftRightDifference(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{10, 4, 8, 3}, []int{15, 1, 11, 22}},
		{[]int{1}, []int{0}},
	}

	for _, test := range tests {
		if output := leftRightDifference(test.input); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

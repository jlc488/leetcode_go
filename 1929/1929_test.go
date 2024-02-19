package leetcode1929

import (
	"slices"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
	}
	for _, test := range tests {
		if output := getConcatenation(test.input); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

package leetcode2433

import (
	"slices"
	"testing"
)

func TestFindArray(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 0, 3, 1}, []int{5, 7, 2, 3, 2}},
	}

	for _, test := range tests {
		if output := findArray(test.input); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

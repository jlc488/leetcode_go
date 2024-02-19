package leetcode1365

import (
	"slices"
	"testing"
)

func TestSmallerNumbersThanCurrent1(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
		{[]int{7, 7, 7, 7}, []int{0, 0, 0, 0}},
	}

	for _, test := range tests {
		if output := smallerNumbersThanCurrent1(test.input); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestSmallerNumbersThanCurrent2(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
		{[]int{7, 7, 7, 7}, []int{0, 0, 0, 0}},
	}

	for _, test := range tests {
		if output := smallerNumbersThanCurrent1(test.input); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

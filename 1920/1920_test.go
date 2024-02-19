package leetcode1920

import (
	"slices"
	"testing"
)

func TestBuildArray1(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}
	for _, test := range tests {
		if output := buildArray1(test.input); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}
func TestBuildArray2(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}
	for _, test := range tests {
		if output := buildArray2(test.input); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

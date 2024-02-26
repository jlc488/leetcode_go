package leetcode1389

import (
	"slices"
	"testing"
)

func TestCreateTargetArray(t *testing.T) {
	var tests = []struct {
		nums     []int
		index    []int
		expected []int
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}, []int{0, 4, 1, 3, 2}},
		{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}, []int{0, 1, 2, 3, 4}},
		{[]int{1}, []int{0}, []int{1}},
	}
	for _, test := range tests {
		if output := createTargetArray(test.nums, test.index); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.nums, test.expected, output)
		}
	}
}

package leetcode2974

import (
	"slices"
	"testing"
)

func TestNumberGame(t *testing.T) {

	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{5, 4, 2, 3}, []int{3, 2, 5, 4}},
		{[]int{2, 5}, []int{5, 2}},
	}

	for _, test := range tests {
		got := numberGame(test.nums)
		if !slices.Equal(got, test.expected) {
			t.Errorf("got: %v, expected: %v", got, test.expected)
		}
	}
}

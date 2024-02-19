package leetcode1720

import (
	"slices"
	"testing"
)

func TestDecode(t *testing.T) {
	var tests = []struct {
		encoded  []int
		first    int
		expected []int
	}{
		{[]int{1, 2, 3}, 1, []int{1, 0, 2, 1}},
		{[]int{6, 2, 7, 3}, 4, []int{4, 2, 0, 7, 4}},
	}

	for _, test := range tests {
		if output := decode(test.encoded, test.first); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.encoded, test.expected, output)
		}
	}
}

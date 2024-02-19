package leetcode1512

import "testing"

func TestNumIdenticalPairs1(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
	}

	for _, test := range tests {
		if output := numIdenticalPairs1(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}
func TestNumIdenticalPairs2(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
	}

	for _, test := range tests {
		if output := numIdenticalPairs2(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

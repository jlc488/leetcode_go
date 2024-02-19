package leetcode1637

import "testing"

func TestMaxWidthOfVerticalArea(t *testing.T) {
	var tests = []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}, 1},
		{[][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}, 3},
	}

	for _, test := range tests {
		if output := maxWidthOfVerticalArea(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

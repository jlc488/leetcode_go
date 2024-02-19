package leetcode2011

import "testing"

func TestFinalValueAfterOperation(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{[]string{"--X", "X++", "X++"}, 1},
		{[]string{"++X", "++X", "X++"}, 3},
		{[]string{"X++", "++X", "--X", "X--"}, 0},
	}
	for _, test := range tests {
		if output := finalValueAfterOperations(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

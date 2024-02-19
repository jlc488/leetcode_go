package leetcode2114

import "testing"

func TestMostWordsFound(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6},
		{[]string{"please wait", "continue to fight", "continue to win"}, 3},
	}
	for _, test := range tests {
		if output := mostWordsFound(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

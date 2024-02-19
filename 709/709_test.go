package leetcode709

import "testing"

func TestToLowerCase(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
	}

	for _, test := range tests {
		if output := toLowerCase(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

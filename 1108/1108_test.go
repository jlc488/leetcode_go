package leetcode1108

import "testing"

func TestDefangIPaddr(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
	}

	for _, test := range tests {
		if output := defangIPaddr(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

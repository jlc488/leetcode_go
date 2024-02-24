package leetcode1816

import "testing"

func TestTruncateSentence(t *testing.T) {
	var tests = []struct {
		s        string
		k        int
		expected string
	}{
		{"Hello how are you Contestant", 4, "Hello how are you"},
		{"What is the solution to this problem", 4, "What is the solution"},
		{"chopper is not a tanuki", 5, "chopper is not a tanuki"},
	}
	for _, test := range tests {
		if output := truncateSentence(test.s, test.k); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.s, test.expected, output)
		}
	}
}

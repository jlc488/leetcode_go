package leetcode1528

import "testing"

func TestRestoreString(t *testing.T) {
	var tests = []struct {
		s        string
		indices  []int
		expected string
	}{
		{"codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
		{"abc", []int{0, 1, 2}, "abc"},
	}
	for _, test := range tests {
		if output := restoreString1(test.s, test.indices); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.s, test.expected, output)
		}
	}
	for _, test := range tests {
		if output := restoreString2(test.s, test.indices); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.s, test.expected, output)
		}
	}
}

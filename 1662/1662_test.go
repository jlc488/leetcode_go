package leetcode1662

import "testing"

func TestArrayStringsAreEqual(t *testing.T) {
	var tests = []struct {
		word1    []string
		word2    []string
		expected bool
	}{
		{[]string{"ab", "c"}, []string{"a", "bc"}, true},
		{[]string{"a", "cb"}, []string{"ab", "c"}, false},
		{[]string{"abc", "d", "defg"}, []string{"abcddefg"}, true},
	}

	for _, test := range tests {
		if output := arrayStringsAreEqual(test.word1, test.word2); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.word1, test.expected, output)
		}
	}
}

package leetcode2942

import (
	"slices"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	var tests = []struct {
		input    []string
		target   byte
		expected []int
	}{
		{[]string{"leet", "code"}, byte('e'), []int{0, 1}},
		{[]string{"abc", "bcd", "aaaa", "cbc"}, byte('a'), []int{0, 2}},
		{[]string{"abc", "bcd", "aaaa", "cbc"}, byte('z'), []int{}},
	}
	for _, test := range tests {
		if output := findWordsContaining(test.input, test.target); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

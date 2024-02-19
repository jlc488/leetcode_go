package leetcode771

import "testing"

func TestNumJewelsInStones(t *testing.T) {

	var tests = []struct {
		jewels   string
		stones   string
		expected int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, test := range tests {
		if output := numJewelsInStones(test.jewels, test.stones); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, %v received: %v", test.jewels, test.stones, test.expected, output)
		}
	}
}

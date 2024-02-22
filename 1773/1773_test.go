package leetcode1773

import "testing"

func TestCountMatches(t *testing.T) {
	var tests = []struct {
		items     [][]string
		ruleKey   string
		ruleValue string
		expected  int
	}{
		{[][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver", 1},
		{[][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}, "type", "phone", 2},
	}
	for _, test := range tests {
		if output := countMatches(test.items, test.ruleKey, test.ruleValue); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.items, test.expected, output)
		}
	}
}

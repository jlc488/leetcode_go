package leetcode2798

import "testing"

func TestNumberOfEmployeeWhoMetTarget(t *testing.T) {
	var tests = []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4}, 2, 3},
		{[]int{5, 1, 4, 2, 2}, 6, 0},
	}
	for _, test := range tests {
		if output := numberOfEmployeesWhoMetTarget(test.input, test.target); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

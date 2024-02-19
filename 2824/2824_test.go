package leetcode2824

import "testing"

func TestCountPairs(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output int
	}{
		{[]int{-1, 1, 2, 3, 1}, 2, 3},
		{[]int{-6, 2, 5, -2, -7, -1, 3}, -2, 10},
	}
	for _, test := range tests {
		if output := countPairs(test.input, test.target); output != test.output {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.output, output)
		}
	}
}

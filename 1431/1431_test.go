package leetcode1431

import (
	"slices"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {

	var tests = []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, test := range tests {
		if output := kidsWithCandies(test.candies, test.extraCandies); !slices.Equal(output, test.expected) {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.candies, test.extraCandies, output)
		}
	}
}

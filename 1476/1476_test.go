package leeetcode1476

import "testing"

func TestSubrectangleQueries_GetValue(t *testing.T) {

	var tests = []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}, 1},
	}

	for _, test := range tests {
		sr := Constructor(test.input)
		if output := sr.GetValue(0, 2); output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestSubrectangleQueries_Update(t *testing.T) {

	var tests = []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}, 1},
	}

	for _, test := range tests {
		sr := Constructor(test.input)
		sr.UpdateSubrectangle(0, 0, 3, 2, 5)
		if output := sr.GetValue(0, 2); output != 5 {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, 10, output)
		}
	}
}

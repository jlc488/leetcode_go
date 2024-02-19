package leetcode1

import (
	"slices"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	actual := twoSum1([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	if slices.Equal(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
func TestTwoSum2(t *testing.T) {
	actual := twoSum2([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	if slices.Equal(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

}

package leetcode412

import (
	"slices"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	actual := fizzBuzz(15)
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

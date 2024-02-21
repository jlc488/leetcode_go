package leetcode1221

import "testing"

func TestBalancedStringSplit(t *testing.T) {
	s := "RLRRLLRLRL"
	expected := 4
	if balancedStringSplit(s) != expected {
		t.Errorf("expected: %d, got: %d", expected, balancedStringSplit(s))
	}
}

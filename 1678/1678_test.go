package leetcode1678

import "testing"

func TestInterpret(t *testing.T) {
	command := "G()(al)"
	expected := "Goal"
	if interpret1(command) != expected {
		t.Errorf("expected: %s, got: %s", expected, interpret1(command))
	}
}

func TestInterpret2(t *testing.T) {
	command := "G()(al)"
	expected := "Goal"
	if interpret2(command) != expected {
		t.Errorf("expected: %s, got: %s", expected, interpret2(command))
	}
}

package leetcode1678

import "testing"

func TestInterpret(t *testing.T) {
	command := "G()(al)"
	expected := "Goal"
	if interpret(command) != expected {
		t.Errorf("expected: %s, got: %s", expected, interpret(command))
	}
}

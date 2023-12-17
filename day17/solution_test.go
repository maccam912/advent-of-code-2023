package day17

import (
	"testing"
)

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 102 {
		t.Errorf("Expected 102, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 94 {
		t.Errorf("Expected 94, got %d", answer)
	}

	answer = B("example_input_2.txt")
	if answer != 71 {
		t.Errorf("Expected 71, got %d", answer)
	}
}

package template

import (
	"testing"
)

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 0 {
		t.Errorf("Expected 0, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 0 {
		t.Errorf("Expected 0, got %d", answer)
	}
}

package day18

import (
	"testing"
)

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 62 {
		t.Errorf("Expected 62, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 952408144115 {
		t.Errorf("Expected 952408144115, got %d", answer)
	}
}

package day15

import (
	"testing"
)

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 1320 {
		t.Errorf("Expected 1320, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 145 {
		t.Errorf("Expected 145, got %d", answer)
	}
}

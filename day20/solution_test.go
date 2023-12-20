package day20

import (
	"testing"
)

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 32000000 {
		t.Errorf("Expected 32000000, got %d", answer)
	}

	answer = A("example_input_2.txt")
	if answer != 11687500 {
		t.Errorf("Expected 11687500, got %d", answer)
	}
}

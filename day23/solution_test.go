package day23

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	grid := parseInput("example_input.txt")
	if len(grid.grid) != 213 {
		t.Errorf("Expected 213, got %d", len(grid.grid))
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 94 {
		t.Errorf("Expected 94, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 0 {
		t.Errorf("Expected 0, got %d", answer)
	}
}

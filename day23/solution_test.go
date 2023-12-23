package day23

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	grid := parseInput("example_input.txt", 'A')
	if len(grid.grid) != 55 {
		t.Errorf("Expected 55, got %d", len(grid.grid))
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
	if answer != 154 {
		t.Errorf("Expected 154, got %d", answer)
	}
}

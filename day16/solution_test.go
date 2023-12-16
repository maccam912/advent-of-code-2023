package day16

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	grid := parseInput("example_input.txt")
	if grid[Coord{0, 0}].symbol != '.' {
		t.Errorf("Expected '.', got %v", grid[Coord{0, 0}].symbol)
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 46 {
		t.Errorf("Expected 46, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 51 {
		t.Errorf("Expected 51, got %d", answer)
	}
}

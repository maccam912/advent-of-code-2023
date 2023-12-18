package day17

import (
	"testing"

	"github.com/maccam912/advent-of-code-2023/util"
)

func TestParseInput(t *testing.T) {
	grid := parseInput("example_input.txt")
	if grid.cells[util.Coord{Row: 0, Col: 0}] != 2 {
		t.Errorf("Expected 2, got %d", grid.cells[util.Coord{Row: 0, Col: 0}])
	}
}

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

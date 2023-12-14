package day14

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	platform := parseInput("example_input.txt")

	if platform.width != 10 {
		t.Errorf("Expected width 10, got %d", platform.width)
	}

	if platform.height != 10 {
		t.Errorf("Expected height 10, got %d", platform.height)
	}

	if len(platform.objects) != 35 {
		t.Errorf("Expected 35 objects, got %d", len(platform.objects))
	}

	if platform.objects[Coord{row: 0, col: 0}] != 'O' {
		t.Errorf("Expected # at 0,0, got %c", platform.objects[Coord{row: 0, col: 0}])
	}

	if platform.objects[Coord{row: 0, col: 5}] != '#' {
		t.Errorf("Expected . at 0,5, got %c", platform.objects[Coord{row: 9, col: 9}])
	}
}

func TestTiltNorth(t *testing.T) {
	platform := parseInput("example_input.txt")
	platform.Tilt('N')

	if len(platform.objects) != 35 {
		t.Errorf("Expected 35 objects, got %d", len(platform.objects))
	}

	if platform.objects[Coord{row: 0, col: 1}] != 'O' {
		t.Errorf("Expected # at 0,1, got %c", platform.objects[Coord{row: 0, col: 1}])
	}

	if platform.objects[Coord{row: 0, col: 2}] != 'O' {
		t.Errorf("Expected # at 0,2, got %c", platform.objects[Coord{row: 0, col: 2}])
	}
}

func TestLoad(t *testing.T) {
	platform := parseInput("example_input.txt")
	platform.Tilt('N')

	if platform.CalculateLoad() != 136 {
		t.Errorf("Expected 136, got %d", platform.CalculateLoad())
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 136 {
		t.Errorf("Expected 136, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt", 1000)
	if answer != 64 {
		t.Errorf("Expected 64, got %d", answer)
	}
}

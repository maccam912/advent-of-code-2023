package day13

import (
	"testing"
)

func TestMirrorVertical(t *testing.T) {
	field := Field{
		width:  3,
		height: 3,
		locs: map[Coord]bool{
			Coord{row: 0, col: 0}: true,
			Coord{row: 0, col: 1}: true,
			Coord{row: 0, col: 2}: true,
			Coord{row: 1, col: 0}: true,
			Coord{row: 1, col: 1}: true,
			Coord{row: 1, col: 2}: true,
			Coord{row: 2, col: 0}: true,
			Coord{row: 2, col: 1}: true,
			Coord{row: 2, col: 2}: true,
		},
	}

	mirrored := field.CheckVerticalSymmetry(1, 'A')

	if !mirrored {
		t.Errorf("Expected true, got false")
	}
	mirrored = field.CheckVerticalSymmetry(2, 'A')

	if !mirrored {
		t.Errorf("Expected true, got false")
	}
}

func TestMirrorNotVertical(t *testing.T) {
	field := Field{
		width:  3,
		height: 3,
		locs: map[Coord]bool{
			Coord{row: 0, col: 0}: true,
			Coord{row: 0, col: 2}: true,
			Coord{row: 1, col: 0}: true,
			Coord{row: 1, col: 1}: true,
			Coord{row: 1, col: 2}: true,
			Coord{row: 2, col: 0}: true,
			Coord{row: 2, col: 1}: true,
			Coord{row: 2, col: 2}: true,
		},
	}

	mirrored := field.CheckVerticalSymmetry(1, 'A')

	if mirrored {
		t.Errorf("Expected false, got true")
	}
	mirrored = field.CheckVerticalSymmetry(2, 'A')

	if mirrored {
		t.Errorf("Expected false, got true")
	}
}

func TestMirrorHorizontal(t *testing.T) {
	field := Field{
		width:  3,
		height: 3,
		locs: map[Coord]bool{
			Coord{row: 0, col: 0}: true,
			Coord{row: 0, col: 1}: true,
			Coord{row: 0, col: 2}: true,
			Coord{row: 1, col: 0}: true,
			Coord{row: 1, col: 1}: true,
			Coord{row: 1, col: 2}: true,
			Coord{row: 2, col: 0}: true,
			Coord{row: 2, col: 1}: true,
			Coord{row: 2, col: 2}: true,
		},
	}

	mirrored := field.CheckHorizontalSymmetry(1, 'A')

	if !mirrored {
		t.Errorf("Expected true, got false")
	}
	mirrored = field.CheckHorizontalSymmetry(2, 'A')

	if !mirrored {
		t.Errorf("Expected true, got false")
	}
}

func TestVertSymmetryOnField(t *testing.T) {
	fields := parseInput("example_input.txt")

	val := fields[0].FindSymmetryColumn('A')
	if val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
}

func TestHorizSymmetryOnField(t *testing.T) {
	fields := parseInput("example_input.txt")

	val := fields[1].FindSymmetryRow('A')
	if val != 4 {
		t.Errorf("Expected 5, got %d", val)
	}
}

func TestParseInput(t *testing.T) {
	fields := parseInput("example_input.txt")

	if len(fields) != 2 {
		t.Errorf("Expected 2 fields, got %d", len(fields))
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 405 {
		t.Errorf("Expected 405, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 400 {
		t.Errorf("Expected 400, got %d", answer)
	}
}

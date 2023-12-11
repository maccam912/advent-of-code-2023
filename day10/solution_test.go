package day10

import (
	"testing"
	// Import any additional packages needed for testing
)

func TestParseInput(t *testing.T) {
	g := NewGraph()
	g.ParseFile("example_input.txt")
	if len(g.edges[Coord{1, 3}]) != 2 {
		t.Errorf("Expected Coord{1, 2}, got %v", g.edges[Coord{1, 1}][0])
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 4 {
		t.Errorf("Expected 4, got %d", answer)
	}
	answer = A("example_input_2.txt")
	if answer != 8 {
		t.Errorf("Expected 8, got %d", answer)
	}
}

func TestB(t *testing.T) {
	// answer := B("example_input.txt")
	// if answer != 1 {
	// 	t.Errorf("Expected 1, got %d", answer)
	// }
	// answer = B("example_input_inside_8.txt")
	// if answer != 8 {
	// 	t.Errorf("Expected 8, got %d", answer)
	// }
	answer := B("example_input_inside_10.txt")
	if answer != 10 {
		t.Errorf("Expected 10, got %d", answer)
	}
	// answer = B("example_input_inside_4.txt")
	// if answer != 4 {
	// 	t.Errorf("Expected 4, got %d", answer)
	// }
}

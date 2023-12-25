package day25

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	graph := parseInput("input.txt")
	if len(graph.nodes) != 15 {
		t.Errorf("Expected 15 nodes, got %d", len(graph.nodes))
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 0 {
		t.Errorf("Expected 0, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 0 {
		t.Errorf("Expected 0, got %d", answer)
	}
}

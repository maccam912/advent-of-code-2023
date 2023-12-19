package day19

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	input := parseInput("example_input.txt")
	if len(input.workflows) != 11 {
		t.Errorf("Expected 11 workflows, got %d", len(input.workflows))
	}
	if len(input.parts) != 5 {
		t.Errorf("Expected 5 parts, got %d", len(input.parts))
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 19114 {
		t.Errorf("Expected 19114, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 0 {
		t.Errorf("Expected 0, got %d", answer)
	}
}

package day06

import (
	"testing"
	// Import any additional packages needed for testing
)

func TestParseInput(t *testing.T) {
	races := parseInput("example_input.txt")
	if races[0].Time != 7 {
		t.Errorf("Expected time to be 7, but was %d\n", races[0].Time)
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 288 {
		t.Errorf("Expected answer to be 288, but was %d\n", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 71503 {
		t.Errorf("Expected answer to be 71503, but was %d\n", answer)
	}
}

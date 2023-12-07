package day07

import (
	"testing"
	// Import any additional packages needed for testing
)

func TestParseInput(t *testing.T) {
	hands := parseInput("example_input.txt")
	if hands[0].cards[0] != '3' {
		t.Errorf("Expected 3, got %v", hands[0].cards[0])
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 6440 {
		t.Errorf("Expected 6440, got %v", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 5905 {
		t.Errorf("Expected 5905, got %v", answer)
	}
}

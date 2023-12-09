package day09

import (
	"testing"
	// Import any additional packages needed for testing
)

func TestParseInput(t *testing.T) {
	nums := parseInput("example_input.txt")
	if nums[0][0] != 0 {
		t.Errorf("Expected 0, got %v", nums[0][0])
	}

	if nums[1][1] != 3 {
		t.Errorf("Expected 3, got %v", nums[1][1])
	}
}

func TestDiffs(t *testing.T) {
	nums := parseInput("example_input.txt")
	diffs := diffs(nums[0])
	if diffs[0] != 3 {
		t.Errorf("Expected 3, got %v", diffs[0])
	}

	if diffs[4] != 3 {
		t.Errorf("Expected 3, got %v", diffs[1])
	}
}

func TestFindNext(t *testing.T) {
	nums := parseInput("example_input.txt")
	next := findNext(nums[0])
	if next != 18 {
		t.Errorf("Expected 18, got %v", next)
	}
	next = findNext(nums[1])
	if next != 28 {
		t.Errorf("Expected 28, got %v", next)
	}
	next = findNext(nums[2])
	if next != 68 {
		t.Errorf("Expected 68, got %v", next)
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 114 {
		t.Errorf("Expected 114, got %v", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 2 {
		t.Errorf("Expected 2, got %v", answer)
	}
}

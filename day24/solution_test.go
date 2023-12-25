package day24

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	sky := parseInput("example_input.txt")
	if len(sky.hailstones) != 5 {
		t.Errorf("Expected 5 hailstones, got %d", len(sky.hailstones))
	}
}

func TestCheckIntersections(t *testing.T) {
	sky := parseInput("example_input.txt")
	count := sky.CheckIntersections(7, 27)
	if count != 2 {
		t.Errorf("Expected 2 intersections, got %d", count)
	}
}

func TestA(t *testing.T) {
	answer := A("input.txt")
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

package day05

import (
	"testing"
	// Import any additional packages needed for testing
)

func TestParseInput(t *testing.T) {
	almanac := parseInput("example_input.txt")
	if len(almanac.Seeds) != 4 {
		t.Errorf("Expected 2 seeds, got %d", len(almanac.Seeds))
	}
	if almanac.Seeds[0] != 79 {
		t.Errorf("Expected first seed to be 79, got %d", almanac.Seeds[0])
	}

	if almanac.Seed2Soil.Ranges[0].Dst != 50 {
		t.Errorf("Expected first seed2soil range dst to be 50, got %d", almanac.Seed2Soil.Ranges[0].Dst)
	}

	if almanac.Seed2Soil.Get(79) != 81 {
		t.Errorf("Expected seed2soil.Get(79) to be 81, got %d", almanac.Seed2Soil.Get(98))
	}

	if almanac.GetLocation(79) != 82 {
		t.Errorf("Expected GetLocation(79) to be 82, got %d", almanac.GetLocation(79))
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 35 {
		t.Errorf("Expected answer to be 35, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 46 {
		t.Errorf("Expected answer to be 35, got %d", answer)
	}
}

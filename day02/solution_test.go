package day02

import (
	"testing"
	// Import any additional packages needed for testing
)

func TestParseSubet(t *testing.T) {
	expected := Subset{Red: 4, Green: 0, Blue: 3}
	actual := parseSubset("3 blue, 4 red")
	if actual != expected {
		t.Errorf("parseSubset() = %v, want %v", actual, expected)
	}
}

func TestParseInput(t *testing.T) {
	actual := parseInput("example_input.txt")
	expected := []Game{
		{ID: 1, Subsets: []Subset{{Red: 4, Green: 0, Blue: 3}, {Red: 1, Green: 2, Blue: 6}, {Red: 0, Green: 2, Blue: 0}}},
		{ID: 2, Subsets: []Subset{{Red: 0, Green: 2, Blue: 1}, {Red: 1, Green: 3, Blue: 4}, {Red: 0, Green: 1, Blue: 1}}},
		{ID: 3, Subsets: []Subset{{Red: 20, Green: 8, Blue: 6}, {Red: 4, Green: 13, Blue: 5}, {Red: 1, Green: 5, Blue: 0}}},
		{ID: 4, Subsets: []Subset{{Red: 3, Green: 1, Blue: 6}, {Red: 6, Green: 3, Blue: 0}, {Red: 14, Green: 3, Blue: 15}}},
		{ID: 5, Subsets: []Subset{{Red: 6, Green: 3, Blue: 1}, {Red: 1, Green: 2, Blue: 2}}},
	}

	for i, game := range actual {
		if game.ID != expected[i].ID {
			t.Errorf("parseInput() = %v, want %v", actual, expected)
		}
		for j, subset := range game.Subsets {
			if subset != expected[i].Subsets[j] {
				t.Errorf("parseInput() = %v, want %v", actual, expected)
			}
		}
	}
}

func TestA(t *testing.T) {
	actual := A("example_input.txt")
	expected := 8
	if actual != expected {
		t.Errorf("A() = %v, want %v", actual, expected)
	}
}

func TestB(t *testing.T) {
	actual := B("example_input.txt")
	expected := 2286
	if actual != expected {
		t.Errorf("A() = %v, want %v", actual, expected)
	}
}

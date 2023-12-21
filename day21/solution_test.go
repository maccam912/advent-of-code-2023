package day21

import (
	"testing"

	"github.com/maccam912/advent-of-code-2023/util"
)

func TestParseInput(t *testing.T) {
	garden := parseInput("example_input.txt")
	if len(garden.plots) != 81 {
		t.Errorf("Expected 81 plots, got %d", len(garden.plots))
	}
}

func TestFindDistances(t *testing.T) {
	garden := parseInput("example_input.txt")
	c := garden.Get(util.Coord{Row: 1, Col: 15})
	if !c {
		t.Errorf("Expected true, got %t", c)
	}
	c = garden.Get(util.Coord{Row: 1, Col: 16})
	if c {
		t.Errorf("Expected false, got %t", c)
	}
	c = garden.Get(util.Coord{Row: -2, Col: 15})
	if !c {
		t.Errorf("Expected true, got %t", c)
	}
	c = garden.Get(util.Coord{Row: -3, Col: 15})
	if c {
		t.Errorf("Expected false, got %t", c)
	}
	distances := garden.FindDistances(1, false)
	if len(distances) != 2 {
		t.Errorf("Expected 4 plots, got %d", len(distances))
	}
	distances = garden.FindDistances(6, false)
	if len(distances) != 16 {
		t.Errorf("Expected 16 plots, got %d", len(distances))
	}
	distances = garden.FindDistances(10, false)
	if len(distances) != 50 {
		t.Errorf("Expected 50 plots, got %d", len(distances))
	}
	distances = garden.FindDistances(50, false)
	if len(distances) != 1594 {
		t.Errorf("Expected 1594 plots, got %d", len(distances))
	}
	distances = garden.FindDistances(100, false)
	if len(distances) != 6536 {
		t.Errorf("Expected 6536 plots, got %d", len(distances))
	}
	// distances = garden.FindDistances(500, false)
	// if len(distances) != 167004 {
	// 	t.Errorf("Expected 167004 plots, got %d", len(distances))
	// }
	// distances = garden.FindDistances(1000)
	// if len(distances) != 668697 {
	// 	t.Errorf("Expected 668697 plots, got %d", len(distances))
	// }
	// distances = garden.FindDistances(5000)
	// if len(distances) != 16733044 {
	// 	t.Errorf("Expected 16733044 plots, got %d", len(distances))
	// }
}

func TestInBoundsCoord(t *testing.T) {
	garden := parseInput("example_input.txt")
	for i := -100; i < 100; i++ {
		for j := -100; j < 100; j++ {
			result := garden.InBoundsCoord(util.Coord{Row: i, Col: j})
			if result.Row < 0 || result.Row >= garden.height {
				t.Errorf("Expected row in bounds, got %d", result.Row)
			}
			if result.Col < 0 || result.Col >= garden.width {
				t.Errorf("Expected col in bounds, got %d", result.Col)
			}
		}
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 44613306015022 {
		t.Errorf("Expected 44613306015022, got %d", answer)
	}
}

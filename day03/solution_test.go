package day03

import "testing"

// Import any additional packages needed for testing

func TestParseGame(t *testing.T) {
	game, nums, symbols := parseGame("example_input.txt")

	if game[Coord{0, 0}] != '4' {
		t.Errorf("Expected ., got %c", game[Coord{0, 0}])
	}

	if game[Coord{9, 9}] != '.' {
		t.Errorf("Expected ., got %c", game[Coord{9, 9}])
	}

	if nums[Coord{0, 0}].Num != 467 {
		t.Errorf("Expected 467, got %d", nums[Coord{0, 0}])
	}
	if nums[Coord{0, 1}].Num != 467 {
		t.Errorf("Expected 467, got %d", nums[Coord{0, 0}])
	}
	if nums[Coord{0, 2}].Num != 467 {
		t.Errorf("Expected 467, got %d", nums[Coord{0, 0}])
	}
	if nums[Coord{9, 1}].Num != 664 {
		t.Errorf("Expected 664, got %d", nums[Coord{9, 1}])
	}
	if symbols[Coord{1, 3}] != '*' {
		t.Errorf("Expected *, got %c", symbols[Coord{1, 3}])
	}
	if symbols[Coord{8, 3}] != '$' {
		t.Errorf("Expected $, got %c", symbols[Coord{1, 3}])
	}
}

func TestAdjacentNumbers(t *testing.T) {
	_, nums, _ := parseGame("example_input.txt")

	adjacent := getAdjacentNumbers(nums, Coord{8, 5})
	if len(adjacent) != 2 {
		t.Errorf("%v", adjacent)
		t.Errorf("Expected 2 adjacent numbers, got %d", len(adjacent))
	}
}

func TestA(t *testing.T) {
	sum := A("example_input.txt")
	if sum != 4361 {
		t.Errorf("Expected 4361, got %d", sum)
	}

	sum = A("example_input_2.txt")
	if sum != 413 {
		t.Errorf("Expected 413, got %d", sum)
	}
}

func TestB(t *testing.T) {
	sum := B("example_input.txt")
	if sum != 467835 {
		t.Errorf("Expected 467835, got %d", sum)
	}

	sum = B("example_input_2.txt")
	if sum != 6756 {
		t.Errorf("Expected 6756, got %d", sum)
	}
}

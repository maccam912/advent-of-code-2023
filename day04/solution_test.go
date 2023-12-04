package day04

import (
	"testing"
	// Import any additional packages needed for testing
)

func TestParseCards(t *testing.T) {
	cards := parseCards("example_input.txt")
	if cards[0].ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", cards[0].ID)
	}
	for _, card := range cards {
		if len(card.WinningNumbers) != 5 {
			t.Errorf("Expected 5 winning numbers, got %d", len(card.WinningNumbers))
		}
		if len(card.YourNumbers) != 8 {
			t.Errorf("Expected 8 your numbers, got %d on card ID %d", len(card.YourNumbers), card.ID)
		}
	}
}

func TestCountingWinningNumbers(t *testing.T) {
	cards := parseCards("example_input.txt")
	if cards[0].CountWinningNumbers() != 4 {
		t.Errorf("Expected 4 winning numbers, got %d", cards[0].CountWinningNumbers())
	}
}

func TestA(t *testing.T) {
	score := A("example_input.txt")
	if score != 13 {
		t.Errorf("Expected score to be 13, got %d", score)
	}
}

func TestB(t *testing.T) {
	score := B("example_input.txt")
	if score != 30 {
		t.Errorf("Expected score to be 30, got %d", score)
	}
}

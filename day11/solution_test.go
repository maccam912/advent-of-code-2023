package day11

import (
	"log/slog"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	slog.SetDefault(l) // configures log package to print with LevelInfo

	// Call all tests.
	os.Exit(m.Run())
}

func TestExpand(t *testing.T) {
	space := parseInput("example_input.txt")
	expandSpace(&space, 1)
	if space.height != 12 {
		t.Errorf("Expected height of 12, got %d", space.height)
	}
	if space.width != 13 {
		t.Errorf("Expected width of 13, got %d", space.width)
	}
}
func TestAB(t *testing.T) {
	answer := AB("example_input.txt", 9)
	if answer != 1030 {
		t.Errorf("Expected 1030, got %d", answer)
	}

	answer = AB("example_input.txt", 99)
	if answer != 8410 {
		t.Errorf("Expected 8410, got %d", answer)
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 374 {
		t.Errorf("Expected 374, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 82000210 {
		t.Errorf("Expected 82000210, got %d", answer)
	}
}

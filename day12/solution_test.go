package day12

import (
	"os"
	"runtime/pprof"
	"testing"
)

func TestParseInput(t *testing.T) {
	rows := parseInput("example_input.txt")
	if len(rows) != 6 {
		t.Errorf("Expected 6, got %d", len(rows))
	}
	if rows[0].groups[0] != 1 {
		t.Errorf("Expected 1, got %d", rows[0].groups[0])
	}
}

func TestIsValid(t *testing.T) {
	row := Row{[]int{1, 1, 1}, []rune{'#', '.', '#', '.', '#', '.'}, NewLookup()}
	if !row.IsValid() {
		t.Errorf("Expected true, got false")
	}

	row2 := Row{[]int{2, 2}, []rune{'#', '#', '.', '.', '#', '.'}, NewLookup()}
	if row2.IsValid() {
		t.Errorf("Expected false, got true")
	}

	row3 := Row{[]int{1, 1, 1}, []rune{'#', '?', '?', '?', '?', '.'}, NewLookup()}
	if !row3.IsValid() {
		t.Errorf("Expected true, got false")
	}
}

func TestCountPossibilities(t *testing.T) {
	row := Row{[]int{1, 1, 1}, []rune{'#', '.', '#', '.', '?', '.'}, NewLookup()}
	// possibilities := row.CountPossibilities()
	possibilities := row.arrangements(string(row.row), row.groups)
	if possibilities != 1 {
		t.Errorf("Expected 1, got %d", possibilities)
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 21 {
		t.Errorf("Expected 21, got %d", answer)
	}
}

func TestB(t *testing.T) {
	f, _ := os.Create("cpu.pprof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	answer := B("example_input.txt")
	if answer != 525152 {
		t.Errorf("Expected 525152, got %d", answer)
	}
}

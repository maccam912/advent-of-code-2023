package day13

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	row int
	col int
}
type Field struct {
	width  int
	height int
	locs   map[Coord]bool
}

func CheckEqual(left, right map[Coord]bool) bool {
	if len(left) != len(right) {
		return false
	}

	for loc := range left {
		if !right[loc] {
			return false
		}
	}

	for loc := range right {
		if !left[loc] {
			return false
		}
	}

	return true
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func CheckAlmostEqual(left, right map[Coord]bool) bool {
	if Abs(len(left)-len(right)) != 1 {
		return false
	}
	diffs := 0
	for loc := range left {
		if !right[loc] {
			diffs++
		}
	}

	for loc := range right {
		if !left[loc] {
			diffs++
		}
	}

	return diffs == 1
}

func (field *Field) CheckVerticalSymmetry(column int, part rune) bool {
	// First create a new set of all locations right of that column
	// Note: column of 1 means the line just before the first column
	// col 2 mirrored over line 2 ends up in 1, col 3 mirrored over line 2 ends up in 0
	// col 2 mirrored over line 1 ends up in 0, col 3 mirrored over line 1 ends up in -1
	right := make(map[Coord]bool)
	for loc := range field.locs {
		// Ignore anything that mirrored will end up with a negative column
		if loc.col >= column {
			dist := 1 + (loc.col - column)
			if column-dist >= 0 {
				right[Coord{row: loc.row, col: column - dist}] = true
			}
		}
	}
	left := make(map[Coord]bool)
	for loc := range field.locs {
		if loc.col < column {
			dist := column - loc.col
			remainingCols := field.width - column
			if dist <= remainingCols {
				left[Coord{row: loc.row, col: loc.col}] = true
			}
		}
	}
	if part == 'A' {
		return CheckEqual(left, right)
	} else {
		return CheckAlmostEqual(left, right)
	}
}

func (field *Field) CheckHorizontalSymmetry(row int, part rune) bool {
	// First create a new set of all locations right of that column
	// Note: column of 1 means the line just before the first column
	// col 2 mirrored over line 2 ends up in 1, col 3 mirrored over line 2 ends up in 0
	// col 2 mirrored over line 1 ends up in 0, col 3 mirrored over line 1 ends up in -1
	top := make(map[Coord]bool)
	for loc := range field.locs {
		// Ignore anything that mirrored will end up with a negative column
		if loc.row >= row {
			dist := 1 + (loc.row - row)
			if row-dist >= 0 {
				top[Coord{row: row - dist, col: loc.col}] = true
			}
		}
	}
	bottom := make(map[Coord]bool)
	for loc := range field.locs {
		if loc.row < row {
			dist := row - loc.row
			remainingRows := field.height - row
			if dist <= remainingRows {
				bottom[Coord{row: loc.row, col: loc.col}] = true
			}
		}
	}
	if part == 'A' {
		return CheckEqual(top, bottom)
	} else {
		return CheckAlmostEqual(top, bottom)
	}
}

func (field *Field) FindSymmetryColumn(part rune) int {
	for col := 1; col < field.width; col++ {
		if field.CheckVerticalSymmetry(col, part) {
			return col
		}
	}
	return -1
}

func (field *Field) FindSymmetryRow(part rune) int {
	for row := 1; row < field.height; row++ {
		if field.CheckHorizontalSymmetry(row, part) {
			return row
		}
	}
	return -1
}

func parseField(input string) Field {
	lines := strings.Split(input, "\n")
	field := Field{locs: make(map[Coord]bool)}

	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				field.locs[Coord{row: row, col: col}] = true
			}
		}
	}
	field.width = len(lines[0])
	field.height = len(lines)
	return field
}

func parseInput(path string) []Field {
	// Start by reading the input file, splitting on double newlines
	// Parse each chunk as a separate field
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	chunks := strings.Split(string(contents), "\n\n")
	fields := make([]Field, len(chunks))
	for i, chunk := range chunks {
		fields[i] = parseField(chunk)
	}

	return fields

}

func A(path string) int {
	sum := 0
	fields := parseInput(path)
	for _, field := range fields {
		col := field.FindSymmetryColumn('A')
		if col != -1 {
			sum += col
		}
		row := field.FindSymmetryRow('A')
		if row != -1 {
			sum += row * 100
		}
	}
	return sum
}

func B(path string) int {
	sum := 0
	fields := parseInput(path)
	for _, field := range fields {
		col := field.FindSymmetryColumn('B')
		if col != -1 {
			sum += col
		}
		row := field.FindSymmetryRow('B')
		if row != -1 {
			sum += row * 100
		}
	}
	return sum
}

func Run() {
	partA := A("day13/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day13/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

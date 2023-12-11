package day11

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Galaxy struct {
	row int
	col int
}

type Space struct {
	width    int
	height   int
	galaxies []*Galaxy
}

func parseInput(path string) Space {
	lines, _ := util.ReadLines(path)
	retval := []*Galaxy{}
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				g := Galaxy{row, col}
				retval = append(retval, &g)
			}
		}
	}
	return Space{len(strings.TrimSpace(lines[0])), len(lines), retval}
}

func expandSpace(space *Space, amount int) {
	emptyRows := make([]bool, 0, space.height)
	emptyCols := make([]bool, 0, space.width)
	for i := 0; i < space.height; i++ {
		emptyRows = append(emptyRows, true)
	}
	for i := 0; i < space.width; i++ {
		emptyCols = append(emptyCols, true)
	}

	for _, g := range space.galaxies {
		emptyRows[g.row] = false
		emptyCols[g.col] = false
	}

	for i := space.height - 1; i >= 0; i-- {
		if emptyRows[i] {
			slog.Info(fmt.Sprintf("Row %d is empty, expanding.", i))
			space.height += amount
			for _, g := range space.galaxies {
				if g.row >= i {
					g.row += amount
				}
			}
		}
	}

	for i := space.width - 1; i >= 0; i-- {
		if emptyCols[i] {
			slog.Info(fmt.Sprintf("Col %d is empty, expanding.", i))
			space.width += amount
			for _, g := range space.galaxies {
				if g.col >= i {
					g.col += amount
				}
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func AB(path string, amount int) int {
	space := parseInput(path)
	expandSpace(&space, amount)
	total_dist := 0

	for galaxy_1 := 0; galaxy_1 < len(space.galaxies)-1; galaxy_1++ {
		for galaxy_2 := galaxy_1 + 1; galaxy_2 < len(space.galaxies); galaxy_2++ {
			a := space.galaxies[galaxy_1]
			b := space.galaxies[galaxy_2]
			dist := abs(a.row-b.row) + abs(a.col-b.col)
			total_dist += dist
		}
	}

	return total_dist
}

func A(path string) int {
	return AB(path, 1)
}

func B(path string) int {
	return AB(path, 999999)
}

func Run() {
	partA := A("day11/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day11/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

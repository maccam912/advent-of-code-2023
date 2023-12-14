package day14

import (
	"fmt"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Coord struct {
	row int
	col int
}

type Platform struct {
	width   int
	height  int
	objects map[Coord]rune
}

func (platform *Platform) Tilt(dir rune) {
	somethingMoved := true

	for somethingMoved {
		somethingMoved = false
		for coord, object := range platform.objects {
			if object == 'O' {
				// It is able to move north if there is a space above it

				if dir == 'N' && coord.row > 0 {
					if _, exists := platform.objects[Coord{row: coord.row - 1, col: coord.col}]; !exists {
						platform.objects[Coord{row: coord.row - 1, col: coord.col}] = 'O'
						delete(platform.objects, coord)
						somethingMoved = true
					}
				}
				if dir == 'S' && coord.row < platform.height-1 {
					if _, exists := platform.objects[Coord{row: coord.row + 1, col: coord.col}]; !exists {
						platform.objects[Coord{row: coord.row + 1, col: coord.col}] = 'O'
						delete(platform.objects, coord)
						somethingMoved = true
					}
				}
				if dir == 'W' && coord.col > 0 {
					if _, exists := platform.objects[Coord{row: coord.row, col: coord.col - 1}]; !exists {
						platform.objects[Coord{row: coord.row, col: coord.col - 1}] = 'O'
						delete(platform.objects, coord)
						somethingMoved = true
					}
				}
				if dir == 'E' && coord.col < platform.width-1 {
					if _, exists := platform.objects[Coord{row: coord.row, col: coord.col + 1}]; !exists {
						platform.objects[Coord{row: coord.row, col: coord.col + 1}] = 'O'
						delete(platform.objects, coord)
						somethingMoved = true
					}
				}
			}
		}
	}
}

func (platform *Platform) CalculateLoad() int {
	load := 0
	for coord, object := range platform.objects {
		if object == 'O' {
			load += platform.height - coord.row
		}
	}
	return load
}

func parseInput(path string) Platform {
	lines, _ := util.ReadLines(path)

	objects := make(map[Coord]rune)

	for row, line := range lines {
		for col, char := range line {
			if char != '.' {
				objects[Coord{row: row, col: col}] = char
			}
		}
	}

	return Platform{
		width:   len(lines[0]),
		height:  len(lines),
		objects: objects,
	}
}

func A(path string) int {
	platform := parseInput(path)
	platform.Tilt('N')
	return platform.CalculateLoad()
}

func B(path string, itercount int) int {
	platform := parseInput(path)
	for i := 0; i < itercount; i++ {
		platform.Tilt('N')
		platform.Tilt('W')
		platform.Tilt('S')
		platform.Tilt('E')
		fmt.Printf("Iteration %d: %d\n", i, platform.CalculateLoad())
	}
	return platform.CalculateLoad()
}

func Run() {
	partA := A("day14/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day14/input.txt", 1000)
	fmt.Printf("Part B: %v\n", partB)
}

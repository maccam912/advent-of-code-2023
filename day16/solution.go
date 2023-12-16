package day16

import (
	"fmt"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Coord struct {
	row int
	col int
}

type Cell struct {
	N      bool
	S      bool
	E      bool
	W      bool
	symbol rune
}

type Beam struct {
	loc Coord
	dir rune
}

func parseInput(path string) map[Coord]*Cell {
	grid := make(map[Coord]*Cell)
	lines, _ := util.ReadLines(path)
	for row, line := range lines {
		for col, char := range line {
			grid[Coord{row, col}] = &Cell{false, false, false, false, char}
		}
	}
	return grid
}

func debugGrid(grid map[Coord]*Cell) {
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			cell := grid[Coord{row, col}]
			if cell.N || cell.S || cell.E || cell.W {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func solve(path string, loc Coord, dir rune) int {
	grid := parseInput(path)
	beams := make([]Beam, 0)
	beams = append(beams, Beam{loc, dir})
	for len(beams) > 0 {
		// pop beam off stack
		beam := beams[0]
		beams = beams[1:]
		// Split logic depending on which way it is going

		if beam.dir == 'E' {
			newLoc := Coord{beam.loc.row, beam.loc.col + 1}
			if cell, ok := grid[newLoc]; ok {
				// Check symbol
				if cell.symbol == '.' || cell.symbol == '-' {
					// No change of direction, check if there has been a beam going east
					if !cell.E {
						// We haven't been here before, add to stack
						cell.E = true
						beams = append(beams, Beam{newLoc, 'E'})
					}
				} else if cell.symbol == '|' {
					if !cell.N {
						cell.N = true
						beams = append(beams, Beam{newLoc, 'N'})
					}
					if !cell.S {
						cell.S = true
						beams = append(beams, Beam{newLoc, 'S'})
					}
				} else if cell.symbol == '/' {
					if !cell.N {
						cell.N = true
						beams = append(beams, Beam{newLoc, 'N'})
					}
				} else if cell.symbol == '\\' {
					if !cell.S {
						cell.S = true
						beams = append(beams, Beam{newLoc, 'S'})
					}
				}
			}
		} else if beam.dir == 'N' {
			newLoc := Coord{beam.loc.row - 1, beam.loc.col}
			if cell, ok := grid[newLoc]; ok {
				// Check symbol
				if cell.symbol == '.' || cell.symbol == '|' {
					// No change of direction, check if there has been a beam going east
					if !cell.N {
						// We haven't been here before, add to stack
						cell.N = true
						beams = append(beams, Beam{newLoc, 'N'})
					}
				} else if cell.symbol == '-' {
					if !cell.E {
						cell.E = true
						beams = append(beams, Beam{newLoc, 'E'})
					}
					if !cell.W {
						cell.W = true
						beams = append(beams, Beam{newLoc, 'W'})
					}
				} else if cell.symbol == '/' {
					if !cell.W {
						cell.W = true
						beams = append(beams, Beam{newLoc, 'E'})
					}
				} else if cell.symbol == '\\' {
					if !cell.E {
						cell.E = true
						beams = append(beams, Beam{newLoc, 'W'})
					}
				}
			}
		} else if beam.dir == 'S' {
			newLoc := Coord{beam.loc.row + 1, beam.loc.col}
			if cell, ok := grid[newLoc]; ok {
				// Check symbol
				if cell.symbol == '.' || cell.symbol == '|' {
					// No change of direction, check if there has been a beam going east
					if !cell.S {
						// We haven't been here before, add to stack
						cell.S = true
						beams = append(beams, Beam{newLoc, 'S'})
					}
				} else if cell.symbol == '-' {
					if !cell.E {
						cell.E = true
						beams = append(beams, Beam{newLoc, 'E'})
					}
					if !cell.W {
						cell.W = true
						beams = append(beams, Beam{newLoc, 'W'})
					}
				} else if cell.symbol == '/' {
					if !cell.E {
						cell.E = true
						beams = append(beams, Beam{newLoc, 'W'})
					}
				} else if cell.symbol == '\\' {
					if !cell.W {
						cell.W = true
						beams = append(beams, Beam{newLoc, 'E'})
					}
				}
			}
		} else if beam.dir == 'W' {
			newLoc := Coord{beam.loc.row, beam.loc.col - 1}
			if cell, ok := grid[newLoc]; ok {
				// Check symbol
				if cell.symbol == '.' || cell.symbol == '-' {
					// No change of direction, check if there has been a beam going east
					if !cell.W {
						// We haven't been here before, add to stack
						cell.W = true
						beams = append(beams, Beam{newLoc, 'W'})
					}
				} else if cell.symbol == '|' {
					if !cell.N {
						cell.N = true
						beams = append(beams, Beam{newLoc, 'N'})
					}
					if !cell.S {
						cell.S = true
						beams = append(beams, Beam{newLoc, 'S'})
					}
				} else if cell.symbol == '/' {
					if !cell.S {
						cell.S = true
						beams = append(beams, Beam{newLoc, 'S'})
					}
				} else if cell.symbol == '\\' {
					if !cell.N {
						cell.N = true
						beams = append(beams, Beam{newLoc, 'N'})
					}
				}
			}
		}
	}
	// Now check the grid for the number of cells that have been visited
	sum := 0
	for _, cell := range grid {
		if cell.N || cell.S || cell.E || cell.W {
			sum++
		}
	}
	// debugGrid(grid)
	return sum
}

func A(path string) int {
	return solve(path, Coord{0, -1}, 'E')
}

func B(path string) int {
	grid := parseInput(path)
	maxRow := 0
	maxCol := 0
	for coord := range grid {
		if coord.row > maxRow {
			maxRow = coord.row
		}
		if coord.col > maxCol {
			maxCol = coord.col
		}
	}

	maxEnergized := 0
	for row := 0; row <= maxRow; row++ {
		loc := Coord{row, -1}
		dir := 'E'
		score := solve(path, loc, dir)
		if score > maxEnergized {
			maxEnergized = score
		}

		loc = Coord{row, maxCol + 1}
		dir = 'W'
		score = solve(path, loc, dir)
		if score > maxEnergized {
			maxEnergized = score
		}
	}

	for col := 0; col <= maxCol; col++ {
		loc := Coord{-1, col}
		dir := 'S'
		score := solve(path, loc, dir)
		if score > maxEnergized {
			maxEnergized = score
		}

		loc = Coord{maxRow + 1, col}
		dir = 'N'
		score = solve(path, loc, dir)
		if score > maxEnergized {
			maxEnergized = score
		}
	}

	return maxEnergized
}

func Run() {
	partA := A("day16/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day16/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

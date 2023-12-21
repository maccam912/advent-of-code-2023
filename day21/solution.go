package day21

import (
	"fmt"
	"os"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Garden struct {
	plots map[util.Coord]bool
	// edges  map[util.Coord][]util.Coord
	start  util.Coord
	width  int
	height int
}

func parseInput(path string) Garden {
	garden := Garden{
		plots: make(map[util.Coord]bool),
	}

	lines, _ := util.ReadLines(path)
	for row, line := range lines {
		for col, char := range line {
			if char == '.' {
				garden.plots[util.Coord{Row: row, Col: col}] = true
			} else if char == 'S' {
				garden.start = util.Coord{Row: row, Col: col}
				garden.plots[util.Coord{Row: row, Col: col}] = true
			}
		}
	}

	// garden.edges = make(map[util.Coord][]util.Coord)
	// for plot := range garden.plots {
	// 	for _, neighbor := range plot.Neighbors() {
	// 		if _, ok := garden.plots[neighbor]; ok {
	// 			garden.edges[plot] = append(garden.edges[plot], neighbor)
	// 		}
	// 	}
	// }

	garden.width = len(lines[0])
	garden.height = len(lines)

	return garden
}

func (garden *Garden) InBoundsCoord(c util.Coord) util.Coord {
	row := c.Row % garden.height
	col := c.Col % garden.width
	if row < 0 {
		row += garden.height
	}
	if col < 0 {
		col += garden.width
	}
	return util.Coord{Row: row, Col: col}
}

func (garden *Garden) Get(c util.Coord) bool {
	return garden.plots[garden.InBoundsCoord(c)]
}

func (garden *Garden) FindDistances(steps int, recordAnswers bool) map[util.Coord]bool {
	currentPositions := make(map[util.Coord]bool)
	currentPositions[garden.start] = true
	lines := []string{}
	for i := 0; i < steps; i++ {
		newPositions := make(map[util.Coord]bool)
		for currentPosition := range currentPositions {
			neighbors := currentPosition.Neighbors()
			for _, neighbor := range neighbors {
				if garden.Get(neighbor) {
					newPositions[neighbor] = true
				}
			}
		}
		if recordAnswers {
			fmt.Println(i)
			line := fmt.Sprintf("%d, %d\n", i+1, len(newPositions))
			lines = append(lines, line)
		}
		currentPositions = newPositions
	}
	if recordAnswers {
		file, _ := os.OpenFile("day21/answers.txt", os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()
		file.WriteString("steps, plots\n")
		for _, line := range lines {
			file.WriteString(line)
		}
	}
	return currentPositions
}

func A(path string) int {
	garden := parseInput(path)
	distances := garden.FindDistances(64, false)
	return len(distances)
}

func B(path string) int {
	x := 202301
	return (14881 * 14881 * x) - (914821 * x) + 3682
}

func Run() {
	partA := A("day21/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day21/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

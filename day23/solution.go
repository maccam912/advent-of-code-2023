package day23

import (
	"fmt"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Node struct {
	coord util.Coord
	char  rune
	edges []*Node
}

type Grid struct {
	grid   map[util.Coord]*Node
	width  int
	height int
}

func parseInput(path string) Grid {
	grid := Grid{make(map[util.Coord]*Node), 0, 0}
	lines, _ := util.ReadLines(path)
	for row, line := range lines {
		for col, char := range strings.TrimSpace(line) {
			if char != '#' {
				node := Node{util.Coord{Row: row, Col: col}, char, []*Node{}}
				grid.grid[node.coord] = &node
			}
		}
	}
	grid.height = len(lines)
	grid.width = len(strings.TrimSpace(lines[0]))

	for _, node := range grid.grid {
		if node.char == '.' {
			for _, neighbor := range node.coord.Neighbors() {
				if _, ok := grid.grid[neighbor]; ok {
					node.edges = append(node.edges, grid.grid[neighbor])
				}
			}
		} else {
			// Node must be a slope, only connect to the direction it's going
			if node.char == '^' {
				coordUp := util.Coord{Row: node.coord.Row - 1, Col: node.coord.Col}
				node.edges = append(node.edges, grid.grid[coordUp])
			} else if node.char == 'v' {
				coordDown := util.Coord{Row: node.coord.Row + 1, Col: node.coord.Col}
				node.edges = append(node.edges, grid.grid[coordDown])
			} else if node.char == '<' {
				coordLeft := util.Coord{Row: node.coord.Row, Col: node.coord.Col - 1}
				node.edges = append(node.edges, grid.grid[coordLeft])
			} else if node.char == '>' {
				coordRight := util.Coord{Row: node.coord.Row, Col: node.coord.Col + 1}
				node.edges = append(node.edges, grid.grid[coordRight])
			} else {
				panic("Some character we don't recognize!")
			}
		}
	}
	return grid
}

func FindLongestPath(grid *Grid, path []*Node) int {
	endCoord := util.Coord{Row: grid.height - 1, Col: grid.width - 2}
	if path[len(path)-1].coord == endCoord {
		return len(path)
	}

	// Otherwise, we're not at exit. Keep exploring
	neighbors := path[len(path)-1].edges
	if len(neighbors) == 0 {
		return 0
	}
	pathLengths := []int{}
	for _, neighbor := range neighbors {
		// Check if we've already been here
		alreadyVisited := false
		for _, node := range path {
			if node.coord == neighbor.coord {
				alreadyVisited = true
				break
			}
		}
		if !alreadyVisited {
			clonedPath := make([]*Node, len(path)+1)
			copy(clonedPath, path)
			clonedPath[len(clonedPath)-1] = neighbor

			pathLengths = append(pathLengths, FindLongestPath(grid, clonedPath))
		}
	}
	maxDist := 0
	for _, dist := range pathLengths {
		if dist > maxDist {
			maxDist = dist
		}
	}
	return maxDist
}

func A(path string) int {
	grid := parseInput(path)
	start := util.Coord{Row: 0, Col: 1}
	return FindLongestPath(&grid, []*Node{grid.grid[start]})
}

func B(path string) int {
	return 0
}

func Run() {
	partA := A("dayXX/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("dayXX/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

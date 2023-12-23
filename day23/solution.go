package day23

import (
	"fmt"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Edge struct {
	dest   *Node
	weight int
}

type Node struct {
	coord util.Coord
	char  rune
	edges []*Edge
}

type Grid struct {
	grid   map[util.Coord]*Node
	width  int
	height int
}

func parseInput(path string, part rune) Grid {
	grid := Grid{make(map[util.Coord]*Node), 0, 0}
	lines, _ := util.ReadLines(path)
	for row, line := range lines {
		for col, char := range strings.TrimSpace(line) {
			if char != '#' {
				node := Node{util.Coord{Row: row, Col: col}, char, []*Edge{}}
				grid.grid[node.coord] = &node
			}
		}
	}
	grid.height = len(lines)
	grid.width = len(strings.TrimSpace(lines[0]))

	for _, node := range grid.grid {
		if node.char == '.' || (part == 'B' && node.char != '#') {
			for _, neighbor := range node.coord.Neighbors() {
				if _, ok := grid.grid[neighbor]; ok {
					node.edges = append(node.edges, &Edge{grid.grid[neighbor], 1})
				}
			}
		} else {
			if part == 'A' {
				// Node must be a slope, only connect to the direction it's going
				if node.char == '^' {
					coordUp := util.Coord{Row: node.coord.Row - 1, Col: node.coord.Col}
					node.edges = append(node.edges, &Edge{grid.grid[coordUp], 1})
				} else if node.char == 'v' {
					coordDown := util.Coord{Row: node.coord.Row + 1, Col: node.coord.Col}
					node.edges = append(node.edges, &Edge{grid.grid[coordDown], 1})
				} else if node.char == '<' {
					coordLeft := util.Coord{Row: node.coord.Row, Col: node.coord.Col - 1}
					node.edges = append(node.edges, &Edge{grid.grid[coordLeft], 1})
				} else if node.char == '>' {
					coordRight := util.Coord{Row: node.coord.Row, Col: node.coord.Col + 1}
					node.edges = append(node.edges, &Edge{grid.grid[coordRight], 1})
				} else {
					panic("Some character we don't recognize!")
				}
			}
		}
	}
	merged := true
	for merged {
		merged = false
		for _, node := range grid.grid {
			if MergeGraph(node) {
				// fmt.Printf("Merged %v\n", node.coord)
				// fmt.Printf("Length of grid: %v\n", len(grid.grid))
				delete(grid.grid, node.coord)
				merged = true
				break
			}
		}
	}
	n := grid.grid[util.Coord{Row: 0, Col: 1}]
	fmt.Println(n)
	return grid
}

func MergeGraph(node *Node) bool {
	// Check if this node has two neighbors. If each of those has only 2 and one of the two is this original node, remove the original
	// It is like removing a node from a doubly linked list s
	if len(node.edges) == 2 {
		neighbor1 := node.edges[0].dest
		neighbor2 := node.edges[1].dest
		var ne1 *Edge
		var ne2 *Edge
		if len(neighbor1.edges) == 2 && len(neighbor2.edges) == 2 {
			// Check if neighbor1 and neighbor2 are connected
			for _, edge := range neighbor1.edges {
				if edge.dest == node {
					// save it
					ne1 = edge
					break
				}
			}
			for _, edge := range neighbor2.edges {
				if edge.dest == node {
					// save it
					ne2 = edge
					break
				}
			}
			if ne1 == nil || ne2 == nil {
				return false
			}
			// Remove node from neighbor1 and neighbor2
			sum := ne1.weight + ne2.weight
			// fmt.Printf("New sum: %v\n", sum)
			ne1.dest = neighbor2
			ne1.weight = sum
			ne2.dest = neighbor1
			ne2.weight = sum
			return true
		}
	}
	return false
}

func FindLongestPath(grid *Grid, path []*Edge) int {
	endCoord := util.Coord{Row: grid.height - 1, Col: grid.width - 2}
	if path[len(path)-1].dest.coord == endCoord {
		sum := 0
		for _, edge := range path {
			sum += edge.weight
		}
		return sum
	}

	// Otherwise, we're not at exit. Keep exploring
	neighbors := path[len(path)-1].dest.edges
	if len(neighbors) == 0 {
		return 0
	}
	pathLengths := []int{}
	for _, neighbor := range neighbors {
		// Check if we've already been here
		alreadyVisited := false
		for _, node := range path {
			if node.dest.coord == neighbor.dest.coord {
				alreadyVisited = true
				break
			}
		}
		if !alreadyVisited {
			clonedPath := make([]*Edge, len(path)+1)
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
	grid := parseInput(path, 'A')
	start := util.Coord{Row: 0, Col: 1}
	return FindLongestPath(&grid, []*Edge{&Edge{grid.grid[start], 1}}) - 1
}

func B(path string) int {
	grid := parseInput(path, 'B')
	start := util.Coord{Row: 0, Col: 1}
	return FindLongestPath(&grid, []*Edge{&Edge{grid.grid[start], 1}}) - 1
}

func Run() {
	partA := A("day23/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day23/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

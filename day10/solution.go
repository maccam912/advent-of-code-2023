package day10

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	row int
	col int
}

type Graph struct {
	nodes      map[Coord]rune
	edges      map[Coord][]Coord
	startCoord Coord
	values     map[Coord]int
}

func NewGraph() *Graph {
	return &Graph{
		nodes:  make(map[Coord]rune),
		edges:  make(map[Coord][]Coord),
		values: make(map[Coord]int),
	}
}

func (g *Graph) AddEdge(c1, c2 Coord) {
	g.edges[c1] = append(g.edges[c1], c2)
	g.edges[c2] = append(g.edges[c2], c1)
}

func (g *Graph) RenderValues() string {
	// Find the dimensions of the grid
	maxRow, maxCol := 0, 0
	for coord := range g.edges {
		if coord.row > maxRow {
			maxRow = coord.row
		}
		if coord.col > maxCol {
			maxCol = coord.col
		}
	}

	var result string
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if value, exists := g.values[Coord{row, col}]; exists {
				result += fmt.Sprintf("%d", value%10) // Add the last digit of the value
			} else {
				result += "."
			}
		}
		result += "\n"
	}

	return result
}

func (g *Graph) IsInside(coord Coord) bool {
	// Check if a space is inside or outside the grid.
	// Do this by counting how many pipes are crossed when moving left to the edge of the grid.
	// If the number of pipes crossed is odd, the space is inside the grid. Even, outside.
	// Pretend the "center" point is somewhere upperish leftish. Tiles that count as a crossing:
	// |, F, 7

	currentLoc := coord
	crossings := 0
	for currentLoc.col >= 0 && currentLoc.col < 1000 {
		// Check if it is one of the tiles
		if _, exists := g.values[currentLoc]; exists {
			if g.nodes[currentLoc] == '|' || g.nodes[currentLoc] == 'F' || g.nodes[currentLoc] == '7' {
				crossings++
			}
		}
		if coord.col < g.startCoord.col {
			currentLoc.col--
		} else {
			currentLoc.col++
		}
	}
	return crossings%2 == 1
}

func (g *Graph) ParseFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sConnections := 0 // To count the connections to 'S'

	for row, line := range grid {
		for col, char := range line {
			g.nodes[Coord{row, col}] = char
			if char == 'S' {
				g.startCoord = Coord{row, col}
			}

			// Check for vertical connections
			if row+1 < len(grid) {
				below := grid[row+1][col]
				if (char == '|' || char == 'F' || char == '7' || char == 'S') && (below == '|' || below == 'J' || below == 'L' || below == 'S') {
					g.AddEdge(Coord{row, col}, Coord{row + 1, col})
				}
			}

			// Check for horizontal connections
			if col+1 < len(line) {
				right := line[col+1]
				if (char == '-' || char == 'F' || char == 'L' || char == 'S') && (right == '-' || right == 'J' || right == '7' || right == 'S') {
					g.AddEdge(Coord{row, col}, Coord{row, col + 1})
				}
			}

			if char == 'S' {
				if row+1 < len(grid) && (grid[row+1][col] == '|' || grid[row+1][col] == 'J' || grid[row+1][col] == 'L' || grid[row+1][col] == 'S') {
					sConnections++
				}
				if col+1 < len(line) && (line[col+1] == '-' || line[col+1] == 'J' || line[col+1] == '7' || line[col+1] == 'S') {
					sConnections++
				}
			}
		}
	}

	return nil
}

func A(path string) int {
	g := NewGraph()
	err := g.ParseFile(path)
	if err != nil {
		panic(err) // or handle the error appropriately
	}

	// Initialize the values map with the starting position
	g.values[g.startCoord] = 0

	// Queue for BFS
	var queue []Coord
	queue = append(queue, g.startCoord)

	// step := 1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentValue := g.values[current]
		for _, neighbor := range g.edges[current] {
			if existingValue, exists := g.values[neighbor]; !exists || existingValue > currentValue+1 {
				g.values[neighbor] = currentValue + 1
				queue = append(queue, neighbor)
			}
			// if (step % 100) == 0 {
			// 	fmt.Println(g.RenderValues())
			// 	time.Sleep(100 * time.Millisecond)
			// }
			// step += 1
		}
	}

	// Find the maximum value in g.values
	maxDistance := 0
	for _, v := range g.values {
		if v > maxDistance {
			maxDistance = v
		}
	}

	return maxDistance
}

func B(path string) int {
	g := NewGraph()
	err := g.ParseFile(path)
	if err != nil {
		panic(err) // or handle the error appropriately
	}

	g.values[g.startCoord] = 0

	// Queue for BFS
	var queue []Coord
	queue = append(queue, g.startCoord)

	// step := 1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentValue := g.values[current]
		for _, neighbor := range g.edges[current] {
			if existingValue, exists := g.values[neighbor]; !exists || existingValue > currentValue+1 {
				g.values[neighbor] = currentValue + 1
				queue = append(queue, neighbor)
			}
			// if (step % 100) == 0 {
			// 	fmt.Println(g.RenderValues())
			// 	time.Sleep(100 * time.Millisecond)
			// }
			// step += 1
		}
	}

	insideCount := 0
	for row := 0; row < 1000; row++ {
		for col := 0; col < 1000; col++ {
			// if coord is not a key in in g.values, it isn't part of the loop
			coord := Coord{row, col}
			if _, exists := g.values[coord]; exists {
				continue
			}

			if g.IsInside(coord) {
				// fmt.Println(coord)
				insideCount++
			}
		}
	}
	return insideCount
}

func Run() {
	partA := A("day10/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day10/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

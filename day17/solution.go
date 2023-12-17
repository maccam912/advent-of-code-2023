package day17

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Node struct {
	x, y        int
	dir         Direction
	consecutive int
}

type State struct {
	node     Node
	heatLoss int
	priority int
	index    int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	state := x.(*State)
	state.index = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	old[n-1] = nil
	state.index = -1
	*pq = old[0 : n-1]
	return state
}

func neighbors(n Node, city [][]int, part rune) []Node {
	var neighbors []Node

	directions := []Direction{North, East, South, West}
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}

	for i, dir := range directions {
		// Skip a 180 degree turn
		if dir == (n.dir+2)%4 {
			continue
		}
		newX := n.x + dx[i]
		newY := n.y + dy[i]

		// Check if the new position is within the city boundaries
		if newX < 0 || newX >= len(city[0]) || newY < 0 || newY >= len(city) {
			continue
		}

		// Check if the new move is in the same direction and doesn't exceed the limit
		if part == 'A' {
			if n.dir == dir && n.consecutive == 3 {
				continue
			}
		}

		if part == 'B' {
			if n.dir != dir && n.consecutive < 4 {
				// If it is not in the same direction, but has moved less than 4, skip
				continue
			}
			if n.dir == dir && n.consecutive == 10 {
				continue
			}
		}

		// Create a new node for the valid neighbor
		newNode := Node{
			x:           newX,
			y:           newY,
			dir:         dir,
			consecutive: 1,
		}

		// If the new move is in the same direction, increment the consecutive moves
		if n.dir == dir {
			newNode.consecutive = n.consecutive + 1
		}

		// Filter out any that get to goal in les than 4 moves
		if part == 'B' && newNode.x == len(city[0])-1 && newNode.y == len(city)-1 && newNode.consecutive < 4 {
			continue
		}
		neighbors = append(neighbors, newNode)
	}

	return neighbors
}

func heuristic(n Node, goal Node) int {
	return abs(n.x-goal.x) + abs(n.y-goal.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func debugCity(city [][]int, visited map[Node]bool) {
	// Create a new 2D slice with the same dimensions as the city
	debug := make([][]string, len(city))
	for i := range debug {
		debug[i] = make([]string, len(city[i]))
		for j := range debug[i] {
			debug[i][j] = "."
		}
	}

	// Create a local map to count the visits for each coordinate
	visitCount := make(map[Node]int)

	// Count the visits for each coordinate
	for node := range visited {
		visitCount[Node{x: node.x, y: node.y}]++
	}

	// Mark the visited nodes in the debug slice
	for node, count := range visitCount {
		debug[node.y][node.x] = strconv.Itoa(count)[0:1]
	}

	// Print the debug slice
	for _, row := range debug {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

// func getPath(goal Node, visited map[Node]bool) []Node {
// 	var path []Node
// 	for node := goal; visited[node]; node = Node{x: node.parent_x, y: node.parent_y} {
// 		path = append([]Node{node}, path...)
// 	}
// 	return path
// }

func shortestPath(city [][]int, start Node, goal Node, part rune) int {
	// Initialize the priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Initialize the map of visited nodes
	visited := make(map[Node]bool)

	// Push the start node into the priority queue
	startState := &State{node: start, heatLoss: 0, priority: heuristic(start, goal)}
	heap.Push(&pq, startState)

	// While the priority queue is not empty
	for pq.Len() > 0 {
		// Pop the node with the highest priority
		currState := heap.Pop(&pq).(*State)
		currNode := currState.node
		if currNode.x == len(city[0])-1 && currNode.y == len(city)-1 {
			// path := getPath(&currNode)
			// fmt.Println(path)
			return currState.heatLoss
		}

		// If we have already visited this node, skip it
		if visited[currNode] {
			continue
		}

		// Mark this node as visited
		visited[currNode] = true

		// For each neighbor of the current node
		for _, neighbor := range neighbors(currNode, city, part) {
			// Calculate the heat loss for the neighbor
			heatLoss := currState.heatLoss + city[neighbor.y][neighbor.x]

			// If we have not visited the neighbor yet, or if the new path to the neighbor is shorter
			if !visited[neighbor] {
				// Update the priority and heat loss of the neighbor
				neighborState := &State{node: neighbor, heatLoss: heatLoss, priority: heatLoss + heuristic(neighbor, goal)}
				heap.Push(&pq, neighborState)
			}
		}
		// debugCity(city, visited)
	}

	// If there is no path to the goal, return -1
	return -1
}

func parseInput(path string) ([][]int, Node, Node, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, Node{}, Node{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var city [][]int
	var goal Node

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := make([]int, len(line))
		for x, char := range strings.Split(line, "") {
			row[x], err = strconv.Atoi(char)
			if err != nil {
				return nil, Node{}, Node{}, err
			}
		}
		city = append(city, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, Node{}, Node{}, err
	}

	// The start is always at the top left
	start := Node{x: 0, y: 0, dir: East}

	// The goal is always at the bottom right
	goal = Node{x: len(city[0]) - 1, y: len(city) - 1}

	return city, start, goal, nil
}

func A(path string) int {
	// Parse the input file
	city, start, goal, err := parseInput(path)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return -1
	}

	// Find the shortest path from the start node to the goal node
	shortestPathHeatLoss := shortestPath(city, start, goal, 'A')
	if shortestPathHeatLoss == -1 {
		fmt.Println("No path found from start to goal")
		return -1
	}

	return shortestPathHeatLoss
}

func B(path string) int {
	// Parse the input file
	city, start, goal, err := parseInput(path)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return -1
	}

	// Find the shortest path from the start node to the goal node
	shortestPathHeatLoss := shortestPath(city, start, goal, 'B')
	if shortestPathHeatLoss == -1 {
		fmt.Println("No path found from start to goal")
		return -1
	}

	return shortestPathHeatLoss
}

func Run() {
	partA := A("day17/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day17/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

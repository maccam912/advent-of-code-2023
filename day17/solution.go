package day17

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/maccam912/advent-of-code-2023/util"
)

type State struct {
	path []util.Coord
	cost int
}

type PQ []State

func (q PQ) Len() int           { return len(q) }
func (q PQ) Less(i, j int) bool { return q[i].cost < q[j].cost }
func (q PQ) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ) Push(x any)        { *q = append(*q, x.(State)) }
func (q *PQ) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }

type Grid struct {
	cells         map[util.Coord]int
	width, height int
}

func parseInput(path string) Grid {
	lines, _ := util.ReadLines(path)
	m := map[util.Coord]int{}
	for row, line := range lines {
		for col, char := range line {
			cost, _ := strconv.Atoi(string(char))
			m[util.Coord{Row: row, Col: col}] = cost
		}
	}
	return Grid{cells: m, width: len(lines[0]), height: len(lines)}
}

func NoUTurns(path []util.Coord, node util.Coord) bool {
	if len(path) <= 1 {
		return true
	}
	return path[len(path)-2] != node
}

func ConsecutiveCount(path []util.Coord, node util.Coord) int {
	// How many nodes in the same direction have we had?
	rowDir := node.Row - path[len(path)-1].Row
	colDir := node.Col - path[len(path)-1].Col
	dir := util.Coord{Row: rowDir, Col: colDir}
	count := 1

	for i := len(path) - 2; i >= 0; i-- {
		a := path[i]
		b := path[i+1]
		rowDir := b.Row - a.Row
		colDir := b.Col - a.Col
		thisDir := util.Coord{Row: rowDir, Col: colDir}
		if thisDir != dir {
			return count
		}
		count++
	}
	return count
}

func InBounds(grid Grid, node util.Coord) bool {
	return node.Row >= 0 && node.Row < grid.height && node.Col >= 0 && node.Col < grid.width
}

func CheckAll(grid Grid, path []util.Coord, node util.Coord, lb, ub int) bool {
	status := NoUTurns(path, node) && ConsecutiveCount(path, node) < ub && InBounds(grid, node)
	// Also check that if a turn is happening, it is more than lb
	dir := util.Coord{Row: node.Row - path[len(path)-1].Row, Col: node.Col - path[len(path)-1].Col}
	if len(path) <= 1 {
		return status
	}
	prevDir := util.Coord{Row: path[len(path)-1].Row - path[len(path)-2].Row, Col: path[len(path)-1].Col - path[len(path)-2].Col}
	if dir != prevDir {
		// Turning, check lb
		return status && ConsecutiveCount(path[:len(path)-1], path[len(path)-1]) > lb
	}
	return status
}

func DebugPath(grid Grid, path []util.Coord) {
	lines := []string{}
	// First make field of grid costs
	for row := 0; row < grid.height; row++ {
		line := ""
		for col := 0; col < grid.width; col++ {
			line += strconv.Itoa(grid.cells[util.Coord{Row: row, Col: col}])
		}
		lines = append(lines, line)
	}

	// Then add path
	for i, node := range path {
		if i > 0 {
			// Show an arrow, not an X
			if node.Row > path[i-1].Row {
				lines[node.Row] = lines[node.Row][:node.Col] + "v" + lines[node.Row][node.Col+1:]
			} else if node.Row < path[i-1].Row {
				lines[node.Row] = lines[node.Row][:node.Col] + "^" + lines[node.Row][node.Col+1:]
			} else if node.Col > path[i-1].Col {
				lines[node.Row] = lines[node.Row][:node.Col] + ">" + lines[node.Row][node.Col+1:]
			} else if node.Col < path[i-1].Col {
				lines[node.Row] = lines[node.Row][:node.Col] + "<" + lines[node.Row][node.Col+1:]
			}
		} else {
			lines[node.Row] = lines[node.Row][:node.Col] + "X" + lines[node.Row][node.Col+1:]
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func GetKey(path []util.Coord, node util.Coord) string {
	consecutiveCount := ConsecutiveCount(path, node)
	return fmt.Sprintf("%v, %v", path[len(path)-consecutiveCount:], node)
}

func Solve(path string, lb, ub int) int {
	grid := parseInput(path)
	start := util.Coord{Row: 0, Col: 0}
	end := util.Coord{Row: grid.height - 1, Col: grid.width - 1}

	q := PQ{}
	heap.Init(&q)
	heap.Push(&q, State{path: []util.Coord{start}, cost: 0})

	visited := map[string]int{}
	visitedCount := 1
	steps := 0

	for q.Len() > 0 {
		steps++
		// get lowest cost state
		state := heap.Pop(&q).(State)
		if state.path[len(state.path)-1] == end {
			if ConsecutiveCount(state.path[:len(state.path)-1], end) <= lb {
				// Not enough, ignore
				// fmt.Println()
			} else {
				DebugPath(grid, state.path)
				return state.cost
			}
		}

		lastNode := state.path[len(state.path)-1]
		// distFromExit := grid.height - lastNode.Row + grid.width - lastNode.Col

		// add neighbors to queue
		// if steps%100 == 0 {
		// 	DebugPath(grid, state.path)
		// 	fmt.Printf("Cost: %d, Dist from exit: %d\n", state.cost, distFromExit)
		// }
		// Check up
		upCoord := util.Coord{Row: lastNode.Row - 1, Col: lastNode.Col}
		// Check for top row, check that prev node is not up (no 180)
		key := GetKey(state.path, upCoord)
		if visited[key] < visitedCount && CheckAll(grid, state.path, upCoord, lb, ub) {
			pathCopy := make([]util.Coord, len(state.path))
			copy(pathCopy, state.path)
			heap.Push(&q, State{
				path: append(pathCopy, upCoord),
				cost: state.cost + grid.cells[upCoord],
			})
			visited[key]++
		}

		// Check down
		downCoord := util.Coord{Row: lastNode.Row + 1, Col: lastNode.Col}
		// Check for bottom row, check that prev node is not down (no 180)
		key = GetKey(state.path, downCoord)
		if visited[key] < visitedCount && CheckAll(grid, state.path, downCoord, lb, ub) {
			pathCopy := make([]util.Coord, len(state.path))
			copy(pathCopy, state.path)
			heap.Push(&q, State{
				path: append(pathCopy, downCoord),
				cost: state.cost + grid.cells[downCoord],
			})
			visited[key]++
		}

		// Check left
		leftCoord := util.Coord{Row: lastNode.Row, Col: lastNode.Col - 1}
		// Check for left col, check that prev node is not left (no 180)
		key = GetKey(state.path, leftCoord)
		if visited[key] < visitedCount && CheckAll(grid, state.path, leftCoord, lb, ub) {
			pathCopy := make([]util.Coord, len(state.path))
			copy(pathCopy, state.path)
			heap.Push(&q, State{
				path: append(pathCopy, leftCoord),
				cost: state.cost + grid.cells[leftCoord],
			})
			visited[key]++
		}

		// Check right
		rightCoord := util.Coord{Row: lastNode.Row, Col: lastNode.Col + 1}
		// Check for right col, check that prev node is not right (no 180)
		key = GetKey(state.path, rightCoord)
		if visited[key] < visitedCount && CheckAll(grid, state.path, rightCoord, lb, ub) {
			pathCopy := make([]util.Coord, len(state.path))
			copy(pathCopy, state.path)
			heap.Push(&q, State{
				path: append(pathCopy, rightCoord),
				cost: state.cost + grid.cells[rightCoord],
			})
			visited[key]++
		}
	}
	return -1
}

func A(path string) int {
	return Solve(path, 0, 4)
}

func B(path string) int {
	return Solve(path, 3, 11)
}

func Run() {
	partA := A("day17/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day17/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

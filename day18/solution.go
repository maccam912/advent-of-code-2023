package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Inst struct {
	Dir   rune
	Steps int
	Color string
}

func parseInput(path string, part rune) []Inst {
	lines, _ := util.ReadLines(path)
	insts := []Inst{}
	for _, line := range lines {
		fields := strings.Fields(line)
		dir := rune(fields[0][0])
		steps, _ := strconv.Atoi(fields[1])
		color := fields[2]
		if part == 'A' {
			inst := Inst{dir, steps, color}
			insts = append(insts, inst)
		} else {
			// convert color hex value to integer
			colorInt64, _ := strconv.ParseInt(color[2:8], 16, 0)
			colorInt := int(colorInt64)
			dir := colorInt % 16
			dirRune := 'U'
			if dir == 0 {
				dirRune = 'R'
			} else if dir == 1 {
				dirRune = 'D'
			} else if dir == 2 {
				dirRune = 'L'
			}
			amt := colorInt / 16

			inst := Inst{dirRune, amt, ""}
			insts = append(insts, inst)
		}
	}
	return insts
}

type Cell struct {
	dug      bool
	exterior bool
	color    string
}

type Grid struct {
	cells      map[util.Coord]Cell
	upperLeft  util.Coord
	lowerRight util.Coord
}

func (grid *Grid) Print() {
	for row := grid.upperLeft.Row; row <= grid.lowerRight.Row; row++ {
		for col := grid.upperLeft.Col; col <= grid.lowerRight.Col; col++ {
			cell := grid.cells[util.Coord{Row: row, Col: col}]
			if cell.dug {
				fmt.Print("#")
			} else if cell.exterior {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (grid *Grid) FloodFillExterior() {
	grid.upperLeft = util.Coord{Row: grid.upperLeft.Row - 1, Col: grid.upperLeft.Col - 1}
	grid.lowerRight = util.Coord{Row: grid.lowerRight.Row + 1, Col: grid.lowerRight.Col + 1}
	q := []util.Coord{}
	q = append(q, grid.upperLeft)
	visited := map[util.Coord]bool{}
	step := 0

	for len(q) > 0 {
		step++
		curr := q[0]
		if step%1000000 == 0 {
			fmt.Println(step)
			fmt.Println(curr)
		}
		q = q[1:]
		if visited[curr] {
			continue
		}
		visited[curr] = true
		grid.cells[curr] = Cell{dug: false, exterior: true, color: ""}
		// for the four neighbors, see if they are dug. If not, add to queue
		above := util.Coord{Row: curr.Row - 1, Col: curr.Col}
		below := util.Coord{Row: curr.Row + 1, Col: curr.Col}
		left := util.Coord{Row: curr.Row, Col: curr.Col - 1}
		right := util.Coord{Row: curr.Row, Col: curr.Col + 1}

		_, exists := grid.cells[above]
		if above.Row >= grid.upperLeft.Row && (!exists || (!grid.cells[above].dug && !grid.cells[above].exterior)) {
			q = append(q, above)
		}
		_, exists = grid.cells[below]
		if below.Row <= grid.lowerRight.Row && (!exists || (!grid.cells[below].dug && !grid.cells[below].exterior)) {
			q = append(q, below)
		}
		_, exists = grid.cells[left]
		if left.Col >= grid.upperLeft.Col && (!exists || (!grid.cells[left].dug && !grid.cells[left].exterior)) {
			q = append(q, left)
		}
		_, exists = grid.cells[right]
		if right.Col <= grid.lowerRight.Col && (!exists || (!grid.cells[right].dug && !grid.cells[right].exterior)) {
			q = append(q, right)
		}
	}

}

func (grid *Grid) DigInterior() {
	for row := grid.upperLeft.Row; row <= grid.lowerRight.Row; row++ {
		for col := grid.upperLeft.Col; col <= grid.lowerRight.Col; col++ {
			_, exists := grid.cells[util.Coord{Row: row, Col: col}]
			if !exists {
				grid.cells[util.Coord{Row: row, Col: col}] = Cell{dug: true, exterior: false, color: ""}
			}
		}
	}
}

func (grid *Grid) CountDugSquares() int {
	dugSquares := 0
	for _, cell := range grid.cells {
		if cell.dug {
			dugSquares++
		}
	}
	return dugSquares
}

func (grid *Grid) FollowPlan(plan []Inst) {
	curr := util.Coord{Row: 0, Col: 0}
	grid.cells[curr] = Cell{dug: true, color: ""}

	for _, inst := range plan {
		if inst.Dir == 'U' {
			for i := 0; i < inst.Steps; i++ {
				curr.Row--
				grid.cells[curr] = Cell{dug: true, color: inst.Color}
			}
		} else if inst.Dir == 'D' {
			for i := 0; i < inst.Steps; i++ {
				curr.Row++
				grid.cells[curr] = Cell{dug: true, color: inst.Color}
			}
		} else if inst.Dir == 'L' {
			for i := 0; i < inst.Steps; i++ {
				curr.Col--
				grid.cells[curr] = Cell{dug: true, color: inst.Color}
			}
		} else if inst.Dir == 'R' {
			for i := 0; i < inst.Steps; i++ {
				curr.Col++
				grid.cells[curr] = Cell{dug: true, color: inst.Color}
			}
		}
		if curr.Col < grid.upperLeft.Col {
			grid.upperLeft.Col = curr.Col
		}
		if curr.Col > grid.lowerRight.Col {
			grid.lowerRight.Col = curr.Col
		}
		if curr.Row < grid.upperLeft.Row {
			grid.upperLeft.Row = curr.Row
		}
		if curr.Row > grid.lowerRight.Row {
			grid.lowerRight.Row = curr.Row
		}
	}
}

func (grid *Grid) GetPoints(insts []Inst) []util.Coord {
	curr := util.Coord{Row: 0, Col: 0}
	points := []util.Coord{curr}
	for _, inst := range insts {
		if inst.Dir == 'U' {
			curr.Row -= inst.Steps
		} else if inst.Dir == 'D' {
			curr.Row += inst.Steps
		} else if inst.Dir == 'L' {
			curr.Col -= inst.Steps
		} else if inst.Dir == 'R' {
			curr.Col += inst.Steps
		}
		points = append(points, curr)
	}
	return points
}

func (grid *Grid) GetArea(insts []Inst) int {
	curr := util.Coord{Row: 0, Col: 0}
	area := 0
	extra := 0
	for _, inst := range insts {
		if inst.Dir == 'U' {
			curr.Row -= inst.Steps
			extra += inst.Steps
		} else if inst.Dir == 'D' {
			curr.Row += inst.Steps
		} else if inst.Dir == 'R' {
			curr.Col += inst.Steps
			// Add to area
			area += curr.Row * inst.Steps
			extra += inst.Steps
		} else if inst.Dir == 'L' {
			curr.Col -= inst.Steps
			// Subtract from area
			area -= curr.Row * inst.Steps
		}
	}
	return abs(area) + extra + 1
}

func A(path string) int {
	instructions := parseInput(path, 'A')
	grid := Grid{cells: map[util.Coord]Cell{}, upperLeft: util.Coord{Row: 0, Col: 0}, lowerRight: util.Coord{Row: 0, Col: 0}}
	// grid.FollowPlan(instructions)
	// grid.FloodFillExterior()
	// grid.DigInterior()
	// grid.Print()
	// return grid.CountDugSquares()
	return grid.GetArea(instructions)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func B(path string) int {
	instructions := parseInput(path, 'B')
	grid := Grid{cells: map[util.Coord]Cell{}, upperLeft: util.Coord{Row: 0, Col: 0}, lowerRight: util.Coord{Row: 0, Col: 0}}
	return grid.GetArea(instructions)
}

func Run() {
	partA := A("day18/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day18/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

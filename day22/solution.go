package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Coord3 struct {
	x int
	y int
	z int
}

type Brick struct {
	a Coord3
	b Coord3
}

func (brick *Brick) GetBottomRect() (util.Coord, util.Coord) {
	return util.Coord{Col: brick.a.x, Row: brick.a.y}, util.Coord{Col: brick.b.x, Row: brick.b.y}
}

func (brick *Brick) MoveDown() {
	brick.a.z--
	brick.b.z--
}

func (brick *Brick) IsSupported(bricks []*Brick) bool {
	bottom_z := min(brick.a.z, brick.b.z)
	if bottom_z == 1 {
		return true
	}

	thisBottomRectA, thisBottomRectB := brick.GetBottomRect()
	for _, other := range bricks {
		if other.a.z == bottom_z-1 || other.b.z == bottom_z-1 {
			otherBottomRectA, otherBottomRectB := other.GetBottomRect()
			// If this north greater than other south, or this south less than other north
			// then check if this west less than other east, or this east greater than other west
			thisNorth := max(thisBottomRectA.Row, thisBottomRectB.Row)
			thisSouth := min(thisBottomRectA.Row, thisBottomRectB.Row)
			thisEast := max(thisBottomRectA.Col, thisBottomRectB.Col)
			thisWest := min(thisBottomRectA.Col, thisBottomRectB.Col)
			otherNorth := max(otherBottomRectA.Row, otherBottomRectB.Row)
			otherSouth := min(otherBottomRectA.Row, otherBottomRectB.Row)
			otherEast := max(otherBottomRectA.Col, otherBottomRectB.Col)
			otherWest := min(otherBottomRectA.Col, otherBottomRectB.Col)
			if thisNorth >= otherSouth && thisSouth <= otherNorth && thisWest <= otherEast && thisEast >= otherWest {
				return true
			}
		}
	}
	return false
}

func parseInput(path string) []*Brick {
	bricks := []*Brick{}
	// Each line is like 1,1,8~1,1,9, with left of ~ being a, right being b
	lines, _ := util.ReadLines(path)
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), "~")
		_a := parts[0]
		_b := parts[1]
		a_coords := strings.Split(_a, ",")
		b_coords := strings.Split(_b, ",")
		a_x, _ := strconv.Atoi(a_coords[0])
		a_y, _ := strconv.Atoi(a_coords[1])
		a_z, _ := strconv.Atoi(a_coords[2])
		a := Coord3{x: a_x, y: a_y, z: a_z}
		b_x, _ := strconv.Atoi(b_coords[0])
		b_y, _ := strconv.Atoi(b_coords[1])
		b_z, _ := strconv.Atoi(b_coords[2])
		b := Coord3{x: b_x, y: b_y, z: b_z}
		brick := Brick{a: a, b: b}
		bricks = append(bricks, &brick)
	}
	return bricks
}

func SettleBricks(bricks []*Brick) int {
	moved := true
	fell := map[int]bool{}
	for moved {
		moved = false
		for i, brick := range bricks {
			if !brick.IsSupported(bricks) {
				brick.MoveDown()
				fell[i] = true
				moved = true
			}
		}
	}
	return len(fell)
}

func A(path string) int {
	bricks := parseInput(path)
	SettleBricks(bricks)
	count := 0
	for i := range bricks {
		brickRemoved := make([]*Brick, 0)
		brickRemoved = append(brickRemoved, bricks[:i]...)
		brickRemoved = append(brickRemoved, bricks[i+1:]...)

		safeToRemove := true
		for _, other := range brickRemoved {
			if !other.IsSupported(brickRemoved) {
				safeToRemove = false
				break
			}
		}
		if safeToRemove {
			count++
		}
	}
	return count
}

func B(path string) int {
	bricks := parseInput(path)
	SettleBricks(bricks)
	count := 0
	for i := range bricks {
		brickRemoved := make([]*Brick, 0)
		for j, other := range bricks {
			if i != j {
				newOther := Brick{a: other.a, b: other.b}
				brickRemoved = append(brickRemoved, &newOther)
			}
		}
		count += SettleBricks(brickRemoved)
	}
	return count
}

func Run() {
	partA := A("day22/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day22/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

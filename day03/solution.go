package day03

import (
	"fmt"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
	"github.com/samber/lo"
)

type Coord struct {
	Row int
	Col int
}

type PartNumber struct {
	Id  int
	Num int
}

func parseGame(path string) (map[Coord]rune, map[Coord]PartNumber, map[Coord]rune) {
	lines, _ := util.ReadLines(path)
	nextPartNum := 0

	maxRow := len(lines)
	maxCol := len(strings.TrimSpace(lines[0]))

	game := make(map[Coord]rune)
	nums := make(map[Coord]PartNumber)
	symbols := make(map[Coord]rune)

	for row, line := range lines {
		for col, char := range strings.TrimSpace(line) {
			game[Coord{row, col}] = char
		}
	}

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			// Check if Coord{row, col} is a digit
			if game[Coord{row, col}] >= '0' && game[Coord{row, col}] <= '9' {
				// Check if Coord{row, col} already a key in nums
				if _, ok := nums[Coord{row, col}]; !ok {
					num := 0
					cols := []int{col}
					// If not, extract number there
					for innerCol := col; innerCol <= maxCol; innerCol++ {
						if game[Coord{row, innerCol}] >= '0' && game[Coord{row, innerCol}] <= '9' {
							char := game[Coord{row, innerCol}]
							dig := int(char - '0')
							num *= 10
							num += dig
							cols = append(cols, innerCol)
						} else {
							// Not a digit, we have num and coords
							for _, c := range cols {
								partNumber := PartNumber{nextPartNum, num}
								nums[Coord{row, c}] = partNumber
							}
							nextPartNum += 1
							break
						}
					}
				}
			}
		}
	}

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			// Check if Coord{row, col} is a symbol (is not a ., or a digit)
			if game[Coord{row, col}] != '.' && (game[Coord{row, col}] < '0' || game[Coord{row, col}] > '9') {
				s := game[Coord{row, col}]
				symbols[Coord{row, col}] = s
			}
		}
	}

	return game, nums, symbols
}

func getAdjacentNumbers(nums map[Coord]PartNumber, coord Coord) []PartNumber {
	adjacent := []PartNumber{}

	// Check row above, col left to row below, col right. Eight cells around current Coord
	for row := coord.Row - 1; row <= coord.Row+1; row++ {
		for col := coord.Col - 1; col <= coord.Col+1; col++ {
			// If a key exists for row, col in nums, add it to adjacent list
			if num, ok := nums[Coord{row, col}]; ok {
				adjacent = append(adjacent, num)
			}
		}
	}

	// Deduplicate list by converting to set, then back
	adjacentSet := make(map[PartNumber]bool)
	for _, num := range adjacent {
		adjacentSet[num] = true
	}
	retval := []PartNumber{}
	for num := range adjacentSet {
		retval = append(retval, num)
	}
	return retval
}

func A(path string) int {
	// Find sum of numbers adjacent to a symbol
	_, nums, symbols := parseGame(path)
	adjacentSet := make(map[PartNumber]bool)
	for coord := range symbols {
		adjacent := getAdjacentNumbers(nums, coord)
		for _, num := range adjacent {
			adjacentSet[num] = true
		}
	}

	sum := 0
	for num := range adjacentSet {
		sum += num.Num
	}

	return sum
}

func B(path string) int {
	game, nums, symbols := parseGame(path)
	gearRatios := []int{}
	for coord := range symbols {
		s := game[coord]
		if s == '*' {
			adjacent := getAdjacentNumbers(nums, coord)
			if len(adjacent) == 2 {
				gearRatio := adjacent[0].Num * adjacent[1].Num
				gearRatios = append(gearRatios, gearRatio)
			}
		}
	}
	return lo.Sum(gearRatios)
}

// Run runs the day 1 challenge.
func Run() {
	partA := A("day03/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day03/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

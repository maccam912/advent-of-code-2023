package day01

import (
	"fmt"

	"github.com/maccam912/advent-of-code-2023/util"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

// getDigit gets the digit at the specified index of the line.
//
// If the digit is spelled out, it returns the digit. Otherwise, it returns
// false.
func getDigit(line string, index int) (int, bool) {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, word := range words {
		if len(line) >= index+len(word) && line[index:index+len(word)] == word {
			return i + 1, true
		}
	}

	if line[index] >= '0' && line[index] <= '9' {
		return int(line[index] - '0'), true
	}

	return 0, false
}

// getPartBNumFromLine finds the first two digits in the given string and returns them.
func getPartBNumFromLine(line string) int {
	var firstInt, lastInt int
	for i := 0; i < len(line); i++ {
		digit, success := getDigit(line, i)
		if success {
			firstInt = digit
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		digit, success := getDigit(line, i)
		if success {
			lastInt = digit
			break
		}
	}

	return 10*firstInt + lastInt
}

// getNumFromLine finds the first and last digits in the given string and returns them.
func getNumFromLine(line string) int {
	var firstInt, lastInt int
	// find first digit
	for _, char := range line {
		if char >= '0' && char <= '9' {
			// convert char to int
			firstInt = int(char - '0')
			break
		}
	}
	// find last digit
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			// convert char to int
			lastInt = int(line[i] - '0')
			break
		}
	}

	return 10*firstInt + lastInt
}

// A solves part A of the day 1 challenge.
func A(path string) int {
	lines, _ := util.ReadLines(path)
	ints := lop.Map(lines, func(line string, _ int) int {
		return getNumFromLine(line)
	})

	return lo.Sum(ints)
}

// B solves part B of the day 1 challenge.
func B(path string) int {
	lines, _ := util.ReadLines(path)
	ints := lop.Map(lines, func(line string, _ int) int {
		return getPartBNumFromLine(line)
	})

	return lo.Sum(ints)
}

// Run runs the day 1 challenge.
func Run() {
	partA := A("day01/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day01/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
	"github.com/samber/lo"
)

type Subset struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID      int
	Subsets []Subset
}

func parseSubset(subset string) Subset {
	// A subset is formatted like this:
	// 3 blue, 4 red
	// <Number> <Color>, <Number> <Color>
	// Note: some colors may not be present

	colors := strings.Split(subset, ",")
	retval := Subset{Red: 0, Green: 0, Blue: 0}

	for _, color := range colors {
		trimmedColor := strings.TrimSpace(color)
		parts := strings.Split(trimmedColor, " ")
		number := parts[0]
		color := parts[1]
		switch color {
		case "red":
			retval.Red, _ = strconv.Atoi(number)
		case "green":
			retval.Green, _ = strconv.Atoi(number)
		case "blue":
			retval.Blue, _ = strconv.Atoi(number)
		}
	}
	return retval
}

func parseGame(line string) Game {
	// A game is formatted like this:
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// Game <ID>: <Subset 1>; <Subset 2>; <Subset 3>
	game := Game{}
	parts := strings.Split(line, ":")
	game.ID, _ = strconv.Atoi(parts[0][5:])
	subsets := strings.Split(parts[1], ";")
	for _, subset := range subsets {
		game.Subsets = append(game.Subsets, parseSubset(subset))
	}
	return game
}

func parseInput(path string) []Game {
	lines, _ := util.ReadLines(path)
	games := lo.Map(lines, func(line string, _ int) Game {
		return parseGame(line)
	})
	return games
}

func A(path string) int {
	// max red = 12
	// max green = 13
	// max blue = 14
	sum := 0
	games := parseInput(path)
	for _, game := range games {
		possible := true
		for _, subset := range game.Subsets {
			if subset.Red > 12 {
				possible = false
			}
			if subset.Green > 13 {
				possible = false
			}
			if subset.Blue > 14 {
				possible = false
			}
		}
		if possible {
			sum += game.ID
		}
	}
	return sum
}

func B(path string) int {
	games := parseInput(path)
	sum := 0
	for _, game := range games {
		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, subset := range game.Subsets {
			minRed = max(minRed, subset.Red)
			minGreen = max(minGreen, subset.Green)
			minBlue = max(minBlue, subset.Blue)
		}
		power := minRed * minGreen * minBlue
		sum += power
	}
	return sum
}

func Run() {
	fmt.Println("Part A:", A("day02/input.txt"))
	fmt.Println("Part B:", B("day02/input.txt"))
}

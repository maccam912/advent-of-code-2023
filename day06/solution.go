package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Race struct {
	Time   int
	Record int
}

func parseInput(path string) []Race {
	lines, _ := util.ReadLines(path)
	times := strings.Fields(lines[0])
	records := strings.Fields(lines[1])
	races := []Race{}
	for i := 1; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		record, _ := strconv.Atoi(records[i])

		races = append(races, Race{time, record})
	}
	return races
}

func parseInputB(path string) []Race {
	lines, _ := util.ReadLines(path)
	times := strings.Fields(lines[0])
	records := strings.Fields(lines[1])
	time, _ := strconv.Atoi(strings.Join(times[1:], ""))
	record, _ := strconv.Atoi(strings.Join(records[1:], ""))

	races := []Race{}
	races = append(races, Race{time, record})
	return races
}

func scoreRace(race Race) int {
	score := 0
	for i := 0; i <= race.Time; i++ {
		dist := i * (race.Time - i)
		if dist > race.Record {
			score += 1
		}
	}
	return score
}

func A(path string) int {
	races := parseInput(path)
	total := 1
	for _, race := range races {
		score := scoreRace(race)
		total *= score
	}
	return total
}

func B(path string) int {
	races := parseInputB(path)
	total := 1
	for _, race := range races {
		score := scoreRace(race)
		total *= score
	}
	return total
}

// Run runs the day 1 challenge.
func Run() {
	partA := A("day06/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day06/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

// In main.go at the project root
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/maccam912/advent-of-code-2023/day01" // Import day packages
	"github.com/maccam912/advent-of-code-2023/day02"
	"github.com/maccam912/advent-of-code-2023/day03"
	"github.com/maccam912/advent-of-code-2023/day04"
	"github.com/maccam912/advent-of-code-2023/day05"
	"github.com/maccam912/advent-of-code-2023/day06"
	"github.com/maccam912/advent-of-code-2023/day07"
	"github.com/maccam912/advent-of-code-2023/day08"
	"github.com/maccam912/advent-of-code-2023/day09"
	"github.com/maccam912/advent-of-code-2023/day10"
	"github.com/maccam912/advent-of-code-2023/day11"
	"github.com/maccam912/advent-of-code-2023/day12"
	"github.com/maccam912/advent-of-code-2023/day13"
	"github.com/maccam912/advent-of-code-2023/day14"
	"github.com/maccam912/advent-of-code-2023/day15"
	"github.com/maccam912/advent-of-code-2023/day16"
	"github.com/maccam912/advent-of-code-2023/day17"
	"github.com/maccam912/advent-of-code-2023/day18"
	"github.com/maccam912/advent-of-code-2023/day19"
	"github.com/maccam912/advent-of-code-2023/day20"
	"github.com/maccam912/advent-of-code-2023/day21"
	"github.com/maccam912/advent-of-code-2023/day22"
	"github.com/maccam912/advent-of-code-2023/day23"
	"github.com/maccam912/advent-of-code-2023/day24"
	"github.com/maccam912/advent-of-code-2023/day25"
	// Import other days similarly
)

// Define a slice of functions for each day's solution
var days = []func(){
	day01.Run,
	day02.Run,
	day03.Run,
	day04.Run,
	day05.Run,
	day06.Run,
	day07.Run,
	day08.Run,
	day09.Run,
	day10.Run,
	day11.Run,
	day12.Run,
	day13.Run,
	day14.Run,
	day15.Run,
	day16.Run,
	day17.Run,
	day18.Run,
	day19.Run,
	day20.Run,
	day21.Run,
	day22.Run,
	day23.Run,
	day24.Run,
	day25.Run,
	// Add other days' functions here
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Running all days...")
		for i, dayFunc := range days {
			fmt.Printf("Day %d:\n", i+1)
			dayFunc() // Execute each day's solution
		}

		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day:", os.Args[1])

		return
	}

	if day < 1 || day > len(days) {
		fmt.Println("Day not implemented:", day)

		return
	}

	// Execute the solution for the specified day
	days[day-1]()
}

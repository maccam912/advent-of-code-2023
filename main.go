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

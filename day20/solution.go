package day20

import "fmt"

type Module struct {
	name string
}

func A(path string) int {
	return 0
}

func B(path string) int {
	return 0
}

func Run() {
	partA := A("dayXX/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("dayXX/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

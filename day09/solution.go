package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
	"github.com/samber/lo"
)

func parseInput(path string) [][]int {
	lines, _ := util.ReadLines(path)
	retval := make([][]int, 0)
	for _, line := range lines {
		nums := strings.Fields(line)
		ints := lo.Map(nums, func(s string, _ int) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		retval = append(retval, ints)
	}
	return retval
}

func diffs(nums []int) []int {
	retval := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		retval = append(retval, nums[i]-nums[i-1])
	}
	return retval
}

func checkAllZeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func findNext(nums []int) int {
	diffs := diffs(nums)
	if checkAllZeroes(diffs) {
		// Got to the last layer
		return nums[len(nums)-1]
	} else {
		// Use findNext on diffs, get last number, add to last number in nums
		nextDiff := findNext(diffs)
		return nextDiff + nums[len(nums)-1]
	}
}

func A(path string) int {
	nums := parseInput(path)
	sum := 0
	for _, list := range nums {
		val := findNext(list)
		sum += val
	}
	return sum
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func B(path string) int {
	nums := parseInput(path)
	// reverse all lists in nums
	for i := 0; i < len(nums); i++ {
		reverseSlice(nums[i])
	}

	sum := 0
	for _, list := range nums {
		val := findNext(list)
		sum += val
	}
	return sum
}

// Run runs the day 1 challenge.
func Run() {
	partA := A("day09/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day09/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

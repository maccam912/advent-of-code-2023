package util

import (
	"log"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	// read file from path, return utf-8 string of contents
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contents), "\n")

	return lines, nil
}

// gcd calculates the Greatest Common Divisor using Euclidean algorithm
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm calculates the Least Common Multiple of two numbers
func Lcm(a, b int) int {
	return a / Gcd(a, b) * b // Multiplication after division to prevent overflow
}

// lcmList calculates the LCM of a list of numbers
func LcmList(nums []int) int {
	if len(nums) < 2 {
		return 0 // LCM of less than two numbers is not well-defined
	}

	result := nums[0]
	for _, num := range nums[1:] {
		result = Lcm(result, num)
	}
	return result
}

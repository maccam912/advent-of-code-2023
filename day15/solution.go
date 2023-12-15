package day15

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func hash(s string) int {
	hash := 0
	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func parseInput(path string) []string {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(contents), ",")
}

func A(path string) int {
	steps := parseInput(path)
	hashes := lo.Map(steps, func(s string, _ int) int {
		return hash(s)
	})
	return lo.Sum(hashes)
}

type Lens struct {
	label       string
	focalLength int
}

type Box struct {
	lenses []Lens
}

func (box *Box) insertLens(lens Lens) {
	for i, l := range box.lenses {
		if l.label == lens.label {
			box.lenses[i] = lens
			return
		}
	}

	box.lenses = append(box.lenses, lens)
}

func (box *Box) removeLens(lens Lens) {
	for i, l := range box.lenses {
		if l.label == lens.label {
			box.lenses = append(box.lenses[:i], box.lenses[i+1:]...)
			return
		}
	}
}

func debugBoxes(boxes []*Box) {
	for i, box := range boxes {
		if len(box.lenses) != 0 {
			fmt.Printf("Box %d: ", i)
			for _, lens := range box.lenses {
				fmt.Printf("[%v %v] ", lens.label, lens.focalLength)
			}
			fmt.Printf("\n")
		}
	}
}

func B(path string) int {
	steps := parseInput(path)
	lenses := lo.Map(steps, func(s string, _ int) Lens {
		// check if '=' is in string s
		if strings.Contains(s, "=") {
			// split string s by '='
			split := strings.Split(s, "=")
			// convert string to int
			focalLength, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}
			return Lens{split[0], focalLength}
		} else {
			// must be a -, label is everything except last character
			label := s[:len(s)-1]
			return Lens{label, -1}
		}
	})

	// path is 256 boxes long
	boxes := make([]*Box, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = &Box{}
	}

	for _, lens := range lenses {
		boxNum := hash(lens.label)
		if lens.focalLength == -1 {
			boxes[boxNum].removeLens(lens)
		} else {
			boxes[boxNum].insertLens(lens)
		}
		// fmt.Printf("After %v:\n", lens.label)
		// debugBoxes(boxes)
		// fmt.Println("")
	}

	sum := 0
	for i, box := range boxes {
		for j, lens := range box.lenses {
			sum += (i + 1) * (j + 1) * lens.focalLength
		}
	}
	return sum
}

func Run() {
	partA := A("day15/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day15/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

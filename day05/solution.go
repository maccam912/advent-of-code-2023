package day05

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Range struct {
	Dst int
	Src int
	Len int
}

type Map struct {
	Ranges []Range
}

func (m Map) Get(src int) int {
	for _, r := range m.Ranges {
		if src >= r.Src && src < r.Src+r.Len {
			return r.Dst + (src - r.Src)
		}
	}
	return src
}

type Almanac struct {
	Seeds       []int
	Seed2Soil   Map
	Soil2Fert   Map
	Fert2Water  Map
	Water2Light Map
	Light2Temp  Map
	Temp2Humid  Map
	Humid2Loc   Map
}

func (a Almanac) GetLocation(seed int) int {
	soil := a.Seed2Soil.Get(seed)
	fert := a.Soil2Fert.Get(soil)
	water := a.Fert2Water.Get(fert)
	light := a.Water2Light.Get(water)
	temp := a.Light2Temp.Get(light)
	humid := a.Temp2Humid.Get(temp)
	loc := a.Humid2Loc.Get(humid)
	return loc
}

func parseMap(content string) Map {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	ranges := []Range{}
	for _, line := range lines[1:] {
		intStrings := strings.Split(strings.TrimSpace(line), " ")
		ints := lo.Map(intStrings, func(s string, _ int) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		ranges = append(ranges, Range{ints[0], ints[1], ints[2]})
	}
	return Map{ranges}
}

func parseInput(path string) Almanac {
	// Read in input file
	contents, _ := os.ReadFile(path)
	// First break into chunks when two newlines in a row
	chunks := strings.Split(string(contents), "\n\n")
	// Chunk 1 is just the list of seeds, i.e. seeds: 79 14 55 13
	seeds := strings.Split(strings.TrimSpace(chunks[0]), " ")
	seedInts := lo.Map(seeds[1:], func(s string, _ int) int {
		i, _ := strconv.Atoi(s)
		return i
	})

	return Almanac{seedInts, parseMap(chunks[1]), parseMap(chunks[2]), parseMap(chunks[3]), parseMap(chunks[4]), parseMap(chunks[5]), parseMap(chunks[6]), parseMap(chunks[7])}
}

func A(path string) int {
	almanac := parseInput(path)
	lowestLoc := 999999999
	for _, seed := range almanac.Seeds {
		loc := almanac.GetLocation(seed)
		if loc < lowestLoc {
			lowestLoc = loc
		}
	}
	return lowestLoc
}

func B(path string) int {
	almanac := parseInput(path)
	lowestLoc := 999999999
	for idx := 0; idx < len(almanac.Seeds); idx += 2 {
		fmt.Printf("Starting seed %d\n", idx)
		src := almanac.Seeds[idx]
		for i := 0; i < almanac.Seeds[idx+1]; i++ {
			seedNum := src + i
			loc := almanac.GetLocation(seedNum)
			if loc < lowestLoc {
				lowestLoc = loc
			}
		}
	}
	return lowestLoc
}

// Run runs the day 1 challenge.
func Run() {
	partA := A("day05/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day05/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

// type Cell struct {
// 	value  int
// 	exists bool
// }

type Lookup struct {
	tree map[string]int
}

func stringify(s string, groups []int) string {
	strs := make([]string, len(groups))
	for i, v := range groups {
		strs[i] = strconv.Itoa(v)
	}
	return s + strings.Join(strs, ",")
}

func NewLookup() *Lookup {
	return &Lookup{make(map[string]int)}
}

type Row struct {
	groups  []int
	row     []rune
	answers *Lookup
}

func (r *Row) IsValid() bool {
	numBroken := 0
	for i := 0; i < len(r.row); i++ {
		if r.row[i] == '#' {
			numBroken++
		}
	}
	sumBroken := 0
	for _, group := range r.groups {
		sumBroken += group
	}
	if numBroken > sumBroken {
		return false
	}
	discoveredGroups := []int{}
	currentGroup := 0
	for i := 0; i < len(r.row); i++ {

		if r.row[i] != '?' {
			if r.row[i] == '#' {
				currentGroup++
			} else {
				// must be .
				if currentGroup > 0 {
					discoveredGroups = append(discoveredGroups, currentGroup)
					currentGroup = 0
				}
			}
		} else {
			break
		}
		if i == len(r.row)-1 && r.row[i-1] != '?' {
			if currentGroup > 0 {
				discoveredGroups = append(discoveredGroups, currentGroup)
			}
			if len(discoveredGroups) != len(r.groups) {
				return false
			}
			for i, group := range r.groups {
				if group != discoveredGroups[i] {
					return false
				}
			}
			return true
		}
	}
	if len(discoveredGroups) > len(r.groups) {
		return false
	}
	for i, group := range discoveredGroups {
		if group != r.groups[i] {
			return false
		}
	}
	return true
}

func (r *Row) addAnswer(key string, value int) {
	r.answers.tree[key] = value
}

func (r *Row) getAnswer(key string) (int, bool) {
	val, exists := r.answers.tree[key]
	return val, exists
}

// func (r *Row) CountPossibilities() int {
// 	if !r.IsValid() {
// 		return 0
// 	}

// 	for i := 0; i < len(r.row); i++ {
// 		answer, exists := r.getAnswer(r.row[:i])
// 		if exists {
// 			return answer
// 		}

// 		if r.row[i] == '?' {
// 			r.row[i] = '#'
// 			possibilities := r.CountPossibilities()
// 			r.row[i] = '.'
// 			possibilities += r.CountPossibilities()
// 			r.row[i] = '?'
// 			r.addAnswer(r.row[:i], possibilities)
// 			return possibilities
// 		}
// 	}
// 	return 1
// }

func (r *Row) arrangements(springs string, group []int) int {
	if value, exists := r.getAnswer(stringify(springs, group)); exists {
		return value
	}

	if len(group) == 0 {
		// All groups already found, only allow where the rest of the string is . or ?
		for _, rune := range springs {
			if rune == '#' {
				r.addAnswer(stringify(springs, group), 0)
				return 0
			}
		}
		r.addAnswer(stringify(springs, group), 1)
		return 1
	}

	sum := 0
	for _, g := range group {
		sum += g
	}
	if len(springs) < sum {
		// There are not enough springs left to reach the count listed in groups
		r.addAnswer(stringify(springs, group), 0)
		return 0
	}

	if springs[0] == '.' {
		// Skip the known working ones
		answer := r.arrangements(springs[1:], group)
		r.addAnswer(stringify(springs, group), answer)
		return answer
	}

	// Either # or ?, try extracting a group
	here := 0
	if len(springs) > group[0] && springs[group[0]] != '#' {
		// This is a possible group!
		possible := true
		for i := 0; i < group[0]; i++ {
			if springs[i] == '.' {
				possible = false // No . allowed in this group
			}
		}
		if possible {
			here = r.arrangements(springs[group[0]+1:], group[1:]) // Get answer for remainder
		}
	}

	next := 0
	if springs[0] == '?' {
		next = r.arrangements(springs[1:], group)
	}
	r.addAnswer(stringify(springs, group), next+here)
	return next + here
}

func parseInput(path string) []Row {
	lines, _ := util.ReadLines(path)
	rows := make([]Row, 0, len(lines))
	for _, line := range lines {
		row := Row{}
		fields := strings.Fields(line)
		row.row = []rune(fields[0])
		groupNums := strings.Split(fields[1], ",")
		for _, groupNum := range groupNums {
			num, _ := strconv.Atoi(groupNum)
			row.groups = append(row.groups, num)
		}
		row.answers = NewLookup()
		rows = append(rows, row)
	}
	return rows
}

func A(path string) int {
	input := parseInput(path)
	possibilities := 0
	for _, row := range input {
		// possibilities += row.CountPossibilities()
		possibilities += row.arrangements(string(row.row)+".", row.groups)
	}
	return possibilities
}

func B(path string) int {
	input := parseInput(path)
	for i, row := range input {
		extension := []int{}
		for i := 0; i < 4; i++ {
			extension = append(extension, row.groups...)
		}
		input[i].groups = append(row.groups, extension...)
		runeExtension := []rune{}
		for i := 0; i < 4; i++ {
			runeExtension = append(runeExtension, '?')
			runeExtension = append(runeExtension, row.row...)
		}
		input[i].row = append(row.row, runeExtension...)
	}
	possibilities := 0
	for i, row := range input {
		fmt.Println(i)
		// possibilities += row.CountPossibilities()
		possibilities += row.arrangements(string(row.row)+".", row.groups)
	}
	return possibilities
}

func Run() {
	partA := A("day12/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day12/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

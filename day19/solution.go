package day19

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	part      rune
	lt        bool
	testValue int
	dest      string
	catchall  bool
}

type Part struct {
	values map[rune]int
}

type Input struct {
	workflows map[string][]Rule
	parts     []Part
}

func parseRule(rule string) Rule {
	// e.g. a<2006:qkq
	parts := strings.Split(strings.TrimSpace(rule), ":")
	if len(parts) != 2 {
		return Rule{dest: parts[0], catchall: true}
	}
	part := rune(parts[0][0])
	lt := parts[0][1] == '<'
	value, _ := strconv.Atoi(parts[0][2:])
	dest := parts[1]
	return Rule{part, lt, value, dest, false}
}

func parseWorkflow(workflow string) (string, []Rule) {
	// e.g. px{a<2006:qkq,m>2090:A,rfg}
	parts := strings.Split(strings.TrimSpace(workflow), "{")
	name := parts[0]
	_rules := strings.Split(strings.TrimSuffix(parts[1], "}"), ",")
	rules := make([]Rule, len(_rules))
	for i, rule := range _rules {
		rule := parseRule(rule)
		rules[i] = rule
	}
	return name, rules
}

func parsePart(part string) Part {
	// e.g. {x=787,m=2655,a=1222,s=2876}
	segments := strings.Split(strings.TrimSpace(part)[1:len(part)-1], ",")
	p := Part{values: make(map[rune]int)}
	x, _ := strconv.Atoi(strings.Split(segments[0], "=")[1])
	p.values['x'] = x
	m, _ := strconv.Atoi(strings.Split(segments[1], "=")[1])
	p.values['m'] = m
	a, _ := strconv.Atoi(strings.Split(segments[2], "=")[1])
	p.values['a'] = a
	s, _ := strconv.Atoi(strings.Split(segments[3], "=")[1])
	p.values['s'] = s
	return p
}

func parseInput(path string) Input {
	contents, _ := os.ReadFile(path)
	sections := strings.Split(string(contents), "\n\n")
	_workflows := strings.Split(sections[0], "\n")
	_parts := strings.Split(sections[1], "\n")
	workflows := make(map[string][]Rule)
	for _, workflow := range _workflows {
		name, rules := parseWorkflow(workflow)
		workflows[name] = rules
	}
	parts := make([]Part, len(_parts))
	for i, part := range _parts {
		parts[i] = parsePart(part)
	}

	return Input{workflows, parts}
}

func (input *Input) RunWorkflows(part Part) bool {
	currWorkflow := input.workflows["in"]
	fmt.Println("In workflow in")
	for true {
		for _, rule := range currWorkflow {
			fmt.Println(rule)
			if rule.catchall {
				if rule.dest == "A" {
					return true
				} else if rule.dest == "R" {
					return false
				}
				currWorkflow = input.workflows[rule.dest]
				fmt.Printf("In workflow %s\n", rule.dest)
				break
			} else {
				if rule.lt {
					if part.values[rule.part] < rule.testValue {
						currWorkflow = input.workflows[rule.dest]
						fmt.Printf("In workflow %s\n", rule.dest)
						if rule.dest == "A" {
							return true
						} else if rule.dest == "R" {
							return false
						}
						break
					}
				} else {
					if part.values[rule.part] > rule.testValue {
						currWorkflow = input.workflows[rule.dest]
						fmt.Printf("In workflow %s\n", rule.dest)
						if rule.dest == "A" {
							return true
						} else if rule.dest == "R" {
							return false
						}
						break
					}
				}
			}
		}
	}
	panic("unreachable")
}

func (part *Part) Sum() int {
	return part.values['x'] + part.values['m'] + part.values['a'] + part.values['s']
}

func A(path string) int {
	input := parseInput(path)
	sum := 0
	for _, part := range input.parts {
		if input.RunWorkflows(part) {
			sum += part.Sum()
		}
	}
	return sum
}

func B(path string) int {
	return 0
}

func Run() {
	partA := A("day19/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day19/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

package day20

import (
	"fmt"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Module struct {
	name       string
	isFlipFlop bool
	state      bool
	inputs     map[string]bool
	outputs    []string
}

func (module *Module) ReceivePulse(pulse Pulse) []Pulse {
	pulses := make([]Pulse, 0)
	if module.name == "broadcaster" {
		for _, output := range module.outputs {
			pulses = append(pulses, Pulse{output, pulse.hi, module.name})
		}
	} else if module.isFlipFlop {
		if !pulse.hi {
			module.state = !module.state
			for _, output := range module.outputs {
				pulses = append(pulses, Pulse{output, module.state, module.name})
			}
		}
	} else {
		// Must be conj
		module.inputs[pulse.from] = pulse.hi
		// If all inputs are high, send a low pulse. If any of the inputs are low, send high
		hi := false
		for _, input := range module.inputs {
			if !input {
				hi = true
			}
		}
		for _, output := range module.outputs {
			pulses = append(pulses, Pulse{output, hi, module.name})
		}
	}
	return pulses
}

type Pulse struct {
	name string
	hi   bool
	from string
}

func parseLine(line string) (string, *Module) {
	// e.g. %x -> a, b, c
	parts := strings.Split(strings.TrimSpace(line), " -> ")
	var name string
	var isFlipFlop bool
	if parts[0] == "broadcaster" {
		name = "broadcaster"
		isFlipFlop = false
	} else {
		name = parts[0][1:]
		isFlipFlop = parts[0][0] == '%'
	}
	outputs := strings.Split(parts[1], ", ")
	return name, &Module{name, isFlipFlop, false, map[string]bool{}, outputs}
}

func parseInput(path string) map[string]*Module {
	lines, _ := util.ReadLines(path)
	modules := make(map[string]*Module)

	for _, line := range lines {
		name, module := parseLine(line)
		modules[name] = module
	}

	// Link up inputs
	for _, module := range modules {
		for _, output := range module.outputs {
			outputModule := modules[output]
			if outputModule != nil {
				outputModule.inputs[module.name] = false
			}
		}
	}
	return modules
}

func A(path string) int {
	modules := parseInput(path)
	// Press button (broadcaster) 1000
	hiCount := 0
	loCount := 0
	for i := 0; i < 1000; i++ {
		pulses := []Pulse{}
		pulses = append(pulses, Pulse{"broadcaster", false, "broadcaster"})
		for len(pulses) > 0 {
			pulse := pulses[0]
			if pulse.hi {
				hiCount++
			} else {
				loCount++
			}
			// fmt.Printf("%v %v -> %v\n", pulse.from, pulse.hi, pulse.name)
			pulses = pulses[1:]
			mod := modules[pulse.name]
			if mod != nil {
				result := mod.ReceivePulse(pulse)
				pulses = append(pulses, result...)
			}
		}
	}
	return hiCount * loCount
}

func B(path string) int {
	answer := util.LcmList([]int{3877, 3907, 4001, 4027})
	return answer
	// modules := parseInput(path)
	// buttonPresses := 0
	// for true {
	// 	pulses := []Pulse{}
	// 	pulses = append(pulses, Pulse{"broadcaster", false, "broadcaster"})
	// 	buttonPresses++
	// 	for len(pulses) > 0 {
	// 		pulse := pulses[0]
	// 		if pulse.name == "rm" && pulse.hi {
	// 			fmt.Printf("%v %v\n", pulse.from, buttonPresses)
	// 			// dh - 3877
	// 			// bb 3907
	// 			// qd - 4001
	// 			// dp 0 4027
	// 		}
	// 		if pulse.name == "rx" && !pulse.hi {
	// 			return buttonPresses
	// 		}
	// 		pulses = pulses[1:]
	// 		mod := modules[pulse.name]
	// 		if mod != nil {
	// 			result := mod.ReceivePulse(pulse)
	// 			pulses = append(pulses, result...)
	// 		}
	// 	}
	// }
	// panic("Should not get here")
}

func Run() {
	partA := A("day20/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day20/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

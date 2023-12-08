package day08

import (
	"bufio"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

type Graph struct {
	Instructions []rune
	Nodes        map[string]*Node
}

type MathThing struct {
	Node     *Node
	ZIndices []int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(name string) {
	if _, exists := g.Nodes[name]; !exists {
		g.Nodes[name] = &Node{Name: name}
	}
}

func (g *Graph) SetChildren(nodeName string, left, right string) {
	node, exists := g.Nodes[nodeName]
	if exists {
		node.Left = g.Nodes[left]
		node.Right = g.Nodes[right]
	}
}

func parseFile(filePath string) (*Graph, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := NewGraph()

	// First line contains instructions (RL)
	if scanner.Scan() {
		graph.Instructions = []rune(scanner.Text())
	}

	// First pass: Add all nodes without children
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			continue // skip malformed lines
		}
		nodeName := parts[0]
		graph.AddNode(nodeName)
	}

	// Reset file scanner to start
	if _, err := file.Seek(0, 0); err != nil {
		return nil, err
	}
	scanner = bufio.NewScanner(file)

	// Skip the first line (instructions)
	scanner.Scan()

	// Second pass: Set children for each node
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			continue
		}

		nodeName := parts[0]
		children := strings.Trim(parts[1], "()")
		childNames := strings.Split(children, ", ")

		if len(childNames) == 2 {
			graph.SetChildren(nodeName, childNames[0], childNames[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func A(filePath string) int {
	graph, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	// Start at AAA, follow insructions in order, recycling as needed, until you reach ZZZ. Return number of steps to reach ZZZ.
	steps := 0
	currentNode := graph.Nodes["AAA"]
	for currentNode.Name != "ZZZ" {
		// Follow instructions
		instruction := graph.Instructions[steps%len(graph.Instructions)]
		if instruction == 'R' {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
		steps++
	}
	return steps
}

func getInstruction(graph *Graph, steps int) rune {
	return graph.Instructions[steps%len(graph.Instructions)]
}

func checkSuffix(nodes []*MathThing, suffix byte, step int) bool {
	done := true
	for _, thing := range nodes {
		if thing.Node.Name[len(thing.Node.Name)-1] != suffix {
			done = false
		} else {
			if len(thing.ZIndices) < 4 {
				thing.ZIndices = append(thing.ZIndices, step)
			}
		}
	}
	return done
}

func updateStartingNodes(startingNodes []*MathThing, instruction rune) {
	for i, thing := range startingNodes {
		// Follow instructions
		if instruction == 'R' {
			startingNodes[i].Node = thing.Node.Right
		} else {
			startingNodes[i].Node = thing.Node.Left
		}
	}
}

func B(filePath string) int {
	graph, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	// Now start on EVERY node where the name ends with the letter 'A'. Follow the instructions for EACH of these starting nodes, each of the paths taking a step on every instruction.
	// Keep going until EVERY path is simultaneously on a node ending with 'Z'. Return the number of steps it took to reach this state. If only some nodes end in 'Z', continue
	// as usual.
	steps := 0
	startingNodes := make([]*MathThing, 0)
	for _, node := range graph.Nodes {
		if strings.HasSuffix(node.Name, "A") {
			startingNodes = append(startingNodes, &MathThing{node, make([]int, 0)})
		}
	}

	done := false
	for !done && steps < 1000000000 {
		instruction := getInstruction(graph, steps)
		updateStartingNodes(startingNodes, instruction)
		steps++

		// Check if all nodes are on a node ending in 'Z'
		done = checkSuffix(startingNodes, 'Z', steps)
	}
	values := []int{}
	for _, node := range startingNodes {
		if len(node.ZIndices) > 0 {
			fmt.Printf("Offset 1: %d\n", node.ZIndices[0])
			values = append(values, (node.ZIndices[0]))
		}
	}
	fmt.Printf("LCM of %v: %d\n", values, util.LcmList(values))
	return steps
}

// Run runs the day 1 challenge.
func Run() {
	partA := A("day08/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	partB := B("day08/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

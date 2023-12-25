package day25

import (
	"fmt"
	"os"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Node struct {
	name  string
	edges []*Node
}

type Graph struct {
	nodes map[string]*Node
}

func parseInput(path string) *Graph {
	lines, _ := util.ReadLines(path)
	graph := &Graph{nodes: map[string]*Node{}}
	for _, line := range lines {
		// Line is <name>: <name> <name> <name>
		fields := strings.Fields(strings.TrimSpace(line))
		src := fields[0][:len(fields[0])-1] // remove the colon
		graph.nodes[src] = &Node{name: src, edges: []*Node{}}
		for _, dst := range fields[1:] {
			graph.nodes[dst] = &Node{name: dst, edges: []*Node{}}
		}
	}
	// Now connect
	for _, line := range lines {
		fields := strings.Fields(strings.TrimSpace(line))
		src := fields[0][:len(fields[0])-1] // remove the colon
		for _, dst := range fields[1:] {
			graph.nodes[src].edges = append(graph.nodes[src].edges, graph.nodes[dst])
		}
	}
	file, _ := os.Create("C:\\Users\\macca\\Downloads\\graph.dot")
	defer file.Close()
	for _, node := range graph.nodes {
		for _, edge := range node.edges {
			line := fmt.Sprintf("%s -> %s;\n", node.name, edge.name)
			file.WriteString(line)
		}
	}
	return graph
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

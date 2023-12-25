package day25

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
)

type Edge struct {
	src, dst string
}

type Graph struct {
	nodes map[string]bool
	edges []Edge
}

func (graph *Graph) addEdge(src, dst string) {
	// order alphabetically since it is undirected and we dont want duplicates
	if src < dst {
		graph.edges = append(graph.edges, Edge{src, dst})
	} else {
		graph.edges = append(graph.edges, Edge{dst, src})
	}
}

func (graph *Graph) removeDuplicateEdges() {
	for i := 0; i < len(graph.edges); i++ {
		for j := i + 1; j < len(graph.edges); j++ {
			if graph.edges[i] == graph.edges[j] {
				graph.edges = append(graph.edges[:j], graph.edges[j+1:]...)
				j--
			}
		}
	}
}

func parseInput(path string) *Graph {
	lines, _ := util.ReadLines(path)
	graph := &Graph{nodes: map[string]bool{}, edges: []Edge{}}
	for _, line := range lines {
		// Line is <name>: <name> <name> <name>
		fields := strings.Fields(strings.TrimSpace(line))
		src := fields[0][:len(fields[0])-1] // remove the colon
		graph.nodes[src] = true
		for _, dst := range fields[1:] {
			graph.nodes[dst] = true
			graph.addEdge(src, dst)
		}
	}
	graph.removeDuplicateEdges()
	return graph
}

func findSet(parents map[string]string, node string) string {
	if parents[node] == node {
		return node
	}
	return findSet(parents, parents[node])
}

func unionSet(parent map[string]string, x, y string) {
	xset := findSet(parent, x)
	yset := findSet(parent, y)
	parent[xset] = yset
}

func (g *Graph) KargerMinCut() (int, map[string][]string) {
	minCut := len(g.edges)
	var finalPartitions map[string][]string
	// Iterate several times to improve accuracy
	for minCut > 3 {
		parent := make(map[string]string)
		for _, e := range g.edges {
			parent[e.src] = e.src
			parent[e.dst] = e.dst
		}

		v := len(parent)
		for v > 2 {
			e := g.edges[rand.Intn(len(g.edges))]
			subset1 := findSet(parent, e.src)
			subset2 := findSet(parent, e.dst)
			if subset1 != subset2 {
				unionSet(parent, subset1, subset2)
				v--
			}
		}

		cutEdges := 0
		partitions := make(map[string][]string)
		for _, e := range g.edges {
			subset1 := findSet(parent, e.src)
			subset2 := findSet(parent, e.dst)

			if subset1 != subset2 {
				cutEdges++
			} else {
				partitions[subset1] = append(partitions[subset1], e.src)
				partitions[subset1] = append(partitions[subset1], e.dst)
			}
		}

		if cutEdges < minCut {
			minCut = cutEdges
			finalPartitions = partitions
		}
	}

	// remove duplicate vertices in each partition
	for key := range finalPartitions {
		vertexMap := make(map[string]bool)
		uniqueVertices := []string{}
		for _, vertex := range finalPartitions[key] {
			if _, exists := vertexMap[vertex]; !exists {
				vertexMap[vertex] = true
				uniqueVertices = append(uniqueVertices, vertex)
			}
		}
		finalPartitions[key] = uniqueVertices
	}
	return minCut, finalPartitions
}

func A(path string) int {
	graph := parseInput(path)
	_, partitions := graph.KargerMinCut()
	prod := 1
	for _, vertices := range partitions {
		prod *= len(vertices)
	}
	return prod
}

func B(path string) int {
	return 0
}

func Run() {
	partA := A("day25/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day25/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

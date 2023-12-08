package day08

import (
	"testing"
)

func TestParseFileSimplified(t *testing.T) {
	graph, err := parseFile("example_input.txt")
	if err != nil {
		t.Fatalf("ParseFile returned an error: %v", err)
	}

	// Check instructions
	expectedInstructions := []rune{'R', 'L'}
	if string(graph.Instructions) != string(expectedInstructions) {
		t.Errorf("Instructions do not match expected output. Got %s, want %s", string(graph.Instructions), string(expectedInstructions))
	}

	// Check specific nodes
	testCases := []struct {
		nodeName   string
		leftChild  string
		rightChild string
	}{
		{"AAA", "BBB", "CCC"},
		{"BBB", "DDD", "EEE"},
		// Add more nodes if needed
	}

	for _, tc := range testCases {
		node, ok := graph.Nodes[tc.nodeName]
		if !ok {
			t.Errorf("Node %s not found", tc.nodeName)
			continue
		}
		if node.Left == nil || node.Left.Name != tc.leftChild {
			t.Errorf("Left child of %s is incorrect. Got %v, want %s", tc.nodeName, node.Left, tc.leftChild)
		}
		if node.Right == nil || node.Right.Name != tc.rightChild {
			t.Errorf("Right child of %s is incorrect. Got %v, want %s", tc.nodeName, node.Right, tc.rightChild)
		}
	}
}

func TestA(t *testing.T) {
	expected := 2
	actual := A("example_input.txt")
	if actual != expected {
		t.Errorf("A() = %d, want %d", actual, expected)
	}

	expected = 6
	actual = A("example_input_2.txt")
	if actual != expected {
		t.Errorf("A() = %d, want %d", actual, expected)
	}
}

func TestB(t *testing.T) {
	expected := 6
	actual := B("example_input_3.txt")
	if actual != expected {
		t.Errorf("B() = %d, want %d", actual, expected)
	}
}

package day01

import (
	"testing"
	// Import any additional packages needed for testing
)

func TestA(t *testing.T) {
	expected := 142
	actual := A("example_input.txt") // This assumes Run returns some value to test against

	if actual != expected {
		t.Errorf("Run() = %v, want %v", actual, expected)
	}
}

func TestB(t *testing.T) {
	expected := 281
	actual := B("example_input_2.txt") // This assumes Run returns some value to test against

	if actual != expected {
		t.Errorf("Run() = %v, want %v", actual, expected)
	}
}

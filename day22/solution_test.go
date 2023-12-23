package day22

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	bricks := parseInput("example_input.txt")
	if len(bricks) != 7 {
		t.Errorf("Expected 3 bricks, got %d", len(bricks))
	}
	if bricks[0].a.x != 1 {
		t.Errorf("Expected 1, got %d", bricks[0].a.x)
	}
	if bricks[0].a.y != 0 {
		t.Errorf("Expected 0, got %d", bricks[0].a.y)
	}
	if bricks[0].a.z != 1 {
		t.Errorf("Expected 1, got %d", bricks[0].a.z)
	}
	if bricks[0].b.x != 1 {
		t.Errorf("Expected 1, got %d", bricks[0].b.x)
	}
	if bricks[0].b.y != 2 {
		t.Errorf("Expected 2, got %d", bricks[0].b.y)
	}
	if bricks[0].b.z != 1 {
		t.Errorf("Expected 1, got %d", bricks[0].b.z)
	}
}

func TestIsSupported(t *testing.T) {
	a := Brick{a: Coord3{x: 0, y: 0, z: 2}, b: Coord3{x: 3, y: 3, z: 2}}
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			b := Brick{a: Coord3{x: x, y: y, z: 1}, b: Coord3{x: x, y: y, z: 1}}
			bricks := []*Brick{&a, &b}
			if !a.IsSupported(bricks) {
				t.Errorf("Expected true, got false")
			}
		}
	}
	a = Brick{a: Coord3{x: 1, y: 1, z: 2}, b: Coord3{x: 1, y: 1, z: 2}}
	b := Brick{a: Coord3{x: 1, y: 1, z: 1}, b: Coord3{x: 2, y: 2, z: 1}}
	if !a.IsSupported([]*Brick{&b}) {
		t.Errorf("Expected true, got false")
	}
	c := Brick{a: Coord3{x: 2, y: 1, z: 1}, b: Coord3{x: 2, y: 2, z: 1}}
	if a.IsSupported([]*Brick{&c}) {
		t.Errorf("Expected false, got true")
	}
	if !a.IsSupported([]*Brick{&b, &c}) {
		t.Errorf("Expected true, got false")
	}
}

func TestSettle(t *testing.T) {
	bricks := parseInput("example_input.txt")
	answer := SettleBricks(bricks)
	if answer != 5 {
		t.Errorf("Expected 5, got %d", answer)
	}
}

func TestA(t *testing.T) {
	answer := A("example_input.txt")
	if answer != 5 {
		t.Errorf("Expected 5, got %d", answer)
	}
}

func TestB(t *testing.T) {
	answer := B("example_input.txt")
	if answer != 7 {
		t.Errorf("Expected 7, got %d", answer)
	}
}

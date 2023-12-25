package day24

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
	"github.com/samber/lo"
)

type Hailstone struct {
	px, py, pz int
	vx, vy, vz int
}

type Sky struct {
	hailstones []*Hailstone
}

func (s *Sky) CheckIntersections(lo, hi int) int {
	count := 0
	for i := 0; i < len(s.hailstones); i++ {
		for j := i + 1; j < len(s.hailstones); j++ {
			if s.intersects(s.hailstones[i], s.hailstones[j], lo, hi) {
				count++
			}
		}
	}
	return count
}

func (s *Sky) intersects(a, b *Hailstone, lo, hi int) bool {
	// Calculate determinants
	detA := a.vx*b.vy - a.vy*b.vx
	if detA == 0 {
		// The determinant is zero if the lines are parallel (or coincident)
		return false
	}

	detTA := (b.px-a.px)*b.vy - (b.py-a.py)*b.vx
	detTB := a.vx*(b.py-a.py) - a.vy*(b.px-a.px)

	// Calculate tA and tB
	tA := float64(detTA) / float64(detA)
	tB := float64(detTB) / float64(detA)
	// if tA < 0 || tA-tB < 0 {
	if tA < 0 || tB > 0 {
		// It happened in the past
		return false
	}

	// Check if the intersection point is within the test area
	intersectX := float64(a.px) + tA*float64(a.vx)
	intersectY := float64(a.py) + tA*float64(a.vy)
	fmt.Println(tB)

	return intersectX >= float64(lo) && intersectX <= float64(hi) && intersectY >= float64(lo) && intersectY <= float64(hi)
}

func (s *Sky) parallel(a, b *Hailstone) bool {
	// Do this in 3D now
	return a.vx*b.vy == a.vy*b.vx && a.vx*b.vz == a.vz*b.vx && a.vy*b.vz == a.vz*b.vy
}

func (s *Sky) FindParallelHailstones() []*Hailstone {
	parallel := make([]*Hailstone, 0)
	for i := 0; i < len(s.hailstones); i++ {
		for j := i + 1; j < len(s.hailstones); j++ {
			if s.parallel(s.hailstones[i], s.hailstones[j]) {
				parallel = append(parallel, s.hailstones[i], s.hailstones[j])
			}
		}
	}
	return parallel
}

func parseInput(path string) Sky {
	lines, _ := util.ReadLines(path)
	sky := Sky{hailstones: make([]*Hailstone, 0)}

	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), " @ ")
		_pos := strings.Split(parts[0], ", ")
		_vel := strings.Split(parts[1], ", ")
		pos := lo.Map(_pos, func(s string, _ int) int {
			parsed, _ := strconv.Atoi(strings.TrimSpace(s))
			return parsed
		})
		vel := lo.Map(_vel, func(s string, _ int) int {
			parsed, _ := strconv.Atoi(strings.TrimSpace(s))
			return parsed
		})
		hailstone := Hailstone{px: pos[0], py: pos[1], pz: pos[2], vx: vel[0], vy: vel[1], vz: vel[2]}
		sky.hailstones = append(sky.hailstones, &hailstone)
	}

	return sky
}

func (s *Sky) TranslateToFirstHailstoneReference() *Sky {
	if len(s.hailstones) == 0 {
		return &Sky{hailstones: []*Hailstone{}}
	}

	// Get the first hailstone's position and velocity
	first := s.hailstones[0]
	refPx, refPy, refPz := first.px, first.py, first.pz
	refVx, refVy, refVz := first.vx, first.vy, first.vz

	// Create a new slice to hold the translated hailstones
	translatedHailstones := make([]*Hailstone, len(s.hailstones))

	for i, h := range s.hailstones {
		// Translate position and velocity
		newPx := h.px - refPx
		newPy := h.py - refPy
		newPz := h.pz - refPz
		newVx := h.vx - refVx
		newVy := h.vy - refVy
		newVz := h.vz - refVz

		// Create a new Hailstone with translated values
		translatedHailstones[i] = &Hailstone{
			px: newPx, py: newPy, pz: newPz,
			vx: newVx, vy: newVy, vz: newVz,
		}
	}

	// Return a new Sky with the translated hailstones
	return &Sky{hailstones: translatedHailstones}
}

func A(path string) int {
	sky := parseInput(path)
	count := sky.CheckIntersections(200000000000000, 400000000000000)
	return count
}

func elim(m [][]*big.Rat) [][]*big.Rat {
	size := len(m)
	for i := 0; i < size; i++ {
		t := new(big.Rat).Set(m[i][i])
		for k := range m[i] {
			m[i][k].Quo(m[i][k], t)
		}
		for j := i + 1; j < size; j++ {
			t := new(big.Rat).Set(m[j][i])
			for k := range m[j] {
				m[j][k].Sub(m[j][k], new(big.Rat).Mul(t, m[i][k]))
			}
		}
	}
	for i := size - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			t := new(big.Rat).Set(m[j][i])
			for k := range m[j] {
				m[j][k].Sub(m[j][k], new(big.Rat).Mul(t, m[i][k]))
			}
		}
	}
	return m
}

func B(path string) int {
	sky := parseInput(path)

	// Assuming you have at least 5 hailstones
	if len(sky.hailstones) < 5 {
		log.Fatal("Not enough hailstones")
	}

	// Create a 4x5 matrix to hold the coefficients of the system of equations
	m := make([][]*big.Rat, 4)
	for i := range m {
		m[i] = make([]*big.Rat, 5)
		for j := range m[i] {
			m[i][j] = new(big.Rat)
		}
	}

	// Fill the matrix with coefficients from the first 4 hailstones
	for i, hailstone := range sky.hailstones[:4] {
		m[i][0].SetInt64(int64(hailstone.vx - sky.hailstones[4].vx))
		m[i][1].SetInt64(int64(hailstone.vy - sky.hailstones[4].vy))
		m[i][2].SetInt64(int64(sky.hailstones[4].py - hailstone.py))
		m[i][3].SetInt64(int64(hailstone.px - sky.hailstones[4].px))
		m[i][4].Set(
			new(big.Rat).Sub(
				new(big.Rat).Mul(new(big.Rat).SetInt64(int64(hailstone.px)), new(big.Rat).SetInt64(int64(sky.hailstones[4].vy))),
				new(big.Rat).Mul(new(big.Rat).SetInt64(int64(hailstone.py)), new(big.Rat).SetInt64(int64(sky.hailstones[4].vx))),
			),
		)
	}

	// Perform Gaussian elimination on the matrix
	m = elim(m)

	// Extract the solution from the last column of the matrix
	solution := make([]big.Rat, 4)
	for i := range solution {
		solution[i] = *m[i][4]
	}

	fmt.Println(solution)
	return 0
}

func Run() {
	partA := A("day24/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day24/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

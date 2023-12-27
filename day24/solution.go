package day24

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/maccam912/advent-of-code-2023/util"
	"github.com/samber/lo"
	"gonum.org/v1/gonum/mat"
)

type Hailstone struct {
	px, py, pz float64
	vx, vy, vz float64
}

func CrossProduct[T int | float64](a, b []T) []T {
	return []T{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

type Sky struct {
	hailstones []*Hailstone
}

func (sky *Sky) GenerateSystemOfEquations() (*mat.Dense, *mat.Dense) {
	hailstones := sky.hailstones
	numRows := 3 * len(hailstones)
	A := mat.NewDense(numRows, 6, nil)
	B := mat.NewDense(numRows, 1, nil)

	for i, hailstone := range hailstones {
		pDiff := []float64{-hailstone.px, -hailstone.py, -hailstone.pz} // Difference in position (rock - hailstone)
		vDiff := []float64{-hailstone.vx, -hailstone.vy, -hailstone.vz} // Difference in velocity (rock - hailstone)

		cross := CrossProduct(pDiff, vDiff)

		for j := 0; j < 3; j++ {
			A.Set(3*i+j, j, vDiff[j])    // Coefficients for position variables
			A.Set(3*i+j, 3+j, -pDiff[j]) // Coefficients for velocity variables
			B.Set(3*i+j, 0, cross[j])
		}
	}
	return A, B
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
		hailstone := Hailstone{px: float64(pos[0]), py: float64(pos[1]), pz: float64(pos[2]), vx: float64(vel[0]), vy: float64(vel[1]), vz: float64(vel[2])}
		sky.hailstones = append(sky.hailstones, &hailstone)
	}

	return sky
}

// func createMatrix(sky Sky, numRows int) [][]int {
// 	matrix := make([][]int, numRows)

// 	for i := 0; i < numRows; i++ {
// 		hailstone := sky.hailstones[i]

// 		matrix[i] = []int{
// 			-dy,
// 			dx,
// 			y,
// 			-x,
// 			y*dx - x*dy,
// 		}
// 	}

// 	return matrix
// }

// func convertToBigRat(matrix [][]int) [][]*big.Rat {
// 	rows := len(matrix)
// 	cols := len(matrix[0])
// 	ratMatrix := make([][]*big.Rat, rows)
// 	for i := range ratMatrix {
// 		ratMatrix[i] = make([]*big.Rat, cols)
// 		for j := range ratMatrix[i] {
// 			ratMatrix[i][j] = new(big.Rat).SetInt(big.NewInt(int64(matrix[i][j])))
// 		}
// 	}
// 	return ratMatrix
// }

func gaussianElimination(matrix [][]*big.Rat) [][]*big.Rat {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		// Find the pivot for the column
		pivotRow := i
		for j := i + 1; j < rows; j++ {
			a := matrix[j][i]
			b := matrix[pivotRow][i]
			absB := matrix[pivotRow][i].Abs(b)
			cmp := matrix[j][i].Cmp(absB)

			if a != nil && b != nil && cmp > 0 {
				pivotRow = j
			}
		}

		// Swap the current row with the pivot row
		matrix[i], matrix[pivotRow] = matrix[pivotRow], matrix[i]

		// Ensure pivot is non-nil and non-zero
		if matrix[i][i] == nil || matrix[i][i].Sign() == 0 {
			continue // Skip this column
		}

		// Normalize the pivot row
		pivot := new(big.Rat).Set(matrix[i][i])
		for k := 0; k < cols; k++ {
			if matrix[i][k] != nil {
				matrix[i][k].Quo(matrix[i][k], pivot)
			} else {
				matrix[i][k] = new(big.Rat) // Initialize to 0
			}
		}

		// Eliminate all other elements in the current column
		for j := 0; j < rows; j++ {
			if j != i && matrix[j][i] != nil {
				factor := new(big.Rat).Set(matrix[j][i])
				for k := 0; k < cols; k++ {
					if matrix[i][k] != nil {
						term := new(big.Rat).Mul(factor, matrix[i][k])
						if matrix[j][k] == nil {
							matrix[j][k] = new(big.Rat) // Initialize to 0
						}
						matrix[j][k].Sub(matrix[j][k], term)
					}
				}
			}
		}
	}

	return matrix
}

func A(path string) int {
	sky := parseInput(path)
	count := sky.CheckIntersections(200000000000000, 400000000000000)
	return count
}

// SolveSystem solves the system of equations Ax = B.
func SolveSystem(A, B *mat.Dense) mat.Dense {
	var x mat.Dense
	if err := x.Solve(A, B); err != nil {
		fmt.Println("Failed to solve the system:", err)
	}
	return x
}

func B(path string) int {
	sky := parseInput(path)
	A, B := sky.GenerateSystemOfEquations()
	solved := SolveSystem(A, B)
	fmt.Printf("%v\n\n%v\n", mat.Formatted(A), mat.Formatted(B))
	fmt.Printf("Solved: %v\n", mat.Formatted(&solved))
	return 0
}

func Run() {
	partA := A("day24/input.txt")
	fmt.Printf("Part A: %v\n", partA)

	partB := B("day24/input.txt")
	fmt.Printf("Part B: %v\n", partB)
}

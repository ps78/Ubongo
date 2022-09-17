package main

import "fmt"

// Vector represents a 3-dimensional int-vector
type Vector [3]int

type Vectorf [3]float64

// Returns a String representation of a vector
func (v Vector) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v[0], v[1], v[2])
}

func (v Vectorf) String() string {
	return fmt.Sprintf("(%f,%f,%f)", v[0], v[1], v[2])
}

func (a Vector) Float64() Vectorf {
	return Vectorf{float64(a[0]), float64(a[1]), float64(a[2])}
}

func (a Vectorf) Add(b Vectorf) Vectorf {
	return Vectorf{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func (a Vectorf) Sub(b Vectorf) Vectorf {
	return Vectorf{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a Vectorf) Div(b float64) Vectorf {
	return Vectorf{a[0] / b, a[1] / b, a[2] / b}
}

func (a Vectorf) Mult(b float64) Vectorf {
	return Vectorf{a[0] * b, a[1] * b, a[2] * b}
}

// GetShiftVectors returns all possible placements of the inner bounding box
// inside the outer bounding box.
// Returns an empty slice if inner does not fit into outer at all
func (outerBoundingBox Vector) GetShiftVectors(innerBoundingBox Vector) []Vector {
	delta := Vector{
		outerBoundingBox[0] - innerBoundingBox[0],
		outerBoundingBox[1] - innerBoundingBox[1],
		outerBoundingBox[2] - innerBoundingBox[2]}

	n := (delta[0] + 1) * (delta[1] + 1) * (delta[2] + 1)

	// return empty vector if there is not fit
	if n <= 0 {
		return make([]Vector, 0)
	} else {
		shifts := make([]Vector, n)
		i := 0
		for x := 0; x <= delta[0]; x++ {
			for y := 0; y <= delta[1]; y++ {
				for z := 0; z <= delta[2]; z++ {
					shifts[i] = [3]int{x, y, z}
					i++
				}
			}
		}
		return shifts
	}
}

// Looks for a in lst (comparing the values)
func FindArray3d(lst []*Array3d, a *Array3d) (bool, int) {
	if lst != nil && a != nil {
		for i, arr := range lst {
			if a.IsEqual(arr) {
				return true, i
			}
		}
	}
	return false, -1
}

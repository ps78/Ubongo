package main

import "fmt"

// Vector represents a 3-dimensional int-vector
type Vector [3]int

// Returns a String representation of a vector
func (v Vector) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v[0], v[1], v[2])
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

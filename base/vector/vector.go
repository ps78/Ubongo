package vector

import (
	"fmt"
	"ubongo/base/vectorf"
)

// V represents a 3-dimensional int-vector
type V [3]int

func (a V) AsVectorf() vectorf.V {
	return vectorf.V{float64(a[0]), float64(a[1]), float64(a[2])}
}

func (a V) String() string {
	return fmt.Sprintf("(%d,%d,%d)", a[0], a[1], a[2])
}

func (a V) Max() int {
	if a[0] >= a[1] && a[0] >= a[2] {
		return a[0]
	} else if a[1] >= a[0] && a[1] >= a[2] {
		return a[1]
	} else {
		return a[2]
	}
}

func (a V) Add(b V) V {
	return V{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func (a V) Sub(b V) V {
	return V{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a V) Mult(b int) V {
	return V{a[0] * b, a[1] * b, a[2] * b}
}

func (a V) Div(b float64) vectorf.V {
	return vectorf.V{float64(a[0]) / b, float64(a[1]) / b, float64(a[2]) / b}
}

func (a V) Flip() V {
	return V{-a[0], -a[1], -a[2]}
}

// GetShiftVectors returns all possible placements of the inner bounding box
// inside the outer bounding box.
// Returns an empty slice if inner does not fit into outer at all
func (outerBoundingBox V) GetShiftVectors(innerBoundingBox V) []V {
	delta := V{
		outerBoundingBox[0] - innerBoundingBox[0],
		outerBoundingBox[1] - innerBoundingBox[1],
		outerBoundingBox[2] - innerBoundingBox[2]}

	n := (delta[0] + 1) * (delta[1] + 1) * (delta[2] + 1)

	// return empty vector if there is not fit
	if n <= 0 {
		return make([]V, 0)
	} else {
		shifts := make([]V, n)
		i := 0
		for x := 0; x <= delta[0]; x++ {
			for y := 0; y <= delta[1]; y++ {
				for z := 0; z <= delta[2]; z++ {
					shifts[i] = V{x, y, z}
					i++
				}
			}
		}
		return shifts
	}
}

// Package vectorf contains the type vectorf and related methods
package vectorf

import (
	"fmt"
	"math"
)

// V represents a 3-dimensional float-vector
type V [3]float64

// Zero represents the zero-vector {0,0,0}
var Zero V = V{0.0, 0.0, 0.0}

// String returns a readable string representation of a
func (a V) String() string {
	return fmt.Sprintf("(%f,%f,%f)", a[0], a[1], a[2])
}

// Max returns the biggest element in a
func (a V) Max() float64 {
	return math.Max(math.Max(a[0], a[1]), a[2])
}

// Add adds vector b to a and returns the result
func (a V) Add(b V) V {
	return V{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

// Sub subtracts vector b from a and returns the result
func (a V) Sub(b V) V {
	return V{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

// Mult multiplies vector a by b and returns the result
func (a V) Mult(b float64) V {
	return V{a[0] * b, a[1] * b, a[2] * b}
}

// Div divides each element of vector a by b and returns the result
func (a V) Div(b float64) V {
	return V{a[0] / b, a[1] / b, a[2] / b}
}

// Flip multiplies each element of vector a by -1
func (a V) Flip() V {
	return V{-a[0], -a[1], -a[2]}
}

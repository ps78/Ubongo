// Package array2d contains the declaration of 2D-array of int8 elements and
// associated methods.
package array2d

import (
	"fmt"
	"ubongo/base/array3d"
)

// A is a 2-dimensional array representing the area of a problem,
// where 0 indicates that the unit square is part of the shape, and -1 is not
type A struct {
	// data contains the actual values, to be access with Get() / Set() methods
	// This array has size [DimX][DimY]
	data [][]int8

	// DimX is the size of the first dimension of the array
	DimX int

	// DimY is hte size of the second dimension of the array
	DimY int
}

// String returns a readable version of the array
func (a *A) String() string {
	if a == nil {
		return "(nil)"
	} else {
		return fmt.Sprintf("<%d-%d>%v", a.DimX, a.DimY, a.data)
	}
}

// New creates a new zeroed 2D array with the given dimensions
// Panics if any of the dimensions are smaller than 1
func New(dimX, dimY int) *A {
	if dimX <= 0 || dimY <= 0 {
		panic("Cannot initialize a array2d with dimensions smaller than 1")
	}

	a := make([][]int8, dimX)
	for i := 0; i < dimX; i++ {
		a[i] = make([]int8, dimY)
	}
	return &A{data: a, DimX: dimX, DimY: dimY}
}

// NewFromData creates a new 2D array from the given data
// Returns nil if data is nil
func NewFromData(data [][]int8) *A {
	if data == nil {
		return nil
	} else {
		dimX := len(data)
		dimY := len(data[0])
		a := make([][]int8, dimX)
		for i := 0; i < dimX; i++ {
			a[i] = make([]int8, dimY)
			copy(a[i], data[i])
		}
		return &A{data: a, DimX: dimX, DimY: dimY}
	}
}

// Get returns element [x][y] of the 2D array.
// Invalid indices will create an exception
func (a *A) Get(x, y int) int8 {
	return a.data[x][y]
}

// Set sets the element [x][y] of the 2D array
// Invalid indices will create an exception
func (a *A) Set(x, y int, value int8) {
	a.data[x][y] = value
}

// Extrude creates a 3D array based on the 2D array A by extruding it by height-steps into the 3rd dimension,
// copying the values to each level. Panics if height is smaller than 1
func (arr *A) Extrude(height int) *array3d.A {
	if height < 1 {
		panic("Cannot extrude array2d with a height smaller than 1")
	}

	a := array3d.New(arr.DimX, arr.DimY, height)
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < height; z++ {
				a.Set(x, y, z, arr.Get(x, y))
			}
		}
	}
	return a
}

// Equal returns true if the dimensions and all elements of a and b are equal
// Accepts nil arguments
func (a *A) IsEqual(b *A) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if a.DimX != b.DimX || a.DimY != b.DimY {
		return false
	} else {
		for x := 0; x < a.DimX; x++ {
			for y := 0; y < a.DimY; y++ {
				if a.Get(x, y) != b.Get(x, y) {
					return false
				}
			}
		}
	}
	return true
}

// Clone creates a deep copy of the given 2D array
func (src *A) Clone() *A {
	if src == nil {
		return nil
	} else {
		cp := New(src.DimX, src.DimY)
		for x := 0; x < src.DimX; x++ {
			for y := 0; y < src.DimY; y++ {
				cp.Set(x, y, src.Get(x, y))
			}
		}
		return cp
	}
}

// Count counts the elements in a that equal lookFor
func (a *A) Count(lookFor int8) int {
	count := 0
	if a != nil {
		for x := 0; x < a.DimX; x++ {
			for y := 0; y < a.DimY; y++ {
				if a.Get(x, y) == lookFor {
					count++
				}
			}
		}
	}
	return count
}

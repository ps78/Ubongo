package utils

import (
	"fmt"
	"github.com/ps17/ubongo/utils"

// Array2d is a 2-dimensional array representing the area of a problem,
// where 0 indicates that the unit square is part of the shape, and -1 is not
type Array2d struct {
	data [][]int8
	DimX int
	DimY int
}

func (a *Array2d) String() string {
	return fmt.Sprintf("<%d-%d>%v", a.DimX, a.DimY, a.data)
}

// Creates a new zeroed 2D array
func NewArray2d(dimX, dimY int) *Array2d {
	a := make([][]int8, dimX)
	for i := 0; i < dimX; i++ {
		a[i] = make([]int8, dimY)
	}
	return &Array2d{data: a, DimX: dimX, DimY: dimY}
}

// Creates a new 2D array from data
func NewArray2dFromData(data [][]int8) *Array2d {
	dimX := len(data)
	dimY := len(data[0])
	a := make([][]int8, dimX)
	for i := 0; i < dimX; i++ {
		a[i] = make([]int8, dimY)
		copy(a[i], data[i])
	}
	return &Array2d{data: a, DimX: dimX, DimY: dimY}
}

// Returns element [x][y] of the 2D array
func (a *Array2d) Get(x, y int) int8 {
	return a.data[x][y]
}

// Sets the element [x][y] of the 2D array
func (a *Array2d) Set(x, y int, value int8) {
	a.data[x][y] = value
}

// Extrudes the given 2D array by height-steps into the 3rd dimension,
// copying the values to each level
func (arr *Array2d) Extrude(height int) *Array3d {
	a := NewArray3d(arr.DimX, arr.DimY, height)
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < height; z++ {
				a.Set(x, y, z, arr.Get(x, y))
			}
		}
	}
	return a
}

// Tests if the 2D arrays a and b contain the same elements
func (a *Array2d) IsEqual(b *Array2d) bool {
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

// Creates a copy of the given 2D array
func (src *Array2d) Clone() *Array2d {
	cp := NewArray2d(src.DimX, src.DimY)
	for x := 0; x < src.DimX; x++ {
		for y := 0; y < src.DimY; y++ {
			cp.Set(x, y, src.Get(x, y))
		}
	}
	return cp
}

// Counts the elements in arr that equal lookFor
func (a *Array2d) Count(lookFor int8) int {
	count := 0
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			if a.Get(x, y) == lookFor {
				count++
			}
		}
	}
	return count
}

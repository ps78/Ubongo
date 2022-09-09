package main

import "fmt"

// Array3d is a 3-dimensional array of int8 representing the shape of a block or the
// volume of a game, where 1 inidicates the presence of a unit cube and 0 absence of one
// the value -1 is used to define a cube that is outside of the game-volume
type Array3d [][][]int8

// Array2d is a 2-dimensional array representing the area of a problem,
// where 0 indicates that the unit square is part of the shape, and -1 is not
type Array2d [][]int8

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

// Creates a [xdim]x[ydim] array of int8 and returns a slice to it
func Make2DArray(xdim int, ydim int) [][]int8 {
	a := make([][]int8, xdim)
	for i := 0; i < xdim; i++ {
		a[i] = make([]int8, ydim)
	}
	return a
}

// Creates a [xdim]x[ydim]x[zdim] array of int8 and returns a slice to it
func Make3DArray(xdim int, ydim int, zdim int) Array3d {
	a := make([][][]int8, xdim)
	for i := 0; i < xdim; i++ {
		a[i] = make([][]int8, ydim)
		for j := 0; j < ydim; j++ {
			a[i][j] = make([]int8, zdim)
		}
	}
	return a
}

// Extrudes the given 2D array by height-steps into the 3rd dimension,
// copying the values to each level
func Extrude2DArray(shape Array2d, height int) Array3d {
	xdim := len(shape)
	ydim := len(shape[0])
	a := Make3DArray(xdim, ydim, height)
	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {
			for z := 0; z < height; z++ {
				a[x][y][z] = shape[x][y]
			}
		}
	}
	return a
}

// Tests if the 2D arrays a and b contain the same elements
func Equal2DArray(a Array2d, b Array2d) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if len(a) != len(b) {
		return false
	} else if len(a[0]) != len(b[0]) {
		return false
	} else {
		for x := 0; x < len(a); x++ {
			for y := 0; y < len(a[x]); y++ {
				if a[x][y] != b[x][y] {
					return false
				}
			}
		}
	}
	return true
}

// Tests if the 3D array a and b contain the same elements
func Equal3DArray(a Array3d, b Array3d) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if len(a) != len(b) {
		return false
	} else if len(a[0]) != len(b[0]) {
		return false
	} else if len(a[0][0]) != len(b[0][0]) {
		return false
	} else {
		for x := 0; x < len(a); x++ {
			for y := 0; y < len(a[x]); y++ {
				for z := 0; z < len(a[x][y]); z++ {
					if a[x][y][z] != b[x][y][z] {
						return false
					}
				}
			}
		}
	}
	return true
}

// Creates a copy of the given 2D array
func Copy2DArray(src Array2d) Array2d {
	cp := make([][]int8, len(src))
	for i := range src {
		cp[i] = make([]int8, len(src[i]))
		copy(cp[i], src[i])
	}
	return cp
}

// Creates a copy of the given 3D array
func Copy3DArray(src Array3d) Array3d {
	cp := make([][][]int8, len(src))
	for i := range src {
		cp[i] = make([][]int8, len(src[i]))
		for j := range src[i] {
			cp[i][j] = make([]int8, len(src[i][j]))
			copy(cp[i][j], src[i][j])
		}
	}
	return cp
}

// Counts the elements in arr that equal lookFor
func CountValues2D(arr Array2d, lookFor int8) int {
	count := 0
	for x, b := range arr {
		for y := range b {
			if arr[x][y] == lookFor {
				count++
			}
		}
	}
	return count
}

// Counts the elements in arr that equal lookFor
func CountValues3D(arr Array3d, lookFor int8) int {
	count := 0
	for x, b := range arr {
		for y, c := range b {
			for z := range c {
				if arr[x][y][z] == lookFor {
					count++
				}
			}
		}
	}
	return count
}

// GetBoundBoxSize returns the dimensions of the given volume
// which correspond to the size of the bounding box
func GetBoundingBoxFromArray(shape Array3d) Vector {
	xdim := len(shape)
	ydim := len(shape[0])
	zdim := len(shape[0][0])
	return Vector{xdim, ydim, zdim}
}

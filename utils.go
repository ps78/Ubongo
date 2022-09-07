package main

import "fmt"

// Array3d is a 3-dimensional array of int8 representing the shape of a block or the
// volume of a game, where 1 inidicates the presence of a unit cube and 0 absence of one
// the value -1 is used to define a cube that is outside of the game-volume
type Array3d [][][]int8

// Array2d is a 2-dimensional array representing the area of a problem,
// where 0 indicates that the unit square is part of the shape, and -1 is not
type Array2d [][]int8

// BoundingBox describes a box with 3 dimensions expressed in unit-squares
type BoundingBox [3]int

// Vector represents a 3-dimensional int-vector
type Vector [3]int

func (v Vector) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v[0], v[1], v[2])
}

// GetShiftVectors returns all possible placements of the inner bounding box
// inside the outer bounding box.
// Returns an empty slice if inner does not fit into outer at all
func (outer BoundingBox) GetShiftVectors(inner BoundingBox) []Vector {
	delta := []int{outer[0] - inner[0], outer[1] - inner[1], outer[2] - inner[2]}
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

// tries to add the given blockShape to the volume, at the position defined by shift
// return true and the updated volume on success, false, nil otherwise
func TryAdd(vol Array3d, block Array3d, shift Vector) (bool, Array3d) {
	result := Copy3DArray(vol)
	box := GetBoundingBoxFromBlockShape(block)
	for x := shift[0]; x < box[0]+shift[0]; x++ {
		for y := shift[1]; y < box[1]+shift[1]; y++ {
			for z := shift[2]; z < box[2]+shift[2]; z++ {
				// only run following test if the block-cube is solid at the current position
				if block[x-shift[0]][y-shift[1]][z-shift[2]] == 1 {
					// if space is part of volume and empty -> ok
					if (vol)[x][y][z] == 0 {
						(result)[x][y][z] = 1 // mark space as occupied
					} else {
						return false, nil // otherwise abort
					}
				}
			}
		}
	}
	return true, result
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

// Checks if the the given volume is full, i.e. all spaces are either 0 or 2
func IsSolved(vol Array3d) bool {
	xdim, ydim, zdim := len(vol), len(vol[0]), len(vol[0][0])
	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {
			for z := 0; z < zdim; z++ {
				if vol[x][y][z] == 1 {
					return false
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

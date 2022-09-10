package main

// Array2d is a 2-dimensional array representing the area of a problem,
// where 0 indicates that the unit square is part of the shape, and -1 is not
type Array2d struct {
	data [][]int8
	DimX int
	DimY int
}

// Array3d is a 3-dimensional array of int8 representing the shape of a block or the
// volume of a game, where 1 inidicates the presence of a unit cube and 0 absence of one
// the value -1 is used to define a cube that is outside of the game-volume
type Array3d struct {
	data [][][]int8
	DimX int
	DimY int
	DimZ int
}

// Creates a new zeroed 2D array
func NewArray2d(dimX int, dimY int) *Array2d {
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

// Creates a new zeroed 3D array
func NewArray3d(xdim int, ydim int, zdim int) *Array3d {
	a := make([][][]int8, xdim)
	for i := 0; i < xdim; i++ {
		a[i] = make([][]int8, ydim)
		for j := 0; j < ydim; j++ {
			a[i][j] = make([]int8, zdim)
		}
	}
	return &Array3d{data: a, DimX: xdim, DimY: ydim, DimZ: zdim}
}

// Creates a new 3D array from data
func NewArray3dFromData(data [][][]int8) *Array3d {
	dimX := len(data)
	dimY := len(data[0])
	dimZ := len(data[0][0])
	a := make([][][]int8, dimX)
	for i := 0; i < dimX; i++ {
		a[i] = make([][]int8, dimY)
		for j := 0; j < dimY; j++ {
			a[i][j] = make([]int8, dimZ)
			copy(a[i][j], data[i][j])
		}
	}
	return &Array3d{data: a, DimX: dimX, DimY: dimY, DimZ: dimZ}
}

// Returns element [x][y] of the 2D array
func (a *Array2d) Get(x int, y int) int8 {
	return a.data[x][y]
}

// Sets the element [x][y] of the 2D array
func (a *Array2d) Set(x int, y int, value int8) {
	a.data[x][y] = value
}

// Returns element [x][y][z] of the 3D array
func (a *Array3d) Get(x int, y int, z int) int8 {
	return a.data[x][y][z]
}

// Sets the element [x][y][z] of the 3D array
func (a *Array3d) Set(x int, y int, z int, value int8) {
	a.data[x][y][z] = value
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

// Tests if the 3D array a and b contain the same elements
func (a *Array3d) IsEqual(b *Array3d) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if a.DimX != b.DimX || a.DimY != b.DimY || a.DimZ != b.DimZ {
		return false
	} else {
		for x := 0; x < a.DimX; x++ {
			for y := 0; y < a.DimY; y++ {
				for z := 0; z < a.DimZ; z++ {
					if a.Get(x, y, z) != b.Get(x, y, z) {
						return false
					}
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

// Creates a copy of the given 3D array
func (src *Array3d) Clone() *Array3d {
	cp := NewArray3d(src.DimX, src.DimY, src.DimZ)
	for x := 0; x < src.DimX; x++ {
		for y := 0; y < src.DimY; y++ {
			for z := 0; z < src.DimZ; z++ {
				cp.Set(x, y, z, src.Get(x, y, z))
			}
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

// Counts the elements in arr that equal lookFor
func (a *Array3d) Count(lookFor int8) int {
	count := 0
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < a.DimZ; z++ {
				if a.Get(x, y, z) == lookFor {
					count++
				}
			}
		}
	}
	return count
}

// GetBoundBoxSize returns the dimensions of the given volume
// which correspond to the size of the bounding box
func (a *Array3d) GetBoundingBox() Vector {
	return Vector{a.DimX, a.DimY, a.DimZ}
}

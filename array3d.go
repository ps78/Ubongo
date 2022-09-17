package main

import "fmt"

// Array3d is a 3-dimensional array of int8 representing the shape of a block or the
// volume of a game, where 1 inidicates the presence of a unit cube and 0 absence of one
// the value -1 is used to define a cube that is outside of the game-volume
type Array3d struct {
	data [][][]int8
	DimX int
	DimY int
	DimZ int
}

func (a *Array3d) String() string {
	return fmt.Sprintf("<%d-%d-%d>%v", a.DimX, a.DimY, a.DimZ, a.data)
}

// A function that can be applied to an Array3d
type Array3dFunc func(x, y, z int, currentValue int8) int8

// Creates a new zeroed 3D array
func NewArray3d(dimX, dimY, dimZ int) *Array3d {
	a := make([][][]int8, dimX)
	for i := 0; i < dimX; i++ {
		a[i] = make([][]int8, dimY)
		for j := 0; j < dimY; j++ {
			a[i][j] = make([]int8, dimZ)
		}
	}
	return &Array3d{data: a, DimX: dimX, DimY: dimY, DimZ: dimZ}
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

// Returns element [x][y][z] of the 3D array
func (a *Array3d) Get(x, y, z int) int8 {
	return a.data[x][y][z]
}

// Sets the element [x][y][z] of the 3D array
func (a *Array3d) Set(x, y, z int, value int8) {
	a.data[x][y][z] = value
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

// Applies the function f to each element of the array
func (a *Array3d) Apply(f Array3dFunc) *Array3d {
	r := NewArray3d(a.DimX, a.DimY, a.DimZ)
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < a.DimZ; z++ {
				r.Set(x, y, z, f(x, y, z, a.Get(x, y, z)))
			}
		}
	}
	return r
}

// Applies f to all elements and returns true if all return true
func (a *Array3d) AllTrue(f func(x, y, z int, currentValue int8) bool) bool {
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < a.DimZ; z++ {
				if !f(x, y, z, a.Get(x, y, z)) {
					return false
				}
			}
		}
	}
	return true
}

// Rotates the array around the z-axis counter-clockwise by 90°
func (a *Array3d) RotateZ() *Array3d {
	r := NewArray3d(a.DimY, a.DimX, a.DimZ)
	for x := 0; x < r.DimX; x++ {
		for y := 0; y < r.DimY; y++ {
			for z := 0; z < r.DimZ; z++ {
				r.Set(x, y, z, a.Get(y, a.DimY-x-1, z))
			}
		}
	}
	return r
}

func (a *Array3d) RotateZ2() *Array3d {
	return a.RotateZ().RotateZ()
}

func (a *Array3d) RotateZ3() *Array3d {
	return a.RotateZ().RotateZ().RotateZ()
}

// Rotates the array around the y-axis counter-clockwise by 90°
func (a *Array3d) RotateY() *Array3d {
	r := NewArray3d(a.DimZ, a.DimY, a.DimX)
	for x := 0; x < r.DimX; x++ {
		for y := 0; y < r.DimY; y++ {
			for z := 0; z < r.DimZ; z++ {
				r.Set(x, y, z, a.Get(a.DimX-z-1, y, x))
			}
		}
	}
	return r
}

func (a *Array3d) RotateY2() *Array3d {
	return a.RotateY().RotateY()
}

func (a *Array3d) RotateY3() *Array3d {
	return a.RotateY().RotateY().RotateY()
}

// Rotates the array around the x-axis counter-clockwise by 90°
func (a *Array3d) RotateX() *Array3d {
	r := NewArray3d(a.DimX, a.DimZ, a.DimY)
	for x := 0; x < r.DimX; x++ {
		for y := 0; y < r.DimY; y++ {
			for z := 0; z < r.DimZ; z++ {
				r.Set(x, y, z, a.Get(x, a.DimY-z-1, y))
			}
		}
	}
	return r
}

func (a *Array3d) RotateX2() *Array3d {
	return a.RotateX().RotateX()
}

func (a *Array3d) RotateX3() *Array3d {
	return a.RotateX().RotateX().RotateX()
}

// calculates the center of gravity of the given array
// relative to the origin which is in the corner of the array
// element [0,0,0]
func (a *Array3d) GetCenterOfGravity() Vectorf {
	var x, y, z, n float64
	for i := 0; i < a.DimX; i++ {
		for j := 0; j < a.DimY; j++ {
			for k := 0; k < a.DimZ; k++ {
				if a.Get(i, j, k) == IS_BLOCK {
					n += 1
					x += float64(i) + 0.5
					y += float64(j) + 0.5
					z += float64(k) + 0.5
				}
			}
		}
	}
	return Vectorf{x / n, y / n, z / n}
}

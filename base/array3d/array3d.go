package array3d

import (
	"fmt"
	"ubongo/base/vector"
	"ubongo/base/vectorf"
)

// A is a 3-dimensional array of int8 representing the shape of a block or the
// volume of a game, where 1 inidicates the presence of a unit cube and 0 absence of one
// the value -1 is used to define a cube that is outside of the game-volume
type A struct {
	data [][][]int8
	DimX int
	DimY int
	DimZ int
}

func (a *A) String() string {
	return fmt.Sprintf("<%d-%d-%d>%v", a.DimX, a.DimY, a.DimZ, a.data)
}

// A function that can be applied to an Array3d
type ApplyFunc func(x, y, z int, currentValue int8) int8

// Creates a new zeroed 3D array
func New(dimX, dimY, dimZ int) *A {
	a := make([][][]int8, dimX)
	for i := 0; i < dimX; i++ {
		a[i] = make([][]int8, dimY)
		for j := 0; j < dimY; j++ {
			a[i][j] = make([]int8, dimZ)
		}
	}
	return &A{data: a, DimX: dimX, DimY: dimY, DimZ: dimZ}
}

// Creates a new 3D array from data
func NewFromData(data [][][]int8) *A {
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
	return &A{data: a, DimX: dimX, DimY: dimY, DimZ: dimZ}
}

// Returns element [x][y][z] of the 3D array
func (a *A) Get(x, y, z int) int8 {
	return a.data[x][y][z]
}

// Sets the element [x][y][z] of the 3D array
func (a *A) Set(x, y, z int, value int8) {
	a.data[x][y][z] = value
}

// Tests if the 3D array a and b contain the same elements
func (a *A) IsEqual(b *A) bool {
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
func (src *A) Clone() *A {
	cp := New(src.DimX, src.DimY, src.DimZ)
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
func (a *A) Count(lookFor int8) int {
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
func (a *A) GetBoundingBox() vector.V {
	return vector.V{a.DimX, a.DimY, a.DimZ}
}

// Applies the function f to each element of the array
func (a *A) Apply(f ApplyFunc) *A {
	r := New(a.DimX, a.DimY, a.DimZ)
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
func (a *A) AllTrue(f func(x, y, z int, currentValue int8) bool) bool {
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
func (a *A) RotateZ() *A {
	r := New(a.DimY, a.DimX, a.DimZ)
	for x := 0; x < r.DimX; x++ {
		for y := 0; y < r.DimY; y++ {
			for z := 0; z < r.DimZ; z++ {
				r.Set(x, y, z, a.Get(y, a.DimY-x-1, z))
			}
		}
	}
	return r
}

func (a *A) RotateZ2() *A {
	return a.RotateZ().RotateZ()
}

func (a *A) RotateZ3() *A {
	return a.RotateZ().RotateZ().RotateZ()
}

// Rotates the array around the y-axis counter-clockwise by 90°
func (a *A) RotateY() *A {
	r := New(a.DimZ, a.DimY, a.DimX)
	for x := 0; x < r.DimX; x++ {
		for y := 0; y < r.DimY; y++ {
			for z := 0; z < r.DimZ; z++ {
				r.Set(x, y, z, a.Get(a.DimX-z-1, y, x))
			}
		}
	}
	return r
}

func (a *A) RotateY2() *A {
	return a.RotateY().RotateY()
}

func (a *A) RotateY3() *A {
	return a.RotateY().RotateY().RotateY()
}

// Rotates the array around the x-axis counter-clockwise by 90°
func (a *A) RotateX() *A {
	r := New(a.DimX, a.DimZ, a.DimY)
	for x := 0; x < r.DimX; x++ {
		for y := 0; y < r.DimY; y++ {
			for z := 0; z < r.DimZ; z++ {
				r.Set(x, y, z, a.Get(x, a.DimY-z-1, y))
			}
		}
	}
	return r
}

func (a *A) RotateX2() *A {
	return a.RotateX().RotateX()
}

func (a *A) RotateX3() *A {
	return a.RotateX().RotateX().RotateX()
}

// calculates the center of gravity of the given array
// relative to the origin which is in the corner of the array
// element [0,0,0]
func (a *A) GetCenterOfGravity() vectorf.V {
	var x, y, z, n float64
	for i := 0; i < a.DimX; i++ {
		for j := 0; j < a.DimY; j++ {
			for k := 0; k < a.DimZ; k++ {
				if a.Get(i, j, k) == 1 {
					n += 1
					x += float64(i) + 0.5
					y += float64(j) + 0.5
					z += float64(k) + 0.5
				}
			}
		}
	}
	return vectorf.V{x / n, y / n, z / n}
}

// Looks for a in lst (comparing the values)
func Find(lst []*A, a *A) (bool, int) {
	if lst != nil && a != nil {
		for i, arr := range lst {
			if a.IsEqual(arr) {
				return true, i
			}
		}
	}
	return false, -1
}

// Creates all 90° rotations of the base 3d array along the x, y and z axis
// A maximum of 24 arrays are returned, but identical rotations are removed,
// hence the number can be smaller (depending on symmetries of the base array)
func (base *A) CreateRotations() []*A {
	arr := make([]*A, 0)

	// helper function that adds el to lst if it is not already in lst
	addIfNotInList := func(lst []*A, el *A) []*A {
		for _, a := range lst {
			if a.IsEqual(el) {
				return lst
			}
		}
		return append(lst, el)
	}

	// the following code generates all possible rotations about 90° along the x, y, z axis for an
	// object in space, in the general case. Some rotations might be identical due to symmetries
	// of the object and will be eliminated
	arr = append(arr, base)
	arr = addIfNotInList(arr, base.RotateZ())
	arr = addIfNotInList(arr, base.RotateZ2())
	arr = addIfNotInList(arr, base.RotateZ3())

	arr = addIfNotInList(arr, base.RotateX())
	arr = addIfNotInList(arr, base.RotateX().RotateZ())
	arr = addIfNotInList(arr, base.RotateX().RotateZ2())
	arr = addIfNotInList(arr, base.RotateX().RotateZ3())

	arr = addIfNotInList(arr, base.RotateX2())
	arr = addIfNotInList(arr, base.RotateX2().RotateZ())
	arr = addIfNotInList(arr, base.RotateX2().RotateZ2())
	arr = addIfNotInList(arr, base.RotateX2().RotateZ3())

	arr = addIfNotInList(arr, base.RotateX3())
	arr = addIfNotInList(arr, base.RotateX3().RotateZ())
	arr = addIfNotInList(arr, base.RotateX3().RotateZ2())
	arr = addIfNotInList(arr, base.RotateX3().RotateZ3())

	arr = addIfNotInList(arr, base.RotateY())
	arr = addIfNotInList(arr, base.RotateY().RotateZ())
	arr = addIfNotInList(arr, base.RotateY().RotateZ2())
	arr = addIfNotInList(arr, base.RotateY().RotateZ3())

	arr = addIfNotInList(arr, base.RotateY3())
	arr = addIfNotInList(arr, base.RotateY3().RotateZ())
	arr = addIfNotInList(arr, base.RotateY3().RotateZ2())
	arr = addIfNotInList(arr, base.RotateY3().RotateZ3())

	return arr
}

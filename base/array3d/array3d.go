// Package array3d contains the declaration of 3D-array of int8 elements and
// associated methods.
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
	// data contains the actual values, to be access with Get() / Set() methods
	// This array has size [DimX][DimY][DimZ]
	data [][][]int8

	// DimX is the size of the first dimension of the array
	DimX int

	// DimY is hte size of the second dimension of the array
	DimY int

	// DimZ is hte size of the third dimension of the array
	DimZ int
}

// ApplyFunc is a function type that can be applied to an array3d.A with Apply()
type ApplyFunc func(x, y, z int, currentValue int8) int8

// String returns a readable version of the array
func (a *A) String() string {
	if a == nil {
		return "(nil)"
	} else {
		return fmt.Sprintf("<%d-%d-%d>%v", a.DimX, a.DimY, a.DimZ, a.data)
	}
}

// New creates a new zeroed 3D array
// Panics if any of the dimensions are smaller than 1
func New(dimX, dimY, dimZ int) *A {
	if dimX <= 0 || dimY <= 0 || dimZ <= 0 {
		panic("Cannot initialize a array3d with dimensions smaller than 1")
	}

	a := make([][][]int8, dimX)
	for i := 0; i < dimX; i++ {
		a[i] = make([][]int8, dimY)
		for j := 0; j < dimY; j++ {
			a[i][j] = make([]int8, dimZ)
		}
	}
	return &A{data: a, DimX: dimX, DimY: dimY, DimZ: dimZ}
}

// NewFromData creates a new 3D array from the given data
// Returns nil if data is nil
func NewFromData(data [][][]int8) *A {
	if data == nil {
		return nil
	} else {
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
}

// Get returns element [x][y][z] of the 3D array
// Invalid indices will create an exception
func (a *A) Get(x, y, z int) int8 {
	return a.data[x][y][z]
}

// Set sets the element [x][y][z] of the 3D array
// Invalid indices will create an exception
func (a *A) Set(x, y, z int, value int8) {
	a.data[x][y][z] = value
}

// Equals tests if the 3D array a and b contain the same elements
func (a *A) Equals(b *A) bool {
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

// Clone creates a deep copy of the given 3D array
func (src *A) Clone() *A {
	if src == nil {
		return nil
	} else {
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
}

// Count counts the elements in a that equal lookFor
func (a *A) Count(lookFor int8) int {
	count := 0
	if a != nil {
		for x := 0; x < a.DimX; x++ {
			for y := 0; y < a.DimY; y++ {
				for z := 0; z < a.DimZ; z++ {
					if a.Get(x, y, z) == lookFor {
						count++
					}
				}
			}
		}
	}
	return count
}

// GetBoundingBox returns the dimensions of the given volume
// which correspond to the size of the bounding box
// Applying this on an nil-instance returns {0,0,0}
func (a *A) GetBoundingBox() vector.V {
	if a == nil {
		return vector.Zero
	} else {
		return vector.V{a.DimX, a.DimY, a.DimZ}
	}
}

// Apply applies the function f to each element of the array, replacing
// the previous value
func (a *A) Apply(f ApplyFunc) *A {
	if a == nil || f == nil {
		return a
	} else {
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
}

// AllTrue return true if the result of applying function f to each element
// of the array is always true. If the array is nil, it returns false
func (a *A) AllTrue(f func(x, y, z int, currentValue int8) bool) bool {
	if a == nil || f == nil {
		return false
	} else {
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
}

// RotateZ rotates the array around the z-axis counter-clockwise by 90°
func (a *A) RotateZ() *A {
	if a == nil {
		return a
	} else {
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
}

// RotateZ2 rotates the array around the z-axis by 180°
func (a *A) RotateZ2() *A {
	return a.RotateZ().RotateZ()
}

// RotateZ3 rotates the array around the z-axis counter-clockwise by 270°
func (a *A) RotateZ3() *A {
	return a.RotateZ().RotateZ().RotateZ()
}

// RotateY rotates the array around the y-axis counter-clockwise by 90°
func (a *A) RotateY() *A {
	if a == nil {
		return a
	} else {
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
}

// RotateY2 rotates the array around the y-axis by 180°
func (a *A) RotateY2() *A {
	return a.RotateY().RotateY()
}

// RotateZ3 rotates the array around the y-axis counter-clockwise by 270°
func (a *A) RotateY3() *A {
	return a.RotateY().RotateY().RotateY()
}

// RotateX rotates the array around the x-axis counter-clockwise by 90°
func (a *A) RotateX() *A {
	if a == nil {
		return a
	} else {
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
}

// RotateX2 rotates the array around the x-axis by 180°
func (a *A) RotateX2() *A {
	return a.RotateX().RotateX()
}

// RotateX3 rotates the array around the x-axis counter-clockwise by 270°
func (a *A) RotateX3() *A {
	return a.RotateX().RotateX().RotateX()
}

// GetCenterOfGravity calculates the center of gravity of the given array
// relative to the origin which is in the corner of the array element [0,0,0]
// Applying this on a nil reference returns {0,0,0}
func (a *A) GetCenterOfGravity() vectorf.V {
	if a == nil {
		return vectorf.Zero
	} else {
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
}

// Find looks for a in lst (comparing the values)
// The first return value indicates if a was found in lst or not
// The second return value is the index of a in lst, or -1 if it was not found
func Find(lst []*A, a *A) (bool, int) {
	if a == nil || len(lst) == 0 {
		return false, -1
	} else {
		if lst != nil && a != nil {
			for i, arr := range lst {
				if a.Equals(arr) {
					return true, i
				}
			}
		}
		return false, -1
	}
}

// CreateRotations creates all 90° rotations of the base 3d array along the x, y and z axis
// A maximum of 24 arrays are returned, but identical rotations are removed,
// hence the number can be smaller (depending on symmetries of the base array)
// Applying this on a nil reference returns an empty array
func (base *A) CreateRotations() []*A {
	arr := make([]*A, 0)

	// helper function that adds el to lst if it is not already in lst
	addIfNotInList := func(lst []*A, el *A) []*A {
		for _, a := range lst {
			if a.Equals(el) {
				return lst
			}
		}
		return append(lst, el)
	}

	if base != nil {
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
	}

	return arr
}

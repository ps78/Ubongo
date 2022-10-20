package utils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArray3d(t *testing.T) {
	dimX := 2
	dimY := 3
	dimZ := 1

	a := NewArray3d(dimX, dimY, dimZ)

	assert.Equal(t, dimX, a.DimX, "Wrong x-dimension")
	assert.Equal(t, dimY, a.DimY, "Wrong y-dimension")
	assert.Equal(t, dimZ, a.DimZ, "Wrong z-dimension")

	for x := 0; x < dimX; x++ {
		for y := 0; y < dimY; y++ {
			for z := 0; z < dimZ; z++ {
				assert.Equal(t, int8(0), a.Get(x, y, z), "Make3DArray returned an array that is not zeroed at all positions")
			}
		}
	}
}

func TestNewArray3dFromData(t *testing.T) {
	data := [][][]int8{{{0, 1}, {1, 2}, {2, 3}}, {{1, 2}, {2, 3}, {3, 4}}, {{2, 3}, {3, 4}, {4, 5}}}
	a := NewArray3dFromData(data)

	assert.Equal(t, len(data), a.DimX, "Wrong x-dimension")
	assert.Equal(t, len(data[0]), a.DimY, "Wrong y-dimension")
	assert.Equal(t, len(data[0][0]), a.DimZ, "Wrong z-dimension")

	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < a.DimZ; z++ {
				data[x][y][z] = 0 // this should not affect a!
				assert.Equal(t, int8(x+y+z), a.Get(x, y, z), "NewArray3dFromData returned an array that did not properly copy data")
			}
		}
	}
}

func TestArray3dString(t *testing.T) {
	a := NewArray3dFromData([][][]int8{{{0}, {1}, {0}}, {{1}, {0}, {1}}})
	exp := "<2-3-1>[[[0] [1] [0]] [[1] [0] [1]]]"
	act := a.String()
	assert.Equal(t, exp, act)
}

func TestArray3dGetSet(t *testing.T) {
	a := NewArray3d(7, 9, 3)
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < a.DimZ; z++ {
				a.Set(x, y, z, int8(x+y-z))
				assert.Equal(t, int8(x+y-z), a.Get(x, y, z), "Get/Set at position (%d,%d,%d) did not work", x, y, z)
			}
		}
	}
}

func TestArray3dIsEqual(t *testing.T) {
	a := NewArray3dFromData([][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 5}}})
	b := NewArray3dFromData([][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 5}}}) // == a
	c := NewArray3dFromData([][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 0}}}) // != a
	d := NewArray3dFromData([][][]int8{{{2}, {3}}})
	var e *Array3d = nil

	assert.True(t, a.IsEqual(b), "Array a and b are equal but Equal3DArray reports they are not")
	assert.False(t, a.IsEqual(c), "Array a and c are not equal but Equal3DArray reports they are")
	assert.False(t, a.IsEqual(d), "Array a and d have different dimensions but Equal3DArray reports they are equal")
	assert.True(t, e.IsEqual(nil))
	assert.False(t, e.IsEqual(a))
	assert.False(t, a.IsEqual(e))
}

func TestArray3dClone(t *testing.T) {
	orig := NewArray3dFromData([][][]int8{{{0, 1}, {1, 2}}, {{1, 2}, {2, 3}}, {{2, 3}, {3, 4}}})
	copy := orig.Clone()
	orig.Set(0, 0, 0, 42) // this should not affect the copy

	assert.True(t, copy.DimX == orig.DimX && copy.DimY == orig.DimY && copy.DimZ == orig.DimZ,
		"Dimensions do not match")

	for x := 0; x < orig.DimX; x++ {
		for y := 0; y < orig.DimY; y++ {
			for z := 0; z < orig.DimZ; z++ {
				assert.Equal(t, int8(x+y+z), copy.Get(x, y, z), "Element [%d][%d][%d] does not match", x, y, z)
			}
		}
	}
}

func TestArray3dCount(t *testing.T) {
	a := NewArray3dFromData([][][]int8{{{1, 3}, {6, 1}}, {{3, 3}, {0, 9}}, {{1, 1}, {9, 0}}})
	assert.Equal(t, 2, a.Count(0))
	assert.Equal(t, 4, a.Count(1))
	assert.Equal(t, 3, a.Count(3))
	assert.Equal(t, 1, a.Count(6))
	assert.Equal(t, 2, a.Count(9))
	assert.Equal(t, 0, a.Count(7))
}

func TestGetBoundingBox(t *testing.T) {
	a := NewArray3d(7, 8, 9)
	box := a.GetBoundingBox()
	assert.True(t, box[0] == a.DimX && box[1] == a.DimY && box[2] == a.DimZ, "Bounding box dimensions are wrong")
}

func TestArray3dApply(t *testing.T) {
	f := func(x, y, z int, currentValue int8) int8 {
		return int8(x + y + z)
	}
	a := NewArray3d(2, 3, 4).Apply(f)

	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < a.DimZ; z++ {
				assert.Equal(t, int8(x+y+z), a.Get(x, y, z))
			}
		}
	}
}

func TestArray3dAllTrue(t *testing.T) {
	a := NewArray3dFromData([][][]int8{{{0, 1}, {1, 2}}, {{1, 2}, {2, 3}}, {{2, 3}, {3, 4}}})

	assert.True(t, a.AllTrue(func(x, y, z int, v int8) bool {
		return int8(x+y+z) == v
	}))

	assert.False(t, a.AllTrue(func(x, y, z int, v int8) bool {
		return int8(x+y) == v
	}))
}

func TestArray3dRotateZ(t *testing.T) {
	orig := NewArray3dFromData([][][]int8{{{0}, {3}}, {{1}, {4}}, {{2}, {5}}})
	exp := NewArray3dFromData([][][]int8{{{3}, {4}, {5}}, {{0}, {1}, {2}}})
	r := orig.RotateZ()

	// rotate once
	assert.True(t, r.IsEqual(exp))

	// rotate 2x 2x should be the identity
	assert.True(t, orig.IsEqual(orig.RotateZ2().RotateZ2()))

	// rotate x and 3x shoudl be the identity
	assert.True(t, orig.IsEqual(orig.RotateZ3().RotateZ()))

	// rotate 4x should be the identity
	assert.True(t, orig.IsEqual(orig.RotateZ().RotateZ().RotateZ().RotateZ()))
}

func TestArray3dRotateY(t *testing.T) {
	orig := NewArray3dFromData([][][]int8{{{3, 0}}, {{4, 1}}, {{5, 2}}})
	exp := NewArray3dFromData([][][]int8{{{5, 4, 3}}, {{2, 1, 0}}})
	r := orig.RotateY()

	// rotate once
	assert.True(t, r.IsEqual(exp))

	// rotate 2x 2x should be the identity
	assert.True(t, orig.IsEqual(orig.RotateY2().RotateY2()))

	// rotate x and 3x shoudl be the identity
	assert.True(t, orig.IsEqual(orig.RotateY3().RotateY()))

	// rotate 4x should be the identity
	assert.True(t, orig.IsEqual(orig.RotateY().RotateY().RotateY().RotateY()))
}

func TestArray3dRotateX(t *testing.T) {
	orig := NewArray3dFromData([][][]int8{{{0, 1, 2}, {3, 4, 5}}})
	exp := NewArray3dFromData([][][]int8{{{3, 0}, {4, 1}, {5, 2}}})
	r := orig.RotateX()

	// rotate once
	assert.True(t, r.IsEqual(exp))

	// rotate 2x 2x should be the identity
	assert.True(t, orig.IsEqual(orig.RotateX2().RotateX2()))

	// rotate x and 3x shoudl be the identity
	assert.True(t, orig.IsEqual(orig.RotateX3().RotateX()))

	// rotate 4x should be the identity
	assert.True(t, orig.IsEqual(orig.RotateX().RotateX().RotateX().RotateX()))
}

func TestArray3dGetCenterOfGravity(t *testing.T) {
	a := NewArray3d(5, 5, 5)
	a.Set(0, 2, 3, 1)
	a.Set(1, 2, 4, 1)
	a.Set(4, 0, 1, 1)

	cog := a.GetCenterOfGravity()
	assert.True(t, math.Abs(13/6.0-cog[0]) < 1e-10)
	assert.True(t, math.Abs(11/6.0-cog[1]) < 1e-10)
	assert.True(t, math.Abs(19/6.0-cog[2]) < 1e-10)
}

func TestFindArray3d(t *testing.T) {
	lst := []*Array3d{
		NewArray3dFromData([][][]int8{{{0, -1, -1}, {0, 0, 0}}, {{0, -1, 0}, {0, -1, -1}}}),
		NewArray3dFromData([][][]int8{{{0, -1}, {0, 0}}, {{0, 0}, {-1, -1}}}),
		NewArray3dFromData([][][]int8{{{-1}, {-1}}, {{0}, {0}}, {{-1}, {0}}}),
		NewArray3dFromData([][][]int8{{{0, -1, -1}}, {{0, -1, 0}}}),
	}

	inList := NewArray3dFromData([][][]int8{{{0, -1}, {0, 0}}, {{0, 0}, {-1, -1}}})
	notInList := NewArray3dFromData([][][]int8{{{-1}, {0}}, {{0}, {-1}}, {{0}, {-1}}})

	found, idx := FindArray3d(lst, inList)
	assert.True(t, found)
	assert.Equal(t, 1, idx)

	found, idx = FindArray3d(lst, notInList)
	assert.False(t, found)
	assert.Equal(t, -1, idx)
}

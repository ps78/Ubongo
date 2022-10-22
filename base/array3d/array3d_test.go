package array3d_test

import (
	"math"
	"testing"

	. "ubongo/base/array3d"
	"ubongo/base/vector"
	"ubongo/base/vectorf"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	a := NewFromData([][][]int8{{{0}, {1}, {0}}, {{1}, {0}, {1}}})
	exp := "<2-3-1>[[[0] [1] [0]] [[1] [0] [1]]]"
	act := a.String()
	assert.Equal(t, exp, act)
}

func TestStringNil(t *testing.T) {
	var a *A = nil
	exp := "(nil)"
	act := a.String()
	assert.Equal(t, exp, act)
}

func TestNew(t *testing.T) {
	dimX := 2
	dimY := 3
	dimZ := 1

	a := New(dimX, dimY, dimZ)

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

func TestNewError(t *testing.T) {
	assert.Panics(t, func() { New(1, 1, 0) })
	assert.Panics(t, func() { New(1, 0, 1) })
	assert.Panics(t, func() { New(0, 1, 1) })
}

func TestNewFromData(t *testing.T) {
	data := [][][]int8{{{0, 1}, {1, 2}, {2, 3}}, {{1, 2}, {2, 3}, {3, 4}}, {{2, 3}, {3, 4}, {4, 5}}}
	a := NewFromData(data)

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

func TestNewFromDataNil(t *testing.T) {
	assert.Nil(t, NewFromData(nil))
}

func TestGetSet(t *testing.T) {
	a := New(7, 9, 3)
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < a.DimZ; z++ {
				a.Set(x, y, z, int8(x+y-z))
				assert.Equal(t, int8(x+y-z), a.Get(x, y, z), "Get/Set at position (%d,%d,%d) did not work", x, y, z)
			}
		}
	}
}

func TestIsEqual(t *testing.T) {
	a := NewFromData([][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 5}}})
	b := NewFromData([][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 5}}}) // == a
	c := NewFromData([][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 0}}}) // != a
	d := NewFromData([][][]int8{{{2}, {3}}})
	var e *A = nil

	assert.True(t, a.Equals(b), "Array a and b are equal but Equal3DArray reports they are not")
	assert.False(t, a.Equals(c), "Array a and c are not equal but Equal3DArray reports they are")
	assert.False(t, a.Equals(d), "Array a and d have different dimensions but Equal3DArray reports they are equal")
	assert.True(t, e.Equals(nil))
	assert.False(t, e.Equals(a))
	assert.False(t, a.Equals(e))
}

func TestClone(t *testing.T) {
	orig := NewFromData([][][]int8{{{0, 1}, {1, 2}}, {{1, 2}, {2, 3}}, {{2, 3}, {3, 4}}})
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

func TestCloneNil(t *testing.T) {
	var a *A = nil
	assert.Nil(t, a.Clone())
}

func TestCount(t *testing.T) {
	a := NewFromData([][][]int8{{{1, 3}, {6, 1}}, {{3, 3}, {0, 9}}, {{1, 1}, {9, 0}}})
	assert.Equal(t, 2, a.Count(0))
	assert.Equal(t, 4, a.Count(1))
	assert.Equal(t, 3, a.Count(3))
	assert.Equal(t, 1, a.Count(6))
	assert.Equal(t, 2, a.Count(9))
	assert.Equal(t, 0, a.Count(7))
}

func TestCountNil(t *testing.T) {
	var a *A = nil
	assert.Equal(t, 0, a.Count(0))
}

func TestGetBoundingBox(t *testing.T) {
	a := New(7, 8, 9)
	box := a.GetBoundingBox()
	assert.True(t, box[0] == a.DimX && box[1] == a.DimY && box[2] == a.DimZ, "Bounding box dimensions are wrong")
}

func TestGetBoundingBoxNil(t *testing.T) {
	var a *A = nil
	box := a.GetBoundingBox()
	assert.Equal(t, vector.Zero, box)
}
func TestApply(t *testing.T) {
	f := func(x, y, z int, currentValue int8) int8 {
		return int8(x + y + z)
	}
	a := New(2, 3, 4).Apply(f)

	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			for z := 0; z < a.DimZ; z++ {
				assert.Equal(t, int8(x+y+z), a.Get(x, y, z))
			}
		}
	}
}

func TestApplyOnNil(t *testing.T) {
	var a *A = nil
	f := func(x, y, z int, currentValue int8) int8 {
		return int8(x + y + z)
	}
	assert.Nil(t, a.Apply(f))
}

func TestApplyNilFunc(t *testing.T) {
	a := New(2, 3, 4)
	b := a.Apply(nil)
	assert.True(t, a.Equals(b))
}

func TestAllTrue(t *testing.T) {
	a := NewFromData([][][]int8{{{0, 1}, {1, 2}}, {{1, 2}, {2, 3}}, {{2, 3}, {3, 4}}})

	assert.True(t, a.AllTrue(func(x, y, z int, v int8) bool {
		return int8(x+y+z) == v
	}))

	assert.False(t, a.AllTrue(func(x, y, z int, v int8) bool {
		return int8(x+y) == v
	}))
}

func TestAllTrueNil(t *testing.T) {
	var a *A = nil

	assert.False(t, a.AllTrue(func(x, y, z int, v int8) bool {
		return true
	}))

	assert.False(t, a.AllTrue(nil))
}

func TestRotateZ(t *testing.T) {
	orig := NewFromData([][][]int8{{{0}, {3}}, {{1}, {4}}, {{2}, {5}}})
	exp := NewFromData([][][]int8{{{3}, {4}, {5}}, {{0}, {1}, {2}}})
	r := orig.RotateZ()

	// rotate once
	assert.True(t, r.Equals(exp))

	// rotate 2x 2x should be the identity
	assert.True(t, orig.Equals(orig.RotateZ2().RotateZ2()))

	// rotate x and 3x shoudl be the identity
	assert.True(t, orig.Equals(orig.RotateZ3().RotateZ()))

	// rotate 4x should be the identity
	assert.True(t, orig.Equals(orig.RotateZ().RotateZ().RotateZ().RotateZ()))
}

func TestRotateY(t *testing.T) {
	orig := NewFromData([][][]int8{{{3, 0}}, {{4, 1}}, {{5, 2}}})
	exp := NewFromData([][][]int8{{{5, 4, 3}}, {{2, 1, 0}}})
	r := orig.RotateY()

	// rotate once
	assert.True(t, r.Equals(exp))

	// rotate 2x 2x should be the identity
	assert.True(t, orig.Equals(orig.RotateY2().RotateY2()))

	// rotate x and 3x shoudl be the identity
	assert.True(t, orig.Equals(orig.RotateY3().RotateY()))

	// rotate 4x should be the identity
	assert.True(t, orig.Equals(orig.RotateY().RotateY().RotateY().RotateY()))
}

func TestRotateX(t *testing.T) {
	orig := NewFromData([][][]int8{{{0, 1, 2}, {3, 4, 5}}})
	exp := NewFromData([][][]int8{{{3, 0}, {4, 1}, {5, 2}}})
	r := orig.RotateX()

	// rotate once
	assert.True(t, r.Equals(exp))

	// rotate 2x 2x should be the identity
	assert.True(t, orig.Equals(orig.RotateX2().RotateX2()))

	// rotate x and 3x shoudl be the identity
	assert.True(t, orig.Equals(orig.RotateX3().RotateX()))

	// rotate 4x should be the identity
	assert.True(t, orig.Equals(orig.RotateX().RotateX().RotateX().RotateX()))
}

func TestRotateNil(t *testing.T) {
	var a *A = nil

	assert.Nil(t, a.RotateX())
	assert.Nil(t, a.RotateX2())
	assert.Nil(t, a.RotateX3())
	assert.Nil(t, a.RotateY())
	assert.Nil(t, a.RotateY2())
	assert.Nil(t, a.RotateY3())
	assert.Nil(t, a.RotateZ())
	assert.Nil(t, a.RotateZ2())
	assert.Nil(t, a.RotateZ3())
}

func TestGetCenterOfGravity(t *testing.T) {
	a := New(5, 5, 5)
	a.Set(0, 2, 3, 1)
	a.Set(1, 2, 4, 1)
	a.Set(4, 0, 1, 1)

	cog := a.GetCenterOfGravity()
	assert.True(t, math.Abs(13/6.0-cog[0]) < 1e-10)
	assert.True(t, math.Abs(11/6.0-cog[1]) < 1e-10)
	assert.True(t, math.Abs(19/6.0-cog[2]) < 1e-10)
}

func TestGetCenterOfGravityNil(t *testing.T) {
	var a *A = nil
	assert.Equal(t, vectorf.Zero, a.GetCenterOfGravity())
}

func TestFind(t *testing.T) {
	lst := []*A{
		NewFromData([][][]int8{{{0, -1, -1}, {0, 0, 0}}, {{0, -1, 0}, {0, -1, -1}}}),
		NewFromData([][][]int8{{{0, -1}, {0, 0}}, {{0, 0}, {-1, -1}}}),
		NewFromData([][][]int8{{{-1}, {-1}}, {{0}, {0}}, {{-1}, {0}}}),
		NewFromData([][][]int8{{{0, -1, -1}}, {{0, -1, 0}}}),
	}

	inList := NewFromData([][][]int8{{{0, -1}, {0, 0}}, {{0, 0}, {-1, -1}}})
	notInList := NewFromData([][][]int8{{{-1}, {0}}, {{0}, {-1}}, {{0}, {-1}}})

	found, idx := Find(lst, inList)
	assert.True(t, found)
	assert.Equal(t, 1, idx)

	found, idx = Find(lst, notInList)
	assert.False(t, found)
	assert.Equal(t, -1, idx)

	found, idx = Find(lst, nil)
	assert.False(t, found)
	assert.Equal(t, -1, idx)

	found, idx = Find(nil, inList)
	assert.False(t, found)
	assert.Equal(t, -1, idx)
}

func TestCreateRotations(t *testing.T) {
	// a completely symmetric, empty 3x3x3-cube should only return one rotation
	a := New(3, 3, 3)
	assert.Equal(t, 1, len(a.CreateRotations()))

	// this case represents the complex yellow block from the game with no symmetries
	b := NewFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {1, 0}, {0, 0}}})
	assert.Equal(t, 24, len(b.CreateRotations()))
}

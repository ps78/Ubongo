package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArray2d(t *testing.T) {
	dimX := 2
	dimY := 3

	a := NewArray2d(dimX, dimY)

	assert.Equal(t, dimX, a.DimX, "Wrong x-dimension")
	assert.Equal(t, dimY, a.DimY, "Wrong y-dimension")

	for x := 0; x < dimX; x++ {
		for y := 0; y < dimY; y++ {
			assert.Equal(t, int8(0), a.Get(x, y), "NewArray2d returned an array that is not zeroed at all positions")
		}
	}
}

func TestNewArray2dFromData(t *testing.T) {
	data := [][]int8{{0, 1}, {1, 2}, {2, 3}}
	a := NewArray2dFromData(data)

	assert.Equal(t, len(data), a.DimX, "Wrong x-dimension")
	assert.Equal(t, len(data[0]), a.DimY, "Wrong y-dimension")

	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			data[x][y] = 0 // this should not affect a!
			assert.Equal(t, int8(x+y), a.Get(x, y), "NewArray2dFromData returned an array that did not propery copy data")
		}
	}
}

func TestArray2dString(t *testing.T) {
	a := NewArray2dFromData([][]int8{{0, 1, 0}, {1, 0, 1}})
	exp := "<2-3>[[0 1 0] [1 0 1]]"
	act := a.String()
	assert.Equal(t, exp, act)
}

func TestArray2dGetSet(t *testing.T) {
	a := NewArray2d(7, 9)
	for x := 0; x < a.DimX; x++ {
		for y := 0; y < a.DimY; y++ {
			a.Set(x, y, int8(x+y))
			assert.Equal(t, int8(x+y), a.Get(x, y), "Get/Set at position (%d,%d) did not work", x, y)
		}
	}
}

func TestArray2dExtrude(t *testing.T) {
	a2d := NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}})
	height := 2
	a := a2d.Extrude(height)

	for x := 0; x < a2d.DimX; x++ {
		for y := 0; y < a2d.DimY; y++ {
			for z := 0; z < height; z++ {
				assert.Equal(t, a2d.Get(x, y), a.Get(x, y, z), "Extrude returned an invalid result a position %s", Vector{x, y, z})
			}
		}
	}
}

func TestArray2dIsEqual(t *testing.T) {
	a := NewArray2dFromData([][]int8{{2, 3, 5, 6}, {7, 8, 6, 2}, {1, 0, -1, 5}})
	b := NewArray2dFromData([][]int8{{2, 3, 5, 6}, {7, 8, 6, 2}, {1, 0, -1, 5}}) // == a
	c := NewArray2dFromData([][]int8{{2, 3, 5, 6}, {7, 8, 6, 2}, {1, 0, -1, 0}}) // != a
	d := NewArray2dFromData([][]int8{{2}, {3}})
	var e *Array2d = nil

	assert.True(t, a.IsEqual(b), "Array a and b are equal but Equal2DArray reports they are not")
	assert.False(t, a.IsEqual(c), "Array a and c are not equal but Equal2DArray reports they are")
	assert.False(t, a.IsEqual(d), "Array a and d have different dimensions but Equal2DArray reports they are equal")
	assert.True(t, e.IsEqual(nil))
	assert.False(t, e.IsEqual(a))
	assert.False(t, a.IsEqual(e))
}

func TestArray2dClone(t *testing.T) {
	orig := NewArray2dFromData([][]int8{{0, 1}, {1, 2}, {2, 3}})
	copy := orig.Clone()
	orig.Set(0, 0, 42) // this should not affect the copy

	assert.True(t, copy.DimX == orig.DimX && copy.DimY == orig.DimY, "Dimensions do not match")

	for x := 0; x < orig.DimX; x++ {
		for y := 0; y < orig.DimY; y++ {
			assert.Equal(t, int8(x+y), copy.Get(x, y), "Element [%d][%d] has wrong value", x, y)
		}
	}
}

func TestArray2dCount(t *testing.T) {
	a := NewArray2dFromData([][]int8{{1, 3, 6, 1}, {3, 3, 0, 9}, {1, 1, 9, 0}})
	assert.Equal(t, 2, a.Count(0))
	assert.Equal(t, 4, a.Count(1))
	assert.Equal(t, 3, a.Count(3))
	assert.Equal(t, 1, a.Count(6))
	assert.Equal(t, 2, a.Count(9))
	assert.Equal(t, 0, a.Count(7))
}

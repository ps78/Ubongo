package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorString(t *testing.T) {
	v := Vector{1, 2, 3}
	assert.Equal(t, "(1,2,3)", fmt.Sprint(v))
}

// Tests the function GetShiftVectors for arguments that result in an non-empty list
func TestGetShiftVectors(t *testing.T) {
	outer := Vector{4, 3, 2}
	inner := Vector{3, 2, 1}
	shifts := outer.GetShiftVectors(inner)
	if len(shifts) != 8 {
		t.Errorf("GetShiftVectors return %d results, expected %d", len(shifts), 8)
	}

	expected := []Vector{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1}}

	for _, v := range expected {
		assert.Contains(t, shifts, v, "Vector %s is missing in output of GetShiftVector()", v)
	}
}

// Tests the function GetShiftVectors for arguments that result in an empty list
func TestGetShiftVectorsEmpty(t *testing.T) {
	outer := Vector{4, 3, 2}
	inner := Vector{5, 2, 1}
	shifts := outer.GetShiftVectors(inner)
	assert.Equal(t, 0, len(shifts), "GetShiftVectors did not return an empty slice")
}

func TestMake2DArray(t *testing.T) {
	xdim := 2
	ydim := 3

	a := Make2DArray(xdim, ydim)

	assert.Equal(t, xdim, len(a), "Wrong x-dimension %d instead of %d", len(a), xdim)
	assert.Equal(t, ydim, len(a[0]), "Wrong y-dimension %d instead of %d", len(a[0]), ydim)

	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {
			assert.Zero(t, a[x][y], "Make2DArray returned an array that is not zeroed at all positions")
		}
	}
}

func TestMake3DArray(t *testing.T) {
	xdim := 2
	ydim := 3
	zdim := 1

	a := Make3DArray(xdim, ydim, zdim)

	assert.Equal(t, xdim, len(a), "Wrong x-dimension %d instead of %d", len(a), xdim)
	assert.Equal(t, ydim, len(a[0]), "Wrong y-dimension %d instead of %d", len(a[0]), ydim)
	assert.Equal(t, zdim, len(a[0][0]), "Wrong z-dimension %d instead of %d", len(a[0][0]), zdim)

	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {
			for z := 0; z < zdim; z++ {
				assert.Zero(t, a[x][y][z], "Make3DArray returned an array that is not zeroed at all positions")
			}
		}
	}
}

func TestExtrude2DArray(t *testing.T) {
	shape := Array2d{{-1, 0, -1}, {-1, 0, 0}}
	xdim := len(shape)
	ydim := len(shape[0])
	height := 2
	a := Extrude2DArray(shape, height)

	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {
			for z := 0; z < height; z++ {
				assert.Equal(t, shape[x][y], a[x][y][z], "Extrude2DArray returned an invalid result a position %s", Vector{x, y, z})
			}
		}
	}
}

func TestEqual2DArray(t *testing.T) {
	a := [][]int8{{2, 3, 5, 6}, {7, 8, 6, 2}, {1, 0, -1, 5}}
	b := [][]int8{{2, 3, 5, 6}, {7, 8, 6, 2}, {1, 0, -1, 5}} // == a
	c := [][]int8{{2, 3, 5, 6}, {7, 8, 6, 2}, {1, 0, -1, 0}} // != a
	d := [][]int8{{2}, {3}}

	assert.True(t, Equal2DArray(a, b), "Array a and b are equal but Equal2DArray reports they are not")
	assert.False(t, Equal2DArray(a, c), "Array a and c are not equal but Equal2DArray reports they are")
	assert.False(t, Equal2DArray(a, d), "Array a and d have different dimensions but Equal2DArray reports they are equal")
}

func TestEqual3DArray(t *testing.T) {
	a := [][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 5}}}
	b := [][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 5}}} // == a
	c := [][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 0}}} // != a
	d := [][][]int8{{{2}, {3}}}

	assert.True(t, Equal3DArray(a, b), "Array a and b are equal but Equal3DArray reports they are not")
	assert.False(t, Equal3DArray(a, c), "Array a and c are not equal but Equal3DArray reports they are")
	assert.False(t, Equal3DArray(a, d), "Array a and d have different dimensions but Equal3DArray reports they are equal")
}

func TestCopy2DArray(t *testing.T) {
	orig := [][]int8{{0, 1}, {1, 2}, {2, 3}}
	copy := Copy2DArray(orig)
	orig[0][0] = 42 // this should not affect the copy

	assert.True(t, len(copy) == len(orig) && len(copy[0]) == len(orig[0]),
		"Dimensions of copy and original do not match: %d,%d instead of 3,2", len(copy), len(copy[0]))

	for i := range orig {
		for j := range orig[i] {
			assert.Equal(t, int8(i+j), copy[i][j], "Element [%d][%d] has wrong value", i, j)
		}
	}
}

func TestCopy3DArray(t *testing.T) {
	orig := [][][]int8{{{0, 1}, {1, 2}}, {{1, 2}, {2, 3}}, {{2, 3}, {3, 4}}}
	copy := Copy3DArray(orig)
	orig[0][0][0] = 42 // this should not affect the copy

	assert.True(t, len(copy) == len(orig) && len(copy[0]) == len(orig[0]) && len(copy[0][0]) == len(orig[0][0]),
		"Dimensions of copy and original do not match: %d,%d,%d instead of 3,2,1",
		len(copy), len(copy[0]), len(copy[0][0]))

	for i := range orig {
		for j := range orig[i] {
			for k := range orig[i][j] {
				assert.Equal(t, int8(i+j+k), copy[i][j][k], "Element [%d][%d][%d] does not match", i, j, k)
			}
		}
	}
}

func TestCountValues2D(t *testing.T) {
	a := Array2d{{1, 3, 6, 1}, {3, 3, 0, 9}, {1, 1, 9, 0}}
	assert.Equal(t, 2, CountValues2D(a, 0))
	assert.Equal(t, 4, CountValues2D(a, 1))
	assert.Equal(t, 3, CountValues2D(a, 3))
	assert.Equal(t, 1, CountValues2D(a, 6))
	assert.Equal(t, 2, CountValues2D(a, 9))
	assert.Equal(t, 0, CountValues2D(a, 7))
}

func TestCountValues3D(t *testing.T) {
	a := Array3d{{{1, 3}, {6, 1}}, {{3, 3}, {0, 9}}, {{1, 1}, {9, 0}}}
	assert.Equal(t, 2, CountValues3D(a, 0))
	assert.Equal(t, 4, CountValues3D(a, 1))
	assert.Equal(t, 3, CountValues3D(a, 3))
	assert.Equal(t, 1, CountValues3D(a, 6))
	assert.Equal(t, 2, CountValues3D(a, 9))
	assert.Equal(t, 0, CountValues3D(a, 7))
}

func TestGetBoundingBoxFromArray(t *testing.T) {
	a := Make3DArray(7, 8, 9)
	box := GetBoundingBoxFromArray(a)
	assert.True(t, box[0] == len(a) && box[1] == len(a[0]) && box[2] == len(a[0][0]),
		"Bounding box dimensions are wrong")
}

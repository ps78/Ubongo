package main

import (
	"testing"
)

// helper function for testing
func arrayContainsVector(vectorArray []Vector, v Vector) bool {
	if vectorArray == nil {
		return false
	}
	for _, vec := range vectorArray {
		if vec == v {
			return true
		}
	}
	return false
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
		if !arrayContainsVector(shifts, v) {
			t.Errorf("Vector %s is missing in output of GetShiftVector()", v)
		}
	}
}

// Tests the function GetShiftVectors for arguments that result in an empty list
func TestGetShiftVectorsEmpty(t *testing.T) {
	outer := Vector{4, 3, 2}
	inner := Vector{5, 2, 1}
	shifts := outer.GetShiftVectors(inner)
	if len(shifts) != 0 {
		t.Errorf("GetShiftVectors did not return an empty slice")
	}
}

func TestMake2DArray(t *testing.T) {
	xdim := 2
	ydim := 3

	a := Make2DArray(xdim, ydim)

	if len(a) != xdim || len(a[0]) != ydim {
		t.Errorf("Make2DArray retuned an array with the wrong dimensions (%dx%d instead of %dx%d)",
			len(a), len(a[0]), xdim, ydim)
	}
	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {
			if a[x][y] != 0 {
				t.Errorf("Make2DArray returned an array that is not zeroed at all positions")
			}
		}
	}
}

func TestCopy2DArray(t *testing.T) {
	orig := [][]int8{{0, 1}, {1, 2}, {2, 3}}
	copy := Copy2DArray(orig)
	orig[0][0] = 42 // this should not affect the copy

	if len(copy) != len(orig) || len(copy[0]) != len(orig[0]) {
		t.Errorf("Dimensions of copy and original do not match: %d,%d instead of 3,2", len(copy), len(copy[0]))
	}
	for i := range orig {
		for j := range orig[i] {
			if copy[i][j] != int8(i+j) {
				t.Errorf("Element [%d][%d] does not match (is %d instead of %d", i, j, copy[i][j], i+j)
			}
		}
	}
}

func TestCopy3DArray(t *testing.T) {
	orig := [][][]int8{{{0, 1}, {1, 2}}, {{1, 2}, {2, 3}}, {{2, 3}, {3, 4}}}
	copy := Copy3DArray(orig)
	orig[0][0][0] = 42 // this should not affect the copy

	if len(copy) != len(orig) || len(copy[0]) != len(orig[0]) || len(copy[0][0]) != len(orig[0][0]) {
		t.Errorf("Dimensions of copy and original do not match: %d,%d,%d instead of 3,2,1",
			len(copy), len(copy[0]), len(copy[0][0]))
	}
	for i := range orig {
		for j := range orig[i] {
			for k := range orig[i][j] {
				if copy[i][j][k] != int8(i+j+k) {
					t.Errorf("Element [%d][%d][%d] does not match (is %d instead of %d",
						i, j, k, copy[i][j][k], i+j+k)
				}
			}
		}
	}
}

func TestArray3dEqual(t *testing.T) {
	a := [][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 5}}}
	b := [][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 5}}} // == a
	c := [][][]int8{{{2, 3}, {5, 6}, {7, 8}}, {{6, 2}, {1, 0}, {-1, 0}}} // != a
	d := [][][]int8{{{2}, {3}}}

	if !Equal3DArray(a, b) {
		t.Errorf("Array a and b are equal but Equal3DArray reports they are not")
	}
	if Equal3DArray(a, c) {
		t.Errorf("Array a and c are not equal but Equal3DArray reports they are")
	}
	if Equal3DArray(a, d) {
		t.Errorf("Array a and d have different dimensions but Equal3DArray reports they are equal")
	}
}

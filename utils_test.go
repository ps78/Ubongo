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

func TestFindArray3d(t *testing.T) {
	block := GetBlockFactory().Get(1)

	for i, a := range block.Shapes {
		ok, idx := FindArray3d(block.Shapes, a)
		assert.True(t, ok)
		assert.Equal(t, i, idx)
	}

	ok, idx := FindArray3d(block.Shapes, NewArray3d(2, 3, 4))
	assert.Equal(t, -1, idx)
	assert.False(t, ok)
}

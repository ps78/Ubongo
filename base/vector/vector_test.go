package vector_test

import (
	"fmt"
	"testing"
	. "ubongo/base/vector"
	"ubongo/base/vectorf"

	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	assert.Equal(t, V{0, 0, 0}, Zero)
}

func TestAsVectorf(t *testing.T) {
	v := V{1, 2, 3}
	vf := v.AsVectorf()
	assert.Equal(t, vectorf.V{1, 2, 3}, vf)
}

func TestString(t *testing.T) {
	v := V{1, 2, 3}
	assert.Equal(t, "(1,2,3)", fmt.Sprint(v))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 77, V{-10, 3, 77}.Max())
	assert.Equal(t, 55, V{5, 55, -1}.Max())
	assert.Equal(t, 11, V{11, -33, 0}.Max())
}

func TestAdd(t *testing.T) {
	a := V{-1, 5, 42}
	b := V{3, 0, 5}
	actual := a.Add(b)
	expected := V{2, 5, 47}
	assert.Equal(t, expected, actual)
}

func TestSub(t *testing.T) {
	a := V{-1, 5, 42}
	b := V{3, 0, 5}
	actual := a.Sub(b)
	expected := V{-4, 5, 37}
	assert.Equal(t, expected, actual)
}

func TestMult(t *testing.T) {
	a := V{-1, 5, 42}
	b := 3
	actual := a.Mult(b)
	expected := V{-3, 15, 126}
	assert.Equal(t, expected, actual)
}

func TestDiv(t *testing.T) {
	a := V{-1, 5, 42}
	b := 2.0
	actual := a.Div(b)
	expected := vectorf.V{-0.5, 2.5, 21.0}
	assert.Equal(t, expected, actual)
}

func TestFlip(t *testing.T) {
	a := V{-1, 5, 42}
	actual := a.Flip()
	expected := V{1, -5, -42}
	assert.Equal(t, expected, actual)
}

// Tests the function GetShiftVectors for arguments that result in an non-empty list
func TestGetShiftVectors(t *testing.T) {
	outer := V{4, 3, 2}
	inner := V{3, 2, 1}
	shifts := outer.GetShiftVectors(inner)
	if len(shifts) != 8 {
		t.Errorf("GetShiftVectors return %d results, expected %d", len(shifts), 8)
	}

	expected := []V{
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
	outer := V{4, 3, 2}
	inner := V{5, 2, 1}
	shifts := outer.GetShiftVectors(inner)
	assert.Equal(t, 0, len(shifts), "GetShiftVectors did not return an empty slice")
}

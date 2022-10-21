package vector_test

import (
	"fmt"
	"testing"
	. "ubongo/base/vector"
	"ubongo/base/vectorf"

	"github.com/stretchr/testify/assert"
)

func TestVectorAsVectorf(t *testing.T) {
	v := V{1, 2, 3}
	vf := v.AsVectorf()
	assert.Equal(t, vectorf.V{1, 2, 3}, vf)
}

func TestVectorString(t *testing.T) {
	v := V{1, 2, 3}
	assert.Equal(t, "(1,2,3)", fmt.Sprint(v))
}

func TestVectorMax(t *testing.T) {
	a := V{-10, 3, 77}
	m := a.Max()
	assert.Equal(t, 77, m)
}

func TestVectorAdd(t *testing.T) {
	a := V{-1, 5, 42}
	b := V{3, 0, 5}
	actual := a.Add(b)
	expected := V{2, 5, 47}
	assert.Equal(t, expected, actual)
}

func TestVectorSub(t *testing.T) {
	a := V{-1, 5, 42}
	b := V{3, 0, 5}
	actual := a.Sub(b)
	expected := V{-4, 5, 37}
	assert.Equal(t, expected, actual)
}

func TestVectorMult(t *testing.T) {
	a := V{-1, 5, 42}
	b := 3
	actual := a.Mult(b)
	expected := V{-3, 15, 126}
	assert.Equal(t, expected, actual)
}

func TestVectorDiv(t *testing.T) {
	a := V{-1, 5, 42}
	b := 2.0
	actual := a.Div(b)
	expected := vectorf.V{-0.5, 2.5, 21.0}
	assert.Equal(t, expected, actual)
}

func TestVectorFlip(t *testing.T) {
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

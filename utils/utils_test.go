package utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorAsVectorf(t *testing.T) {
	v := Vector{1, 2, 3}
	vf := v.AsVectorf()
	assert.Equal(t, Vectorf{1, 2, 3}, vf)
}

func TestVectorString(t *testing.T) {
	v := Vector{1, 2, 3}
	assert.Equal(t, "(1,2,3)", fmt.Sprint(v))
}

func TestVectorMax(t *testing.T) {
	a := Vector{-10, 3, 77}
	m := a.Max()
	assert.Equal(t, 77, m)
}

func TestVectorAdd(t *testing.T) {
	a := Vector{-1, 5, 42}
	b := Vector{3, 0, 5}
	actual := a.Add(b)
	expected := Vector{2, 5, 47}
	assert.Equal(t, expected, actual)
}

func TestVectorSub(t *testing.T) {
	a := Vector{-1, 5, 42}
	b := Vector{3, 0, 5}
	actual := a.Sub(b)
	expected := Vector{-4, 5, 37}
	assert.Equal(t, expected, actual)
}

func TestVectorMult(t *testing.T) {
	a := Vector{-1, 5, 42}
	b := 3
	actual := a.Mult(b)
	expected := Vector{-3, 15, 126}
	assert.Equal(t, expected, actual)
}

func TestVectorDiv(t *testing.T) {
	a := Vector{-1, 5, 42}
	b := 2.0
	actual := a.Div(b)
	expected := Vectorf{-0.5, 2.5, 21.0}
	assert.Equal(t, expected, actual)
}

func TestVectorFlip(t *testing.T) {
	a := Vector{-1, 5, 42}
	actual := a.Flip()
	expected := Vector{1, -5, -42}
	assert.Equal(t, expected, actual)
}

func TestVectorfString(t *testing.T) {
	v := Vectorf{1.5, 2.3, 3.8}
	expected := "(1.5,2.3,3.8)"
	actual := strings.ReplaceAll(fmt.Sprint(v), "0", "")
	assert.Equal(t, expected, actual)
}

func TestVectorfMax(t *testing.T) {
	a := Vectorf{1.5, 0, -5}
	m := a.Max()
	assert.Equal(t, 1.5, m)
}

func TestVectorfAdd(t *testing.T) {
	a := Vectorf{1.5, 0, -5}
	b := Vectorf{-1, 2, 3.5}
	expected := Vectorf{0.5, 2, -1.5}
	assert.Equal(t, expected, a.Add(b))
}

func TestVectorfSub(t *testing.T) {
	a := Vectorf{1.5, 0, -5}
	b := Vectorf{-1, 2, 3.5}
	expected := Vectorf{2.5, -2, -8.5}
	assert.Equal(t, expected, a.Sub(b))
}

func TestVectorfDiv(t *testing.T) {
	a := Vectorf{1.5, 0, -5}
	b := 2.0
	expected := Vectorf{0.75, 0, -2.5}
	assert.Equal(t, expected, a.Div(b))
}
func TestVectorfMult(t *testing.T) {
	a := Vectorf{1.5, 0, -5}
	b := 2.0
	expected := Vectorf{3.0, 0, -10.0}
	assert.Equal(t, expected, a.Mult(b))
}

func TestVectorfFlip(t *testing.T) {
	a := Vectorf{1.5, 0, -5}
	expected := Vectorf{-1.5, 0, 5.0}
	assert.Equal(t, expected, a.Flip())
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

func TestCreatePartitions(t *testing.T) {
	n := 21
	parts := []int{5, 4, 3}
	partLen := 5
	maxCounts := map[int]int{3: 1, 4: 5, 5: 10}
	partitions := CreateParitions(n, parts, maxCounts, partLen)

	assert.Equal(t, 2, len(partitions))
	for _, part := range partitions {
		sum := 0
		count := 0
		for k, v := range part {
			sum += k * v
			count += v
		}
		assert.Equal(t, n, sum)
		assert.Equal(t, partLen, count)
	}
}

func TestCreatePartitionsNoResult(t *testing.T) {
	n := 18
	parts := []int{5, 4, 3}
	partLen := 5
	maxCounts := map[int]int{3: 1, 4: 5, 5: 10}
	partitions := CreateParitions(n, parts, maxCounts, partLen)

	assert.Equal(t, 0, len(partitions))
}

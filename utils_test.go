package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorString(t *testing.T) {
	v := Vector{1, 2, 3}
	assert.Equal(t, "(1,2,3)", fmt.Sprint(v))
}

func TestVectorfString(t *testing.T) {
	v := Vectorf{1.5, 2.3, 3.8}
	expected := "(1.5,2.3,3.8)"
	actual := strings.ReplaceAll(fmt.Sprint(v), "0", "")
	assert.Equal(t, expected, actual)
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
	block := GetBlockFactory().ByNumber(1)

	for i, a := range block.Shapes {
		ok, idx := FindArray3d(block.Shapes, a)
		assert.True(t, ok)
		assert.Equal(t, i, idx)
	}

	ok, idx := FindArray3d(block.Shapes, NewArray3d(2, 3, 4))
	assert.Equal(t, -1, idx)
	assert.False(t, ok)
}

func TestCreatePartitions(t *testing.T) {
	n := 21
	parts := []int{3, 4, 5}
	partLen := 5
	maxCounts := map[int]int{3: 1, 4: 10, 5: 10}
	partitions := CreateParitions(n, parts, maxCounts, partLen)

	assert.Equal(t, 2, len(partitions))
	for _, part := range partitions {
		fmt.Println(part)
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

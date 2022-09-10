package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	shape := NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}})
	height := 2

	g := NewGame(shape, height)

	assert.True(t, g.Shape.IsEqual(shape), "Shape is wrong")
	assert.True(t, g.Volume.IsEqual(shape.Extrude(height)), "Wrong volume")
}

func TestClear(t *testing.T) {
	shape := NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}})
	height := 3
	g := NewGame(shape, height)

	g.Volume.Set(0, 1, 0, 1)
	g.Volume.Set(0, 1, 1, 1)
	g.Volume.Set(1, 0, 0, 1)

	g.Clear()

	assert.True(t, g.Volume.IsEqual(shape.Extrude(height)), "Clear() didn't produce the right result")
}

func TestClone(t *testing.T) {
	shape := NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}})
	height := 2
	g := NewGame(shape, height)
	c := g.Clone()

	assert.True(t, g.Shape.IsEqual(c.Shape), "Shape does not match")
	assert.True(t, g.Volume.IsEqual(c.Volume), "Volume does not match")
}

func TestTryAddBlock(t *testing.T) {
	shape := NewArray2dFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}})
	g := NewGame(shape, 2)
	blockShape := NewBlock8().Shapes[0]
	pos := Vector{0, 0, 0}

	assert.True(t, g.TryAddBlock(blockShape, pos), "TryAddBlock returned no success where it should")

	exp := NewArray3dFromData([][][]int8{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}})
	assert.True(t, exp.IsEqual(g.Volume), "The resulting volume after TryAddBlock is not as expected")
}

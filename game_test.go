package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	shape := Array2d{{-1, 0, -1}, {-1, 0, 0}}
	height := 2
	g := NewGame(shape, height)

	assert.Equal(t, height, g.Height, "Height is wrong")
	assert.True(t, Equal2DArray(shape, g.Shape), "Shape is wrong")
	assert.Equal(t, len(shape), g.Xdim, "Wrong Xdim")
	assert.Equal(t, len(shape[0]), g.Ydim, "Wrong Ydim")
	assert.True(t, Equal3DArray(g.Volume, Extrude2DArray(shape, height)), "Wrong volume")
}

func TestClear(t *testing.T) {
	shape := Array2d{{-1, 0, -1}, {-1, 0, 0}}
	height := 3
	g := NewGame(shape, height)

	g.Volume[0][1][0] = 1
	g.Volume[0][1][1] = 1
	g.Volume[1][0][0] = 1

	g.Clear()

	assert.True(t, Equal3DArray(g.Volume, Extrude2DArray(shape, height)))
}

func TestClone(t *testing.T) {
	shape := Array2d{{-1, 0, -1}, {-1, 0, 0}}
	height := 2
	g := NewGame(shape, height)
	c := g.Clone()

	assert.True(t, Equal2DArray(g.Shape, c.Shape), "Shape does not match")
	assert.Equal(t, g.Height, c.Height, "Height does not match")
	assert.Equal(t, g.Xdim, c.Xdim, "Xdim does not match")
	assert.Equal(t, g.Ydim, c.Ydim, "Ydim does not match")
	assert.True(t, Equal3DArray(g.Volume, c.Volume), "Volume does not match")
}

func TestTryAddBlock(t *testing.T) {
	shape := Array2d{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}}
	g := NewGame(shape, 2)
	blockShape := NewBlock8().Shapes[0]
	pos := Vector{0, 0, 0}

	assert.True(t, g.TryAddBlock(blockShape, pos), "TryAddBlock returned no success where it should")

	exp := Array3d{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}}
	assert.True(t, Equal3DArray(exp, g.Volume), "The resulting volume after TryAddBlock is not as expected")
}

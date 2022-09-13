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

func TestGameString(t *testing.T) {
	g := NewGame(NewArray2d(5, 3), 3)
	s := g.String()
	assert.True(t, len(s) > 10)
}

func TestGameSolution(t *testing.T) {
	f := NewBlockFactory()
	b := []*Block{f.Get(1), f.Get(2)}
	gs := NewGameSolution(b, []*Array3d{b[0].Shapes[0], b[1].Shapes[0]}, []Vector{{0, 0, 0}, {0, 0, 0}})
	s := gs.String()
	assert.True(t, len(s) > 10)
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
	origVolume := g.Volume.Clone()
	blockShape := NewBlock8().Shapes[0]
	pos := Vector{0, 0, 0}

	// test a case where TryAdd should fail
	nok := g.TryAddBlock(blockShape, Vector{3, 4, 1})
	assert.False(t, nok, "TryAddBlock did not return false where it should")
	assert.True(t, g.Volume.IsEqual(origVolume), "The the volume changed after a failed TryAddBlock() call")

	// test a case where it should succeed
	ok := g.TryAddBlock(blockShape, pos)
	assert.True(t, ok, "TryAddBlock returned no success where it should")
	exp := NewArray3dFromData([][][]int8{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}})
	assert.True(t, exp.IsEqual(g.Volume), "The resulting volume after TryAddBlock is not as expected")
}

func TestRemoveBlock(t *testing.T) {
	g := new(Game)
	g.Volume = NewArray3dFromData([][][]int8{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}})
	origVolume := g.Volume.Clone()
	blockShape := NewBlock8().Shapes[0]
	pos := Vector{0, 0, 0}

	// test case where block is outside volume, this should not change the volume
	nok := g.RemoveBlock(blockShape, Vector{3, 4, 1})
	assert.False(t, nok, "RemoveBlock did not return false where it should")
	assert.True(t, g.Volume.IsEqual(origVolume), "The the volume changed after a failed RemoveBlock() call")

	// test case where removal works
	shape := NewArray2dFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}})
	exp := NewGame(shape, 2)
	ok := g.RemoveBlock(blockShape, pos)
	assert.True(t, ok)
	assert.True(t, exp.Volume.IsEqual(g.Volume))
}

func TestSolveNoSolution(t *testing.T) {
	f := NewBlockFactory()
	blocks := []*Block{f.Get(8), f.Get(9), f.Get(12), f.Get(14)}
	area := NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	g := NewGame(area, 2)

	solutions := g.Solve(blocks)

	assert.Equal(t, 0, len(solutions), "Expected 0 solutions, but found %d", len(solutions))
}

func TestSolve(t *testing.T) {
	f := NewBlockFactory()
	blocks := []*Block{f.Get(8), f.Get(9), f.Get(12), f.Get(16)}
	area := NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	g := NewGame(area, 2)

	solutions := g.Solve(blocks)

	assert.Equal(t, 6, len(solutions), "Expected 6 solutions, but found %d", len(solutions))
}

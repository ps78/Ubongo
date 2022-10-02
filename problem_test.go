package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProblem(t *testing.T) {
	bf := GetBlockFactory()
	shape := NewArray2dFromData([][]int8{{0, 0}, {-1, 0}})
	bs := NewBlockset(bf.Green_L, bf.Blue_v)
	p := NewProblem(shape, 2, bs)

	assert.True(t, shape.IsEqual(p.Shape))
	assert.Equal(t, 2, p.Height)
	assert.True(t, bs.IsEqual(p.Blocks))
}

func TestProblemString(t *testing.T) {
	bf := GetBlockFactory()
	p := NewProblem(NewArray2d(2, 2), 2, NewBlockset(bf.Blue_flash))
	s := p.String()
	assert.True(t, len(s) > 10)
}

func TestProblemClone(t *testing.T) {
	bf := GetBlockFactory()
	shape := NewArray2dFromData([][]int8{{0, 0}, {-1, 0}})
	bs := NewBlockset(bf.Green_L, bf.Blue_v)
	p := NewProblem(shape, 2, bs)
	c := p.Clone()

	assert.Equal(t, p.Height, c.Height)
	assert.True(t, shape.IsEqual(p.Shape))
	assert.True(t, bs.IsEqual(c.Blocks))
	assert.Equal(t, p.BoundingBox, c.BoundingBox)
	assert.Equal(t, p.Area, c.Area)
}

func TestProblemIsEqual(t *testing.T) {
	bf := GetBlockFactory()
	a := NewProblem(NewArray2d(2, 2), 2, NewBlockset(bf.Blue_flash))
	b := NewProblem(NewArray2d(2, 2), 2, NewBlockset(bf.Blue_flash))
	c := NewProblem(NewArray2d(3, 2), 3, NewBlockset(bf.Blue_flash))
	d := NewProblem(NewArray2d(2, 2), 2, NewBlockset(bf.Red_flash))
	assert.True(t, a.IsEqual(b))
	assert.False(t, a.IsEqual(c))
	assert.False(t, a.IsEqual(d))
	assert.False(t, a.IsEqual(nil))
}

func TestGenerateProblems(t *testing.T) {
	fp := GetCardFactory()
	fb := GetBlockFactory()

	shape := fp.Get(Easy, 1).Problems[1].Shape
	problems := GenerateProblems(fb, shape, 3, 5, 10)

	assert.Equal(t, 10, len(problems))

	// check that each problem has a solution
	for _, p := range problems {
		g := NewGame(p)
		solutions := g.Solve()
		assert.Less(t, 0, len(solutions))
	}
}

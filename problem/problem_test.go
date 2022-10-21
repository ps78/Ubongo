package problem_test

import (
	"testing"

	"ubongo/base/array2d"

	"ubongo/blockfactory"
	"ubongo/blockset"
	. "ubongo/problem"

	"github.com/stretchr/testify/assert"
)

func TestNewProblem(t *testing.T) {
	bf := blockfactory.GetBlockFactory()
	shape := array2d.NewFromData([][]int8{{0, 0}, {-1, 0}})
	bs := blockset.NewBlockset(bf.Green_L, bf.Blue_v)
	p := NewProblem(shape, 2, bs)

	assert.True(t, shape.IsEqual(p.Shape))
	assert.Equal(t, 2, p.Height)
	assert.True(t, bs.IsEqual(p.Blocks))
}

func TestProblemString(t *testing.T) {
	bf := blockfactory.GetBlockFactory()
	p := NewProblem(array2d.New(2, 2), 2, blockset.NewBlockset(bf.Blue_flash))
	s := p.String()
	assert.True(t, len(s) > 10)
}

func TestProblemClone(t *testing.T) {
	bf := blockfactory.GetBlockFactory()
	shape := array2d.NewFromData([][]int8{{0, 0}, {-1, 0}})
	bs := blockset.NewBlockset(bf.Green_L, bf.Blue_v)
	p := NewProblem(shape, 2, bs)
	c := p.Clone()

	assert.Equal(t, p.Height, c.Height)
	assert.True(t, shape.IsEqual(p.Shape))
	assert.True(t, bs.IsEqual(c.Blocks))
	assert.Equal(t, p.BoundingBox, c.BoundingBox)
	assert.Equal(t, p.Area, c.Area)
}

func TestProblemIsEqual(t *testing.T) {
	bf := blockfactory.GetBlockFactory()
	a := NewProblem(array2d.New(2, 2), 2, blockset.NewBlockset(bf.Blue_flash))
	b := NewProblem(array2d.New(2, 2), 2, blockset.NewBlockset(bf.Blue_flash))
	c := NewProblem(array2d.New(3, 2), 3, blockset.NewBlockset(bf.Blue_flash))
	d := NewProblem(array2d.New(2, 2), 2, blockset.NewBlockset(bf.Red_flash))
	assert.True(t, a.IsEqual(b))
	assert.False(t, a.IsEqual(c))
	assert.False(t, a.IsEqual(d))
	assert.False(t, a.IsEqual(nil))
}

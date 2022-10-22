package problem_test

import (
	"testing"

	"ubongo/base/array2d"

	"ubongo/blockfactory"
	"ubongo/blockset"
	. "ubongo/problem"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	bf := blockfactory.Get()
	shape := array2d.NewFromData([][]int8{{0, 0}, {-1, 0}})
	bs := blockset.New(bf.Green_L, bf.Blue_v)
	p := New(shape, 2, bs)

	assert.True(t, shape.Equals(p.Shape))
	assert.Equal(t, 2, p.Height)
	assert.True(t, bs.Equals(p.Blocks))
}

func TestString(t *testing.T) {
	bf := blockfactory.Get()
	p := New(array2d.New(2, 2), 2, blockset.New(bf.Blue_flash))
	s := p.String()
	assert.True(t, len(s) > 10)

	var nilProb *P = nil
	assert.Equal(t, "(nil)", nilProb.String())
}

func TestClone(t *testing.T) {
	bf := blockfactory.Get()
	shape := array2d.NewFromData([][]int8{{0, 0}, {-1, 0}})
	bs := blockset.New(bf.Green_L, bf.Blue_v)
	p := New(shape, 2, bs)
	c := p.Clone()

	assert.Equal(t, p.Height, c.Height)
	assert.True(t, shape.Equals(p.Shape))
	assert.True(t, bs.Equals(c.Blocks))
	assert.Equal(t, p.BoundingBox, c.BoundingBox)
	assert.Equal(t, p.Area, c.Area)

	var nilProb *P = nil
	assert.Nil(t, nilProb.Clone())
}

func TestEquals(t *testing.T) {
	bf := blockfactory.Get()
	a := New(array2d.New(2, 2), 2, blockset.New(bf.Blue_flash))
	b := New(array2d.New(2, 2), 2, blockset.New(bf.Blue_flash))
	c := New(array2d.New(3, 2), 3, blockset.New(bf.Blue_flash))
	d := New(array2d.New(2, 2), 2, blockset.New(bf.Red_flash))
	assert.True(t, a.Equals(b))
	assert.False(t, a.Equals(c))
	assert.False(t, a.Equals(d))
	assert.False(t, a.Equals(nil))
}

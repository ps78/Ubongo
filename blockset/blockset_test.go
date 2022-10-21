package blockset_test

import (
	"testing"
	"ubongo/block"
	"ubongo/blockfactory"
	. "ubongo/blockset"
	"ubongo/card"
	"ubongo/cardfactory"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	f := blockfactory.Get()
	bs := New(f.ByNumber(10), f.ByNumber(2), f.ByNumber(2), f.ByNumber(4))
	assert.Equal(t, 3, bs.Count)
	assert.Equal(t, 2, bs.Get(0).Number)
	assert.Equal(t, 4, bs.Get(1).Number)
	assert.Equal(t, 10, bs.Get(2).Number)
}

func TestGet(t *testing.T) {
	f := blockfactory.Get()
	bs := New(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	assert.Equal(t, 1, bs.Get(0).Number)
	assert.Equal(t, 5, bs.Get(1).Number)
	assert.Equal(t, 7, bs.Get(2).Number)
	assert.Nil(t, bs.Get(99))

	var nilBs *S = nil
	assert.Nil(t, nilBs.Get(99))
}

func TestString(t *testing.T) {
	f := blockfactory.Get()
	bs := New(f.ByNumber(10), f.ByNumber(2), f.ByNumber(4))
	s := bs.String()
	assert.True(t, len(s) > 10)

	var nilBs *S = nil
	assert.Equal(t, "(nil)", nilBs.String())
}

func TestAsSlice(t *testing.T) {
	f := blockfactory.Get()
	bs := New(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	s := bs.AsSlice()
	expected := []*block.B{f.ByNumber(1), f.ByNumber(5), f.ByNumber(7)}

	assert.Equal(t, len(expected), len(s))
	for i := 0; i < len(s); i++ {
		assert.Equal(t, expected[i].Number, s[i].Number)
	}

	var nilBs *S = nil
	assert.Nil(t, nilBs.AsSlice())
}

func TestAdd(t *testing.T) {
	f := blockfactory.Get()
	bs := New()

	bs.Add(f.ByNumber(15))
	bs.Add(f.ByNumber(16))
	bs.Add(f.ByNumber(16))
	bs.Add(f.ByNumber(2), f.ByNumber(1), f.ByNumber(7))
	bs.Add(nil)
	bs.Add()

	assert.Equal(t, 5, bs.Count)
	assert.Equal(t, 1, bs.Get(0).Number)
	assert.Equal(t, 2, bs.Get(1).Number)
	assert.Equal(t, 7, bs.Get(2).Number)
	assert.Equal(t, 15, bs.Get(3).Number)
	assert.Equal(t, 16, bs.Get(4).Number)

	var nilBs *S = nil
	assert.NotPanics(t, func() { nilBs.Add() })
}

func TestRemove(t *testing.T) {
	f := blockfactory.Get()
	bs := New(f.ByNumber(7), f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	bs.Remove(7)
	assert.Equal(t, 2, bs.Count)
	assert.Equal(t, 1, bs.Get(0).Number)
	assert.Equal(t, 5, bs.Get(1).Number)

	var nilBs *S = nil
	assert.NotPanics(t, func() { nilBs.Remove(99) })
}

func TestRemoveAt(t *testing.T) {
	f := blockfactory.Get()
	bs := New(f.ByNumber(7), f.ByNumber(7), f.ByNumber(5), f.ByNumber(1), f.ByNumber(16))

	assert.Equal(t, 4, bs.Count)

	bs.RemoveAt(3) // remove last
	assert.Equal(t, 3, bs.Count)
	assert.Equal(t, 1, bs.Get(0).Number)
	assert.Equal(t, 5, bs.Get(1).Number)
	assert.Equal(t, 7, bs.Get(2).Number)
	assert.Nil(t, bs.Get(3))

	bs.RemoveAt(99) // remove non-existing
	assert.Equal(t, 3, bs.Count)

	bs.RemoveAt(1) // remove from the center
	assert.Equal(t, 2, bs.Count)
	assert.Equal(t, 1, bs.Get(0).Number)
	assert.Equal(t, 7, bs.Get(1).Number)

	bs.RemoveAt(0) // remove first
	assert.Equal(t, 1, bs.Count)
	assert.Equal(t, 7, bs.Get(0).Number)

	var nilBs *S = nil
	assert.NotPanics(t, func() { nilBs.RemoveAt(99) })
}

func TestBlocksetIsEqual(t *testing.T) {
	f := blockfactory.Get()

	a := New()
	b := New(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))
	c := New(f.ByNumber(5), f.ByNumber(1), f.ByNumber(7))
	d := New(f.ByNumber(5), f.ByNumber(1), f.ByNumber(7), f.ByNumber(7))
	var e *S = nil

	assert.True(t, a.IsEqual(a))
	assert.True(t, b.IsEqual(c))
	assert.True(t, d.IsEqual(d))
	assert.False(t, a.IsEqual(b))
	assert.True(t, b.IsEqual(d))
	assert.False(t, e.IsEqual(a))
	assert.False(t, a.IsEqual(e))
	assert.True(t, e.IsEqual(e))
}

func TestContains(t *testing.T) {
	f := blockfactory.Get()
	bs := New(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	assert.True(t, bs.Contains(7))
	assert.False(t, bs.Contains(10))
	assert.False(t, bs.Contains(0))

	var nilBs *S = nil
	assert.False(t, nilBs.Contains(0))
}

func TestClone(t *testing.T) {
	f := blockfactory.Get()
	orig := New(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))
	cl := orig.Clone()

	assert.Equal(t, orig.Count, cl.Count)
	for i := 0; i < orig.Count; i++ {
		assert.Equal(t, orig.Get(i).Number, cl.Get(i).Number)
	}

	var nilBs *S = nil
	assert.Nil(t, nilBs.Clone())
}

func TestVolume(t *testing.T) {
	f := blockfactory.Get()
	a := New()
	b := New(f.ByName(block.Blue, "v"), f.ByName(block.Red, "stool"))

	assert.Equal(t, 0, a.Volume())
	assert.Equal(t, 8, b.Volume())

	var nilBs *S = nil
	assert.Equal(t, 0, nilBs.Volume())
}

func TestContainsBlockset(t *testing.T) {
	f := cardfactory.Get()
	bs1 := f.Cards[card.Difficult][1].Problems[1].Blocks
	bs2 := f.Cards[card.Difficult][1].Problems[2].Blocks
	bs3 := f.Cards[card.Difficult][1].Problems[3].Blocks
	bs4 := f.Cards[card.Difficult][1].Problems[4].Blocks

	set := []*S{bs1, bs2, bs3}

	assert.True(t, ContainsBlockset(set, bs1))
	assert.True(t, ContainsBlockset(set, bs2))
	assert.True(t, ContainsBlockset(set, bs3))
	assert.False(t, ContainsBlockset(set, bs4))
	assert.False(t, ContainsBlockset(set, nil))
	assert.False(t, ContainsBlockset(nil, nil))
	assert.False(t, ContainsBlockset(set, nil))
	assert.False(t, ContainsBlockset(nil, bs1))
}

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockset(t *testing.T) {
	f := GetBlockFactory()
	bs := NewBlockset(f.ByNumber(10), f.ByNumber(2), f.ByNumber(4))
	assert.Equal(t, 3, bs.Count)
	assert.Equal(t, 2, bs.At(0).Number)
	assert.Equal(t, 4, bs.At(1).Number)
	assert.Equal(t, 10, bs.At(2).Number)
}

func TestBlocksetAdd(t *testing.T) {
	f := GetBlockFactory()
	bs := NewBlockset()

	bs.Add(f.ByNumber(15))
	bs.Add(f.ByNumber(16))
	bs.Add(f.ByNumber(2), f.ByNumber(1), f.ByNumber(7))
	bs.Add(nil)
	bs.Add()

	assert.Equal(t, 5, bs.Count)
	assert.Equal(t, 1, bs.At(0).Number)
	assert.Equal(t, 2, bs.At(1).Number)
	assert.Equal(t, 7, bs.At(2).Number)
	assert.Equal(t, 15, bs.At(3).Number)
	assert.Equal(t, 16, bs.At(4).Number)
}

func TestBlocksetClone(t *testing.T) {
	f := GetBlockFactory()
	orig := NewBlockset(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))
	cl := orig.Clone()

	assert.Equal(t, orig.Count, cl.Count)
	for i := range orig.items {
		assert.Equal(t, orig.items[i].Number, cl.items[i].Number)
	}
}

func TestBlocksetContains(t *testing.T) {
	f := GetBlockFactory()
	bs := NewBlockset(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	assert.True(t, bs.Contains(7))
	assert.False(t, bs.Contains(10))
	assert.False(t, bs.Contains(0))
}

func TestBlocksetAt(t *testing.T) {
	f := GetBlockFactory()
	bs := NewBlockset(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	assert.Equal(t, 1, bs.At(0).Number)
	assert.Equal(t, 5, bs.At(1).Number)
	assert.Equal(t, 7, bs.At(2).Number)
	assert.Nil(t, bs.At(99))
}

func TestBlocksetAsSlice(t *testing.T) {
	f := GetBlockFactory()
	bs := NewBlockset(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	s := bs.AsSlice()
	expected := []*Block{f.ByNumber(1), f.ByNumber(5), f.ByNumber(7)}

	assert.Equal(t, len(expected), len(s))
	for i := 0; i < len(s); i++ {
		assert.Equal(t, expected[i].Number, s[i].Number)
	}
}

func TestBlocksetIsEqual(t *testing.T) {
	f := GetBlockFactory()

	a := NewBlockset()
	b := NewBlockset(f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))
	c := NewBlockset(f.ByNumber(5), f.ByNumber(1), f.ByNumber(7))
	d := NewBlockset(f.ByNumber(5), f.ByNumber(1), f.ByNumber(7), f.ByNumber(7))

	assert.True(t, a.IsEqual(a))
	assert.True(t, b.IsEqual(c))
	assert.True(t, d.IsEqual(d))
	assert.False(t, a.IsEqual(b))
	assert.False(t, b.IsEqual(d))
}

func TestBlocksetVolume(t *testing.T) {
	f := GetBlockFactory()
	a := NewBlockset()
	b := NewBlockset(f.ByName(Blue, "v"), f.ByName(Red, "stool"))

	assert.Equal(t, 0, a.Volume())
	assert.Equal(t, 8, b.Volume())
}

func TestBlocksetRemove(t *testing.T) {
	f := GetBlockFactory()
	bs := NewBlockset(f.ByNumber(7), f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	bs.Remove(7)
	assert.Equal(t, 2, bs.Count)
	assert.Equal(t, 1, bs.items[0].Number)
	assert.Equal(t, 5, bs.items[1].Number)
}

func TestBlocksetRemoveAt(t *testing.T) {
	f := GetBlockFactory()
	bs := NewBlockset(f.ByNumber(7), f.ByNumber(7), f.ByNumber(5), f.ByNumber(1))

	bs.RemoveAt(3) // remove last
	assert.Equal(t, 3, bs.Count)
	assert.Equal(t, 1, bs.At(0).Number)
	assert.Equal(t, 5, bs.At(1).Number)
	assert.Equal(t, 7, bs.At(2).Number)

	bs.RemoveAt(99) // remove non-existing
	assert.Equal(t, 3, bs.Count)

	bs.RemoveAt(1) // remove from the center
	assert.Equal(t, 2, bs.Count)
	assert.Equal(t, 1, bs.At(0).Number)
	assert.Equal(t, 7, bs.At(1).Number)

	bs.RemoveAt(0) // remove first
	assert.Equal(t, 1, bs.Count)
	assert.Equal(t, 7, bs.At(0).Number)
}

func TestGenerateBlocksets(t *testing.T) {
	f := GetBlockFactory()
	vol := 21
	blockCount := 5
	blocksets := GenerateBlocksets(f, vol, blockCount, 10000)

	// this test should return less than the requested number of results
	assert.LessOrEqual(t, len(blocksets), 10000)
	for _, s := range blocksets {
		assert.Equal(t, blockCount, s.Count)
		assert.Equal(t, vol, s.Volume())
		fmt.Println(s)
	}

	// this test should return exactly 3 results
	blocksets = GenerateBlocksets(f, vol, blockCount, 3)
	assert.LessOrEqual(t, len(blocksets), 3)
}

func TestContainsBlockset(t *testing.T) {
	f := GetCardFactory()
	bs1 := f.Cards[Difficult][1].Problems[1].Blocks
	bs2 := f.Cards[Difficult][1].Problems[2].Blocks
	bs3 := f.Cards[Difficult][1].Problems[3].Blocks
	bs4 := f.Cards[Difficult][1].Problems[4].Blocks

	set := []*Blockset{bs1, bs2, bs3}

	assert.True(t, ContainsBlockset(set, bs1))
	assert.True(t, ContainsBlockset(set, bs2))
	assert.True(t, ContainsBlockset(set, bs3))
	assert.False(t, ContainsBlockset(set, bs4))
	assert.False(t, ContainsBlockset(set, nil))
}

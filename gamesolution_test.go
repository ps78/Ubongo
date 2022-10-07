package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGamesolution(t *testing.T) {
	f := GetBlockFactory()

	b1 := f.Blue_flash
	b2 := f.Green_L
	s1 := b1.Shapes[1]
	s2 := b2.Shapes[7]
	v1 := Vector{1, 0, 0}
	v2 := Vector{0, 1, 0}
	gs := NewGameSolution([]*Block{b1, b2}, []*Array3d{s1, s2}, []Vector{v1, v2})

	assert.Equal(t, 2, len(gs.Blocks))
	assert.Equal(t, 2, len(gs.Shapes))
	assert.Equal(t, 2, len(gs.Shifts))
	assert.Equal(t, b1, gs.Blocks[0])
	assert.Equal(t, b2, gs.Blocks[1])
	assert.True(t, s1.IsEqual(gs.Shapes[0]))
	assert.True(t, s2.IsEqual(gs.Shapes[1]))
	assert.Equal(t, v1, gs.Shifts[0])
	assert.Equal(t, v2, gs.Shifts[1])
}

func TestGameSolutionString(t *testing.T) {
	f := GetBlockFactory()
	gs := NewGameSolution([]*Block{f.Blue_v}, []*Array3d{f.Blue_v.Shapes[1]}, []Vector{{0, 1, 0}})
	s := gs.String()
	assert.True(t, len(s) > 10)
}

func TestGameSolutionGetCenterOfGravity(t *testing.T) {
	cf := GetCardFactory()
	p := cf.Get(Difficult, 3).Problems[7]
	g := NewGame(p)
	solutions := g.Solve()

	actual := solutions[0].GetCenterOfGravity()
	expected := Vectorf{2, 1.625, 1}
	assert.Equal(t, expected, actual)
}

package main

import (
	"testing"

	"ubongo/utils"

	"github.com/stretchr/testify/assert"
)

func TestNewGamesolution(t *testing.T) {
	f := GetBlockFactory()

	b1 := f.Blue_flash
	b2 := f.Green_L
	s1 := 1
	s2 := 7
	v1 := utils.Vector{1, 0, 0}
	v2 := utils.Vector{0, 1, 0}
	gs := NewGameSolution([]*Block{b1, b2}, []int{s1, s2}, []utils.Vector{v1, v2})

	assert.Equal(t, 2, len(gs.Blocks))
	assert.Equal(t, 2, len(gs.ShapeIndex))
	assert.Equal(t, 2, len(gs.Shifts))
	assert.Equal(t, b1, gs.Blocks[0])
	assert.Equal(t, b2, gs.Blocks[1])
	assert.Equal(t, s1, gs.ShapeIndex[0])
	assert.Equal(t, s2, gs.ShapeIndex[1])
	assert.Equal(t, v1, gs.Shifts[0])
	assert.Equal(t, v2, gs.Shifts[1])
}

func TestGameSolutionString(t *testing.T) {
	f := GetBlockFactory()
	gs := NewGameSolution([]*Block{f.Blue_v}, []int{1}, []utils.Vector{{0, 1, 0}})
	s := gs.String()
	assert.True(t, len(s) > 10)
}

func TestGameSolutionGetCenterOfGravity(t *testing.T) {
	cf := GetCardFactory()
	p := cf.Get(Difficult, 3).Problems[7]
	g := NewGame(p)
	solutions := g.Solve()

	actual := solutions[0].GetCenterOfGravity()
	expected := utils.Vectorf{2, 1.625, 1}
	assert.Equal(t, expected, actual)
}

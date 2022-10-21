package gamesolution_test

import (
	"testing"

	"ubongo/base/vector"
	"ubongo/base/vectorf"
	"ubongo/block"
	"ubongo/blockfactory"
	"ubongo/card"
	"ubongo/cardfactory"
	"ubongo/game"
	. "ubongo/gamesolution"

	"github.com/stretchr/testify/assert"
)

func TestNewGamesolution(t *testing.T) {
	f := blockfactory.GetBlockFactory()

	b1 := f.Blue_flash
	b2 := f.Green_L
	s1 := 1
	s2 := 7
	v1 := vector.V{1, 0, 0}
	v2 := vector.V{0, 1, 0}
	gs := NewGameSolution([]*block.Block{b1, b2}, []int{s1, s2}, []vector.V{v1, v2})

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
	f := blockfactory.GetBlockFactory()
	gs := NewGameSolution([]*block.Block{f.Blue_v}, []int{1}, []vector.V{{0, 1, 0}})
	s := gs.String()
	assert.True(t, len(s) > 10)
}

func TestGameSolutionGetCenterOfGravity(t *testing.T) {
	cf := cardfactory.GetCardFactory()
	p := cf.Get(card.Difficult, 3).Problems[7]
	g := game.NewGame(p)
	solutions := g.Solve()

	actual := solutions[0].GetCenterOfGravity()
	expected := vectorf.V{2, 1.625, 1}
	assert.Equal(t, expected, actual)
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUbongoDifficultyString(t *testing.T) {

	assert.Equal(t, "Unknown", UbongoDifficulty(99).String())
	assert.NotEqual(t, "Unknown", Easy.String())
	assert.NotEqual(t, "Unknown", Difficult.String())
	assert.NotEqual(t, "Unknown", Insane.String())
}

func TestProblemString(t *testing.T) {
	p := NewProblem("A1", 1, NewArray2d(5, 3), []*Block{})
	s := p.String()
	assert.True(t, len(s) > 10)
}

func TestNewProblem(t *testing.T) {
	shape := NewArray2dFromData([][]int8{{-1, 0, 0}, {0, 0, 0}})
	f := NewBlockFactory()
	blocks := []*Block{f.Get(1), f.Get(2)}
	p := NewProblem("A1", 1, shape, blocks)

	assert.Equal(t, "A1", p.CardId)
	assert.Equal(t, Easy, p.Difficulty)
	assert.Equal(t, "Elephant", p.Animal)
	assert.Equal(t, 1, p.Number)
	assert.True(t, p.Shape.IsEqual(shape))
	assert.Equal(t, len(blocks), len(p.Blocks))
}

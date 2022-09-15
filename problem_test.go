package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUbongoDifficultyString(t *testing.T) {

	assert.Equal(t, "unknown", strings.ToLower(UbongoDifficulty(99).String()))
	assert.NotEqual(t, "unknown", strings.ToLower(Easy.String()))
	assert.NotEqual(t, "unknown", strings.ToLower(Difficult.String()))
	assert.NotEqual(t, "unknown", strings.ToLower(Insane.String()))
}

func TestProblemString(t *testing.T) {
	p := NewProblem(1, Easy, 2, NewArray2d(5, 3), []*Block{})
	s := p.String()
	assert.True(t, len(s) > 10)
}

func TestNewProblem(t *testing.T) {
	f := GetBlockFactory()

	blocks := map[UbongoDifficulty]([]*Block){
		Easy:      []*Block{f.Get(1), f.Get(5), f.Get(9)},
		Difficult: []*Block{f.Get(2), f.Get(6), f.Get(10), f.Get(12)},
		Insane:    []*Block{f.Get(3), f.Get(7), f.Get(12), f.Get(13), f.Get(16)}}

	heights := map[UbongoDifficulty]int{
		Easy:      2,
		Difficult: 2,
		Insane:    3}

	shape := NewArray2dFromData([][]int8{{-1, 0, 0}, {0, 0, 0}})
	for _, diff := range []UbongoDifficulty{Easy, Difficult, Insane} {
		for cardNum := 1; cardNum <= 36; cardNum++ {
			for probNum := 1; probNum <= 10; probNum++ {
				p := NewProblem(cardNum, diff, probNum, shape, blocks[diff])
				assert.Equal(t, diff, p.Difficulty)
				assert.Equal(t, cardNum, p.CardNumber)
				assert.Equal(t, probNum, p.DiceNumber)
				assert.True(t, shape.IsEqual(p.Shape))
				assert.True(t, len(p.Animal) > 0)
				assert.Equal(t, heights[diff], p.Height)
				assert.Equal(t, shape.Count(0), p.Area)
				assert.Equal(t, Vector{2, 3, heights[diff]}, p.BoundingBox)
			}
		}
	}
}

func TestProblemFactoryGet(t *testing.T) {
	f := GetProblemFactory()
	p := f.Get(Difficult, 12, 4)
	assert.NotNil(t, p)

	pnil := f.Get(Easy, 99, 1)
	assert.Nil(t, pnil)
}

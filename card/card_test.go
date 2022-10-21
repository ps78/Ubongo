package card_test

import (
	"strings"
	"testing"

	"ubongo/base/array2d"
	"ubongo/base/vector"
	"ubongo/blockfactory"
	"ubongo/blockset"
	. "ubongo/card"
	"ubongo/cardfactory"
	"ubongo/problem"

	"github.com/stretchr/testify/assert"
)

func TestUbongoDifficultyString(t *testing.T) {
	assert.Equal(t, "unknown", strings.ToLower(UbongoDifficulty(99).String()))
	assert.NotEqual(t, "unknown", strings.ToLower(Easy.String()))
	assert.NotEqual(t, "unknown", strings.ToLower(Difficult.String()))
	assert.NotEqual(t, "unknown", strings.ToLower(Insane.String()))
}

func TestUbongoAnimalString(t *testing.T) {
	assert.Equal(t, "(n/a)", strings.ToLower(UbongoAnimal(99).String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Elephant.String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Gazelle.String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Snake.String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Gnu.String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Ostrich.String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Rhino.String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Giraffe.String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Zebra.String()))
	assert.NotEqual(t, "(n/a)", strings.ToLower(Warthog.String()))
}

func TestCardString(t *testing.T) {
	probs := map[int]*problem.Problem{}
	p := NewCard(1, Easy, Elephant, probs)
	s := p.String()
	assert.True(t, len(s) > 10)
}

func TestCardVerbousString(t *testing.T) {
	c := cardfactory.GetCardFactory().Get(Difficult, 1)
	s := c.VerbousString()
	assert.True(t, len(s) > 100)
}

func TestNewCard(t *testing.T) {
	f := blockfactory.GetBlockFactory()

	animal := Zebra
	shape := array2d.NewFromData([][]int8{{-1, 0, 0}, {0, 0, 0}})
	height := 3
	cardNum := 42
	diceNum := 7
	diff := Insane
	bs := blockset.NewBlockset(f.ByNumber(3), f.ByNumber(7), f.ByNumber(12), f.ByNumber(13), f.ByNumber(16))

	var problems map[int]*problem.Problem = map[int]*problem.Problem{
		diceNum: problem.NewProblem(shape, height, bs)}

	c := NewCard(cardNum, diff, animal, problems)

	assert.Equal(t, cardNum, c.CardNumber)
	assert.Equal(t, diff, c.Difficulty)
	assert.Equal(t, animal, c.Animal)
	assert.Equal(t, 1, len(c.Problems))
	assert.True(t, shape.IsEqual(c.Problems[diceNum].Shape))
	assert.Equal(t, height, c.Problems[diceNum].Height)
	assert.Equal(t, shape.Count(0), c.Problems[diceNum].Area)
	assert.Equal(t, vector.V{shape.DimX, shape.DimY, height}, c.Problems[diceNum].BoundingBox)
}

func TestCardClone(t *testing.T) {
	o := cardfactory.GetCardFactory().Get(Difficult, 13)
	c := o.Clone()

	assert.Equal(t, o.Animal, c.Animal)
	assert.Equal(t, o.CardNumber, c.CardNumber)
	assert.Equal(t, o.Difficulty, c.Difficulty)
	assert.Equal(t, len(o.Problems), len(c.Problems))
	for k, v := range o.Problems {
		assert.True(t, v.IsEqual(c.Problems[k]))
	}
}

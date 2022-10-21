package cardfactory_test

import (
	"testing"
	"ubongo/card"
	. "ubongo/cardfactory"

	"github.com/stretchr/testify/assert"
)

func TestCardFactoryInstance(t *testing.T) {
	f := GetCardFactory()

	easyCount := 0
	easyCards := f.Cards[card.Easy]
	for cardNum, card := range easyCards {
		assert.Equal(t, cardNum, card.CardNumber)
		easyCount += len(card.Problems)
	}

	diffCount := 0
	diffProbs := f.Cards[card.Difficult]
	for cardNum, card := range diffProbs {
		assert.Equal(t, cardNum, card.CardNumber)
		diffCount += len(card.Problems)
	}

	assert.Equal(t, 144, easyCount)
	assert.Equal(t, 360, diffCount)
}

func TestCardFactoryGet(t *testing.T) {
	f := GetCardFactory()

	c := f.Get(card.Difficult, 12)
	assert.NotNil(t, c)
	assert.Equal(t, card.Difficult, c.Difficulty)
	assert.Equal(t, 12, c.CardNumber)

	pnil := f.Get(card.Easy, 99)
	assert.Nil(t, pnil)
}

func TestCardFactoryGetByAnimal(t *testing.T) {
	f := GetCardFactory()

	cards := f.GetByAnimal(card.Difficult, card.Zebra)
	assert.Equal(t, 4, len(cards))
	for _, c := range cards {
		assert.Equal(t, card.Zebra, c.Animal)
		assert.Equal(t, card.Difficult, c.Difficulty)
	}
}

func TestCardFactoryGetAll(t *testing.T) {
	f := GetCardFactory()

	easyCards := f.GetAll(card.Easy)
	assert.Equal(t, 36, len(easyCards))
	for _, c := range easyCards {
		assert.Equal(t, card.Easy, c.Difficulty)
	}

	diffCards := f.GetAll(card.Difficult)
	assert.Equal(t, 36, len(diffCards))
	for _, c := range diffCards {
		assert.Equal(t, card.Difficult, c.Difficulty)
	}
}

func TestCardFactoryGetAllProblems(t *testing.T) {
	f := GetCardFactory()

	easyProbs := f.GetAllProblems(card.Easy)
	assert.Equal(t, 144, len(easyProbs))

	diffProbs := f.GetAllProblems(card.Difficult)
	assert.Equal(t, 360, len(diffProbs))
}

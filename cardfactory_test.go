package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardFactoryInstance(t *testing.T) {
	f := GetCardFactory()

	easyCount := 0
	easyCards := f.Cards[Easy]
	for cardNum, card := range easyCards {
		assert.Equal(t, cardNum, card.CardNumber)
		easyCount += len(card.Problems)
	}

	diffCount := 0
	diffProbs := f.Cards[Difficult]
	for cardNum, card := range diffProbs {
		assert.Equal(t, cardNum, card.CardNumber)
		diffCount += len(card.Problems)
	}

	assert.Equal(t, 144, easyCount)
	assert.Equal(t, 360, diffCount)
}

func TestCardFactoryGet(t *testing.T) {
	f := GetCardFactory()

	c := f.Get(Difficult, 12)
	assert.NotNil(t, c)
	assert.Equal(t, Difficult, c.Difficulty)
	assert.Equal(t, 12, c.CardNumber)

	pnil := f.Get(Easy, 99)
	assert.Nil(t, pnil)
}

func TestCardFactoryGetByAnimal(t *testing.T) {
	f := GetCardFactory()

	cards := f.GetByAnimal(Difficult, Zebra)
	assert.Equal(t, 4, len(cards))
	for _, c := range cards {
		assert.Equal(t, Zebra, c.Animal)
		assert.Equal(t, Difficult, c.Difficulty)
	}
}

func TestCardFactoryGetAll(t *testing.T) {
	f := GetCardFactory()

	easyCards := f.GetAll(Easy)
	assert.Equal(t, 36, len(easyCards))
	for _, c := range easyCards {
		assert.Equal(t, Easy, c.Difficulty)
	}

	diffCards := f.GetAll(Difficult)
	assert.Equal(t, 36, len(diffCards))
	for _, c := range diffCards {
		assert.Equal(t, Difficult, c.Difficulty)
	}
}

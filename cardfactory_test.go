package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardFactoryGet(t *testing.T) {
	f := GetCardFactory()

	p := f.Get(Difficult, 12)
	assert.NotNil(t, p)

	pnil := f.Get(Easy, 99)
	assert.Nil(t, pnil)
}

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

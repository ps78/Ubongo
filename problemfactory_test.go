package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemFactoryGet(t *testing.T) {
	f := GetProblemFactory()

	p := f.Get(Difficult, 12, 4)
	assert.NotNil(t, p)

	pnil := f.Get(Easy, 99, 1)
	assert.Nil(t, pnil)
}

func TestProblemFactory(t *testing.T) {
	f := GetProblemFactory()

	easyCount := 0
	easyProbs := f.Problems[Easy]
	for cardNum, probs := range easyProbs {
		for diceNum, p := range probs {
			assert.Equal(t, cardNum, p.CardNumber)
			assert.Equal(t, diceNum, p.DiceNumber)
			easyCount++
		}
	}

	diffCount := 0
	diffProbs := f.Problems[Difficult]
	for cardNum, probs := range diffProbs {
		for diceNum, p := range probs {
			assert.Equal(t, cardNum, p.CardNumber)
			assert.Equal(t, diceNum, p.DiceNumber)
			diffCount++
		}
	}

	assert.Equal(t, 144, easyCount)
	assert.Equal(t, 360, diffCount)
}

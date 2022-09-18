package main

import (
	"sync"
)

// used to create a thread-safe singleton instance of a problemFactory
var onceProblemFactorySingleton sync.Once

// the singleton
var problemFactoryInstance *ProblemFactory

// Returns the singleton instance of the problem factory
func GetProblemFactory() *ProblemFactory {
	onceProblemFactorySingleton.Do(func() {
		bf := GetBlockFactory()

		f := new(ProblemFactory)

		f.Problems = make(map[UbongoDifficulty]map[int]map[int]*Problem)

		for _, p := range createAllProblems(bf) {
			if _, ok := f.Problems[p.Difficulty]; !ok {
				f.Problems[p.Difficulty] = make(map[int]map[int]*Problem)
			}
			if _, ok := f.Problems[p.Difficulty][p.CardNumber]; !ok {
				f.Problems[p.Difficulty][p.CardNumber] = make(map[int]*Problem)
			}
			f.Problems[p.Difficulty][p.CardNumber][p.DiceNumber] = p
		}

		problemFactoryInstance = f
	})
	return problemFactoryInstance
}

// Returns the problem with the given parameters if it exists, nil otherwise
func (f *ProblemFactory) Get(difficulty UbongoDifficulty, cardNumber, diceNumber int) *Problem {
	if _, okDiff := f.Problems[difficulty]; okDiff {
		if _, okCard := f.Problems[difficulty][cardNumber]; okCard {
			if _, okDice := f.Problems[difficulty][cardNumber][diceNumber]; okDice {
				return f.Problems[difficulty][cardNumber][diceNumber]
			}
		}
	}
	return nil
}

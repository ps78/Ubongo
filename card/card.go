// Package card contains the Card type and related sub-types and methods
package card

import (
	"fmt"
	"sort"
	"ubongo/problem"
)

// UbongoDifficulty is an enum representing the difficulty in the game
type UbongoDifficulty int

// Enumeration values of the UbongoDifficulty enum
const (
	// Easy as in the original game: shapes to be filled with 3 blocks
	Easy UbongoDifficulty = iota
	// Difficult as in the original game: shapes to be filled with 4 blocks
	Difficult
	// Insane is not part of the original game: shapes to be filled with 5 blocks
	Insane
)

// String returns a string representation for the UbongoDifficulty enum
func (s UbongoDifficulty) String() string {
	switch s {
	case Easy:
		return "Easy"
	case Difficult:
		return "Difficult"
	case Insane:
		return "Insane"
	}
	return "Unknown"
}

// UbongoAnimal represents the animal on the cards of the original Ubongo game
type UbongoAnimal int

// Enumeration constants of the UbongoAnimal enumeration type
const (
	Elephant UbongoAnimal = iota
	Gazelle
	Snake
	Gnu
	Ostrich
	Rhino
	Giraffe
	Zebra
	Warthog
)

// AllAnimals can be used to list all elements of the UbongoAnimal enumeration
func AllAnimals() []UbongoAnimal {
	return []UbongoAnimal{Elephant, Gazelle, Snake, Gnu, Ostrich, Rhino, Giraffe, Zebra, Warthog}
}

// String returns a string representation of the UbongoAnimal enum
func (s UbongoAnimal) String() string {
	switch s {
	case Elephant:
		return "Elephant"
	case Gazelle:
		return "Gazelle"
	case Snake:
		return "Snake"
	case Gnu:
		return "Gnu"
	case Ostrich:
		return "Ostrich"
	case Rhino:
		return "Rhino"
	case Giraffe:
		return "Giraffe"
	case Zebra:
		return "Zebra"
	case Warthog:
		return "Warthog"
	}
	return "(N/A)"
}

// C represents a physical card of the Ubongo game with multiple problems
type C struct {
	// CardNumber represents the number printed on each side of a card in the original game, without the letter
	CardNumber int

	// Animal is the name of the animal as printed on the original game
	Animal UbongoAnimal

	// Difficulty of the card and consequently of all problems on the card
	Difficulty UbongoDifficulty

	// Problems is the set of problems on the card. The key of the map corresponds to the dice number
	Problems map[int]*problem.P
}

// String returns a string representation of the card
func (c *C) String() string {
	if c == nil {
		return "(nil)"
	} else {
		return fmt.Sprintf("Card %d (%s-%s) %d problems)",
			c.CardNumber, c.Animal, c.Difficulty, len(c.Problems))
	}
}

// VerbousString returns a string representation including all problems for a card
func (c *C) VerbousString() string {
	if c == nil {
		return "(nil)"
	} else {
		s := fmt.Sprintf("Card %02d %s, %s\n",
			c.CardNumber, c.Animal, c.Difficulty)

		type item struct {
			diceNum int
			p       *problem.P
		}
		probs := make([]*item, 0)
		for diceNumber, p := range c.Problems {
			probs = append(probs, &item{diceNumber, p})
		}
		sort.Slice(probs, func(i, j int) bool {
			return probs[i].diceNum < probs[j].diceNum
		})
		for _, p := range probs {
			s += fmt.Sprintf("\t%2d: Vol=%2d, %s\n", p.diceNum, p.p.Blocks.Volume(), p.p.Blocks)
		}
		return s
	}
}

// New creates a card instance
func New(cardNumber int, difficulty UbongoDifficulty, animal UbongoAnimal, problems map[int]*problem.P) *C {
	var c *C = new(C)

	c.CardNumber = cardNumber
	c.Difficulty = difficulty
	c.Animal = animal

	c.Problems = make(map[int]*problem.P)
	for k, v := range problems {
		c.Problems[k] = v
	}

	return c
}

// Clone creates a deep copy of the given card
func (c *C) Clone() *C {
	if c == nil {
		return nil
	} else {
		n := new(C)

		n.CardNumber = c.CardNumber
		n.Difficulty = c.Difficulty
		n.Animal = c.Animal
		n.Problems = make(map[int]*problem.P)
		for k, v := range c.Problems {
			n.Problems[k] = v.Clone()
		}

		return n
	}
}

package main

import (
	"fmt"
)

// UbongoDifficulty is an enum representing the difficulty in the game
type UbongoDifficulty int

const (
	// Easy as in the original game: shapes to be filled with 3 blocks
	Easy UbongoDifficulty = iota
	// Difficult as in the original game: shapes to be filled with 4 blocks
	Difficult
	// Insane is not part of the original game: shapes to be filled with 5 blocks
	Insane
)

// Returns a string representation for the UbongoDifficulty enum
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

// Card represents a physical card of the Ubongo game with multiple problems
type Card struct {
	// CardId represents the number printed on each side of a card in the original game, without the letter
	CardNumber int

	// Animal is the name of the animal as printed on the original game
	Animal UbongoAnimal

	// Difficulty is either e = easy or d=difficult
	Difficulty UbongoDifficulty

	// The problems on the card. the key of the map corresponds to the dice number
	Problems map[int]*Problem
}

// Returns a string representation of the card
func (c Card) String() string {
	return fmt.Sprintf("Card %d (%s-%s) %d problems)",
		c.CardNumber, c.Animal, c.Difficulty, len(c.Problems))
}

// NewCard creates a card instance
func NewCard(cardNumber int, difficulty UbongoDifficulty, animal UbongoAnimal, problems map[int]*Problem) *Card {
	var c *Card = new(Card)

	c.CardNumber = cardNumber
	c.Difficulty = difficulty
	c.Animal = animal

	c.Problems = make(map[int]*Problem)
	for k, v := range problems {
		c.Problems[k] = v
	}

	return c
}

// Creates a deep copy of the given object p
func (c *Card) Clone() *Card {
	n := new(Card)

	n.CardNumber = c.CardNumber
	n.Difficulty = c.Difficulty
	n.Animal = c.Animal
	n.Problems = make(map[int]*Problem)
	for k, v := range c.Problems {
		n.Problems[k] = v.Clone()
	}

	return n
}

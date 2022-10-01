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
	Elephant = iota
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

var animalByCardNum = map[int]UbongoAnimal{
	1:  Elephant,
	2:  Elephant,
	3:  Elephant,
	4:  Elephant,
	5:  Gazelle,
	6:  Gazelle,
	7:  Gazelle,
	8:  Gazelle,
	9:  Snake,
	10: Snake,
	11: Snake,
	12: Snake,
	13: Gnu,
	14: Gnu,
	15: Gnu,
	16: Gnu,
	17: Ostrich,
	18: Ostrich,
	19: Ostrich,
	20: Ostrich,
	21: Rhino,
	22: Rhino,
	23: Rhino,
	24: Rhino,
	25: Giraffe,
	26: Giraffe,
	27: Giraffe,
	28: Giraffe,
	29: Zebra,
	30: Zebra,
	31: Zebra,
	32: Zebra,
	33: Warthog,
	34: Warthog,
	35: Warthog,
	36: Warthog,
}

// Problem represents a single Ubongo problem to solve
type Problem struct {
	// CardId represents the number printed on each side of a card in the original game, without the letter
	CardNumber int

	// Animal is the name of the animal as printed on the original game
	Animal UbongoAnimal

	// Difficulty is either e = easy or d=difficult
	Difficulty UbongoDifficulty

	// DiceNumber is the problem number as printed on the card corresponding to the dice (1..10)
	DiceNumber int

	// Shape is the 2D shape of the puzzle, first is the index X-direction (horizontal, to the right),
	// the second index is the Y-direction (up)
	Shape *Array2d // -1=not part of volume, 0=empty, 1=occupied by a block

	// Height of the volume to fill with the blocks. This is always 2 for the original game
	Height int

	// The area of the problem in unit squares
	Area int

	// Bounding box of the problem volume
	BoundingBox Vector

	// Blocks is an array of the blocks to be used to fill the volume
	Blocks []*Block
}

// Returns a string representation of the problem
func (p Problem) String() string {
	return fmt.Sprintf("Problem: card %d (%s-%s-%d) (%d blocks, area %d, height %d)",
		p.CardNumber, p.Animal, p.Difficulty, p.DiceNumber, len(p.Blocks), p.Area, p.Height)
}

// creates a problem instance
func NewProblem(cardNumber, diceNumber int, difficulty UbongoDifficulty, height int, animal UbongoAnimal, shape *Array2d, blocks []*Block) *Problem {
	var p *Problem = new(Problem)

	p.CardNumber = cardNumber
	p.DiceNumber = diceNumber
	p.Shape = shape.Clone()
	p.Blocks = make([]*Block, len(blocks))
	copy(p.Blocks, blocks)
	p.Difficulty = difficulty
	p.Height = height
	p.Animal = animal
	p.Area = p.Shape.Count(0)
	p.BoundingBox = Vector{p.Shape.DimX, p.Shape.DimY, p.Height}

	return p
}

// Creates a deep copy of the given object p
func (p *Problem) Clone() *Problem {
	c := new(Problem)

	c.CardNumber = p.CardNumber
	c.DiceNumber = p.DiceNumber
	c.Shape = p.Shape.Clone()
	c.Blocks = make([]*Block, len(p.Blocks))
	copy(c.Blocks, p.Blocks)
	c.Difficulty = p.Difficulty
	c.Height = p.Height
	c.Animal = p.Animal
	c.Area = p.Area
	c.BoundingBox = p.BoundingBox

	return c
}

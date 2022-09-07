package main

import (
	"fmt"
)

// UbongoDifficulty is an enum representing the difficulty in the game
type UbongoDifficulty int8

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

// Problem represents a single Ubongo problem to solve
type Problem struct {
	// CardId represents the number printed on each side of a card in the original game
	CardId string

	// Animal is the name of the animal as printed on the original game
	Animal string

	// Difficulty is either e = easy or d=difficult
	Difficulty UbongoDifficulty

	// Number is the problem number as printed on the card (1..10)
	Number int

	// Shape is the 2D shape of the puzzle, first is the index X-direction (horizontal, to the right),
	// the second index is the Y-direction (up)
	Shape Array2d // -1=not part of volume, 0=empty, 1=occupied by a block

	// Height of the volume to fill with the blocks. This is always 2 for the original game
	Height int

	// The area of the problem in unit squares
	Area int

	// Bounding box of the problem volume
	BoundingBox BoundingBox

	// Blocks is an array of the blocks to be used to fill the volume
	Blocks []*Block
}

// Returns a string representation of the problem
func (p Problem) String() string {
	return fmt.Sprintf("Problem: card %s (%s-%s-%d) (%d blocks, area %d, height %d)",
		p.CardId, p.Animal, p.Difficulty, p.Number, len(p.Blocks), p.Area, p.Height)
}

// GetProblemArea calculates the area in unit squares of a given problem shape
func GetProblemArea(shape Array2d) int {
	var area int = 0
	for x, b := range shape {
		for y := range b {
			if shape[x][y] == 1 {
				area++
			}
		}
	}
	return area
}

// creates a problem instance
func MakeProblem(cardId string, number int, shape Array2d, blocks []*Block) *Problem {
	var p *Problem = new(Problem)

	p.CardId = cardId
	p.Number = number
	p.Shape = shape
	p.Blocks = blocks

	// the first character or the cardId defines the difficulty
	switch p.CardId[0] {
	case 'A':
		p.Difficulty = Easy
	case 'B':
		p.Difficulty = Difficult
	case 'C':
		p.Difficulty = Insane
	}

	// original game: height=2, for insane level, one more
	switch p.Difficulty {
	case Easy:
		p.Height = 2
	case Difficult:
		p.Height = 2
	case Insane:
		p.Height = 3
	}

	// the animal depends on the card number
	switch cardId[1:] {
	case "1", "2", "3", "4":
		p.Animal = "Elephant"
	case "5", "6", "7", "8":
		p.Animal = "Gazelle"
	case "9", "10", "11", "12":
		p.Animal = "Snake"
	case "13", "14", "15", "16":
		p.Animal = "Gnu"
	case "17", "18", "19", "20":
		p.Animal = "Ostrich"
	case "21", "22", "23", "24":
		p.Animal = "Rhino"
	case "25", "26", "27", "28":
		p.Animal = "Giraffe"
	case "29", "30", "31", "32":
		p.Animal = "Zebra"
	case "33", "34", "35", "36":
		p.Animal = "Warthog"
	}

	p.Area = GetProblemArea(p.Shape)
	p.BoundingBox = BoundingBox{len(p.Shape), len(p.Shape[0]), p.Height}

	return p
}

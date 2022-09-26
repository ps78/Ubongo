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

// Represents the singleton problem factory. Get instance with GetProblemFactory()
type ProblemFactory struct {
	Problems map[UbongoDifficulty](map[int](map[int]*Problem))
}

// Problem represents a single Ubongo problem to solve
type Problem struct {
	// CardId represents the number printed on each side of a card in the original game, without the letter
	CardNumber int

	// Animal is the name of the animal as printed on the original game
	Animal string

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
func NewProblem(cardNumber int, difficulty UbongoDifficulty, number int, shape *Array2d, blocks []*Block) *Problem {
	var p *Problem = new(Problem)

	p.CardNumber = cardNumber
	p.DiceNumber = number
	p.Shape = shape.Clone()
	p.Blocks = make([]*Block, len(blocks))
	copy(p.Blocks, blocks)

	p.Difficulty = difficulty

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
	switch cardNumber {
	case 1, 2, 3, 4:
		p.Animal = "Elephant"
	case 5, 6, 7, 8:
		p.Animal = "Gazelle"
	case 9, 10, 11, 12:
		p.Animal = "Snake"
	case 13, 14, 15, 16:
		p.Animal = "Gnu"
	case 17, 18, 19, 20:
		p.Animal = "Ostrich"
	case 21, 22, 23, 24:
		p.Animal = "Rhino"
	case 25, 26, 27, 28:
		p.Animal = "Giraffe"
	case 29, 30, 31, 32:
		p.Animal = "Zebra"
	case 33, 34, 35, 36:
		p.Animal = "Warthog"
	}

	p.Area = p.Shape.Count(0)
	p.BoundingBox = Vector{p.Shape.DimX, p.Shape.DimY, p.Height}

	return p
}

// Creates a deep copy of the given object p
func (p *Problem) Clone() *Problem {
	c := new(Problem)

	c.CardNumber = p.CardNumber
	c.DiceNumber = p.DiceNumber
	c.Difficulty = p.Difficulty
	c.BoundingBox = p.BoundingBox
	c.Animal = p.Animal
	c.Shape = p.Shape.Clone()
	c.Blocks = make([]*Block, len(p.Blocks))
	copy(c.Blocks, p.Blocks)
	c.Area = p.Area

	return c
}

// Creates all problem on one side of a card (as many as there are keys in the blocks-map)
func createProblems(cardNum int, difficulty UbongoDifficulty, topShape, bottomShape *Array2d, blocks map[int][]*Block, f *BlockFactory) []*Problem {

	result := make([]*Problem, 0)

	for k, v := range blocks {
		var shape *Array2d
		if (k <= 4 && difficulty == Easy) || (k <= 5) {
			shape = topShape
		} else {
			shape = bottomShape
		}

		p := NewProblem(cardNum, difficulty, k, shape, v)
		result = append(result, p)
	}

	return result
}

func createAllProblems(f *BlockFactory) []*Problem {
	problems := make([]*Problem, 0)

	// A1
	topShape := NewArray2dFromData([][]int8{{0, -1, -1}, {0, -1, -1}, {0, 0, -1}, {0, 0, 0}})
	bottomShape := NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums := map[int][]*Block{
		1: {f.Red_bighook, f.Green_L, f.Blue_lighter},
		3: {f.Blue_lighter, f.Green_bighook, f.Red_smallhook},
		5: {f.Yellow_bighook, f.Blue_bighook, f.Red_smallhook},
		8: {f.Yellow_smallhook, f.Red_stool, f.Green_flash}}
	problems = append(problems, createProblems(1, Easy, topShape, bottomShape, blockNums, f)...)

	// A2
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, -1, -1}, {0, 0, -1}, {0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Red_stool, f.Green_L, f.Yellow_hello},
		3: {f.Blue_lighter, f.Red_flash, f.Yellow_hello},
		5: {f.Green_L, f.Green_bighook, f.Red_bighook},
		8: {f.Blue_flash, f.Red_stool, f.Green_L}}
	problems = append(problems, createProblems(2, Easy, topShape, bottomShape, blockNums, f)...)

	// A3
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, -1, -1}, {0, 0, -1}, {0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Blue_lighter, f.Yellow_smallhook, f.Blue_bighook},
		3: {f.Red_bighook, f.Green_L, f.Red_stool},
		5: {f.Blue_lighter, f.Red_smallhook, f.Red_stool},
		8: {f.Blue_bighook, f.Green_L, f.Blue_lighter}}
	problems = append(problems, createProblems(3, Easy, topShape, bottomShape, blockNums, f)...)

	// A4
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, -1, -1}, {0, 0, -1}, {0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Green_L, f.Blue_lighter, f.Yellow_bighook},
		3: {f.Yellow_hello, f.Green_L, f.Blue_lighter},
		5: {f.Green_L, f.Red_bighook, f.Blue_bighook},
		8: {f.Yellow_bighook, f.Green_L, f.Blue_bighook}}
	problems = append(problems, createProblems(4, Easy, topShape, bottomShape, blockNums, f)...)

	// B1
	topShape = NewArray2dFromData([][]int8{{0, 0, 0}, {0, 0, -1}, {0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Red_smallhook, f.Blue_lighter, f.Green_L, f.Blue_v},
		2:  {f.Red_bighook, f.Red_smallhook, f.Blue_v, f.Red_flash},
		3:  {f.Blue_v, f.Blue_bighook, f.Green_T, f.Green_L},
		4:  {f.Green_flash, f.Green_T, f.Blue_v, f.Green_L},
		5:  {f.Red_flash, f.Blue_v, f.Green_L, f.Blue_bighook},
		6:  {f.Green_bighook, f.Yellow_smallhook, f.Green_L, f.Blue_v},
		7:  {f.Yellow_bighook, f.Red_smallhook, f.Blue_v, f.Green_T},
		8:  {f.Blue_flash, f.Green_L, f.Blue_v, f.Red_smallhook},
		9:  {f.Red_smallhook, f.Green_L, f.Red_bighook, f.Blue_v},
		10: {f.Red_bighook, f.Yellow_smallhook, f.Blue_v, f.Red_flash}}
	problems = append(problems, createProblems(1, Difficult, topShape, bottomShape, blockNums, f)...)

	// B2
	topShape = NewArray2dFromData([][]int8{{0, 0, -1, -1}, {0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, -1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Green_bighook, f.Yellow_smallhook, f.Green_flash, f.Red_smallhook},
		2:  {f.Red_stool, f.Blue_flash, f.Green_flash, f.Blue_v},
		3:  {f.Yellow_smallhook, f.Green_bighook, f.Green_flash, f.Red_flash},
		4:  {f.Yellow_smallhook, f.Green_bighook, f.Green_flash, f.Red_flash},
		5:  {f.Red_stool, f.Green_bighook, f.Green_L, f.Red_smallhook},
		6:  {f.Yellow_smallhook, f.Red_smallhook, f.Blue_v, f.Blue_bighook},
		7:  {f.Red_smallhook, f.Red_bighook, f.Green_L, f.Blue_v},
		8:  {f.Blue_v, f.Green_L, f.Blue_flash, f.Yellow_smallhook},
		9:  {f.Blue_v, f.Blue_bighook, f.Green_L, f.Red_flash},
		10: {f.Green_bighook, f.Red_smallhook, f.Blue_v, f.Yellow_smallhook}}
	problems = append(problems, createProblems(2, Difficult, topShape, bottomShape, blockNums, f)...)

	// B3
	topShape = NewArray2dFromData([][]int8{{-1, 0}, {0, 0}, {0, 0}, {0, 0}, {0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Red_flash, f.Blue_v, f.Green_L, f.Blue_flash},
		2:  {f.Green_T, f.Red_smallhook, f.Green_flash, f.Blue_v},
		3:  {f.Yellow_hello, f.Green_L, f.Green_T, f.Blue_v},
		4:  {f.Blue_v, f.Red_flash, f.Blue_bighook, f.Green_L},
		5:  {f.Red_smallhook, f.Yellow_hello, f.Yellow_smallhook, f.Blue_v},
		6:  {f.Green_L, f.Blue_v, f.Yellow_smallhook, f.Green_flash},
		7:  {f.Green_L, f.Red_flash, f.Green_bighook, f.Blue_v},
		8:  {f.Red_flash, f.Green_L, f.Blue_v, f.Yellow_hello},
		9:  {f.Green_flash, f.Red_smallhook, f.Red_flash, f.Blue_v},
		10: {f.Red_smallhook, f.Yellow_bighook, f.Blue_v, f.Red_flash}}
	problems = append(problems, createProblems(3, Difficult, topShape, bottomShape, blockNums, f)...)

	// B4
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {0, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Green_L, f.Blue_v, f.Green_bighook, f.Red_smallhook},
		2:  {f.Red_smallhook, f.Blue_bighook, f.Red_flash, f.Blue_v},
		3:  {f.Green_bighook, f.Blue_v, f.Green_L, f.Red_flash},
		4:  {f.Yellow_bighook, f.Red_smallhook, f.Blue_v, f.Green_T},
		5:  {f.Green_L, f.Blue_v, f.Yellow_smallhook, f.Green_bighook},
		6:  {f.Blue_v, f.Green_T, f.Green_L, f.Yellow_bighook},
		7:  {f.Red_smallhook, f.Yellow_smallhook, f.Blue_v, f.Blue_lighter},
		8:  {f.Green_L, f.Green_flash, f.Green_T, f.Blue_v},
		9:  {f.Red_smallhook, f.Yellow_smallhook, f.Blue_v, f.Yellow_hello},
		10: {f.Blue_v, f.Green_T, f.Blue_flash, f.Red_smallhook}}
	problems = append(problems, createProblems(4, Difficult, topShape, bottomShape, blockNums, f)...)

	// B12
	topShape = NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {0, 0, 0}, {0, 0, -1}, {-1, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Red_stool, f.Green_L, f.Blue_v, f.Red_flash},
		2:  {f.Red_stool, f.Yellow_smallhook, f.Green_T, f.Blue_v},
		3:  {f.Red_stool, f.Yellow_smallhook, f.Blue_v, f.Green_L},
		4:  {f.Red_bighook, f.Red_smallhook, f.Blue_v, f.Green_T},
		5:  {f.Green_L, f.Yellow_smallhook, f.Blue_flash, f.Blue_v},
		6:  {f.Yellow_smallhook, f.Green_L, f.Red_stool, f.Green_bighook},
		7:  {f.Blue_v, f.Yellow_hello, f.Yellow_gate, f.Blue_flash},
		8:  {f.Blue_flash, f.Green_L, f.Green_flash, f.Red_smallhook},
		9:  {f.Red_stool, f.Red_flash, f.Yellow_bighook, f.Red_smallhook},
		10: {f.Yellow_smallhook, f.Red_stool, f.Red_smallhook, f.Green_bighook}}
	problems = append(problems, createProblems(12, Difficult, topShape, bottomShape, blockNums, f)...)

	return problems
}

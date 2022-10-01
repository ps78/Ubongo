package main

import (
	"sync"
)

// Represents the singleton problem factory. Get instance with GetProblemFactory()
type ProblemFactory struct {
	// Contains all problems in map with 3 keys: [Difficulty][cardNumber][DiceNumber]
	Problems map[UbongoDifficulty](map[int](map[int]*Problem))
}

// used to create a thread-safe singleton instance of a problemFactory
var onceProblemFactorySingleton sync.Once

// the singleton
var problemFactoryInstance *ProblemFactory

// Returns the singleton instance of the problem factory
func GetProblemFactory() *ProblemFactory {
	// Create the singleton instance
	onceProblemFactorySingleton.Do(func() {
		bf := GetBlockFactory()

		f := new(ProblemFactory)

		f.Problems = make(map[UbongoDifficulty]map[int]map[int]*Problem)

		// create all standard easy and difficult problems in a flat slice
		allProblems := createAllEasyProblems(bf)
		allProblems = append(allProblems, createAllDifficultProblems(bf)...)

		// insert all problems in the 3-level map f.Problems[difficulty][cardNum][DiceNum]
		for _, p := range allProblems {
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

// Creates all easy problems of a specific card (as many as there are keys in the blocks-map)
func createEasyCardProblems(cardNum int, topShape, bottomShape *Array2d, blocks map[int][]*Block, f *BlockFactory) []*Problem {
	result := make([]*Problem, 0)

	for dice, blockset := range blocks {
		var shape *Array2d
		if dice <= 4 {
			shape = topShape
		} else {
			shape = bottomShape
		}
		result = append(result, NewProblem(cardNum, dice, Easy, 2, animalByCardNum[cardNum], shape, blockset))
	}

	return result
}

// Creates all difficult problems of a specific card (as many as there are keys in the blocks-map)
func createDifficultCardProblems(cardNum int, topShape, bottomShape *Array2d, blocks map[int][]*Block, f *BlockFactory) []*Problem {
	result := make([]*Problem, 0)

	for dice, blockset := range blocks {
		var shape *Array2d
		if dice <= 5 {
			shape = topShape
		} else {
			shape = bottomShape
		}
		result = append(result, NewProblem(cardNum, dice, Difficult, 2, animalByCardNum[cardNum], shape, blockset))
	}

	return result
}

// Creates all the problems from the original Ubongo game with the difficulty 'Easy'
// Returns a slice with 144 elements
func createAllEasyProblems(f *BlockFactory) []*Problem {
	problems := make([]*Problem, 0)

	// A1
	topShape := NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, -1, 0}, {-1, 0, 0}, {0, 0, 0}})
	bottomShape := NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, -1, -1}})
	blockNums := map[int][]*Block{
		1: {f.Red_bighook, f.Green_L, f.Blue_lighter},
		3: {f.Blue_lighter, f.Green_bighook, f.Red_smallhook},
		5: {f.Yellow_bighook, f.Blue_bighook, f.Red_smallhook},
		8: {f.Yellow_smallhook, f.Red_stool, f.Green_flash}}
	problems = append(problems, createEasyCardProblems(1, topShape, bottomShape, blockNums, f)...)

	// A2
	blockNums = map[int][]*Block{
		1: {f.Red_stool, f.Green_L, f.Yellow_hello},
		3: {f.Blue_lighter, f.Red_flash, f.Yellow_hello},
		5: {f.Green_L, f.Green_bighook, f.Red_bighook},
		8: {f.Blue_flash, f.Red_stool, f.Green_L}}
	problems = append(problems, createEasyCardProblems(2, topShape, bottomShape, blockNums, f)...)

	// A3
	blockNums = map[int][]*Block{
		1: {f.Blue_lighter, f.Yellow_smallhook, f.Blue_bighook},
		3: {f.Red_bighook, f.Green_L, f.Red_stool},
		5: {f.Blue_lighter, f.Red_smallhook, f.Red_stool},
		8: {f.Blue_bighook, f.Green_L, f.Blue_lighter}}
	problems = append(problems, createEasyCardProblems(3, topShape, bottomShape, blockNums, f)...)

	// A4
	blockNums = map[int][]*Block{
		1: {f.Green_L, f.Blue_lighter, f.Yellow_bighook},
		3: {f.Yellow_hello, f.Green_L, f.Blue_lighter},
		5: {f.Green_L, f.Red_bighook, f.Blue_bighook},
		8: {f.Yellow_bighook, f.Green_L, f.Blue_bighook}}
	problems = append(problems, createEasyCardProblems(4, topShape, bottomShape, blockNums, f)...)

	// A5
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, -1, 0}, {0, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {0, 0, -1}, {0, -1, -1}})
	blockNums = map[int][]*Block{
		1: {f.Yellow_bighook, f.Red_stool, f.Green_L},
		3: {f.Yellow_smallhook, f.Blue_bighook, f.Red_bighook},
		5: {f.Red_stool, f.Yellow_hello, f.Red_smallhook},
		8: {f.Yellow_gate, f.Green_L, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(5, topShape, bottomShape, blockNums, f)...)

	// A6
	blockNums = map[int][]*Block{
		1: {f.Blue_bighook, f.Red_stool, f.Green_L},
		3: {f.Yellow_bighook, f.Green_bighook, f.Red_smallhook},
		5: {f.Yellow_gate, f.Blue_bighook, f.Red_smallhook},
		8: {f.Blue_lighter, f.Green_T, f.Yellow_hello}}
	problems = append(problems, createEasyCardProblems(6, topShape, bottomShape, blockNums, f)...)

	// A7
	blockNums = map[int][]*Block{
		1: {f.Blue_flash, f.Green_L, f.Blue_lighter},
		3: {f.Green_T, f.Blue_lighter, f.Green_flash},
		5: {f.Yellow_hello, f.Blue_lighter, f.Yellow_smallhook},
		8: {f.Yellow_hello, f.Green_L, f.Blue_lighter}}
	problems = append(problems, createEasyCardProblems(7, topShape, bottomShape, blockNums, f)...)

	// A8
	blockNums = map[int][]*Block{
		1: {f.Blue_lighter, f.Red_smallhook, f.Blue_bighook},
		3: {f.Green_bighook, f.Blue_lighter, f.Yellow_smallhook},
		5: {f.Yellow_bighook, f.Green_L, f.Yellow_gate},
		8: {f.Yellow_gate, f.Green_bighook, f.Yellow_smallhook}}
	problems = append(problems, createEasyCardProblems(8, topShape, bottomShape, blockNums, f)...)

	// A9
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	blockNums = map[int][]*Block{
		1: {f.Red_stool, f.Red_bighook, f.Red_flash},
		3: {f.Green_L, f.Green_flash, f.Blue_lighter},
		5: {f.Green_T, f.Red_stool, f.Blue_flash},
		8: {f.Blue_bighook, f.Yellow_smallhook, f.Red_bighook}}
	problems = append(problems, createEasyCardProblems(9, topShape, bottomShape, blockNums, f)...)

	// A10
	blockNums = map[int][]*Block{
		1: {f.Red_smallhook, f.Green_bighook, f.Red_stool},
		3: {f.Red_stool, f.Yellow_smallhook, f.Blue_flash},
		5: {f.Yellow_smallhook, f.Blue_flash, f.Blue_bighook},
		8: {f.Red_smallhook, f.Yellow_bighook, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(10, topShape, bottomShape, blockNums, f)...)

	// A11
	blockNums = map[int][]*Block{
		1: {f.Blue_lighter, f.Green_L, f.Blue_flash},
		3: {f.Green_bighook, f.Green_flash, f.Yellow_smallhook},
		5: {f.Green_bighook, f.Red_smallhook, f.Red_stool},
		8: {f.Red_smallhook, f.Blue_bighook, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(11, topShape, bottomShape, blockNums, f)...)

	// A12
	blockNums = map[int][]*Block{
		1: {f.Yellow_smallhook, f.Red_stool, f.Yellow_hello},
		3: {f.Green_bighook, f.Red_smallhook, f.Blue_lighter},
		5: {f.Green_bighook, f.Green_flash, f.Red_smallhook},
		8: {f.Yellow_smallhook, f.Green_bighook, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(12, topShape, bottomShape, blockNums, f)...)

	// A13
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {0, 0, 0}})
	blockNums = map[int][]*Block{
		1: {f.Blue_bighook, f.Yellow_hello, f.Yellow_smallhook},
		3: {f.Blue_lighter, f.Red_smallhook, f.Blue_flash},
		5: {f.Blue_flash, f.Green_L, f.Red_stool},
		8: {f.Green_flash, f.Blue_lighter, f.Green_L}}
	problems = append(problems, createEasyCardProblems(13, topShape, bottomShape, blockNums, f)...)

	// A14
	blockNums = map[int][]*Block{
		1: {f.Yellow_gate, f.Red_smallhook, f.Red_bighook},
		3: {f.Red_smallhook, f.Blue_bighook, f.Red_stool},
		5: {f.Blue_lighter, f.Red_flash, f.Blue_flash},
		8: {f.Green_bighook, f.Green_T, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(14, topShape, bottomShape, blockNums, f)...)

	// A15
	blockNums = map[int][]*Block{
		1: {f.Yellow_smallhook, f.Red_stool, f.Green_bighook},
		3: {f.Green_flash, f.Red_smallhook, f.Blue_lighter},
		5: {f.Green_L, f.Red_stool, f.Green_flash},
		8: {f.Green_bighook, f.Green_L, f.Blue_flash}}
	problems = append(problems, createEasyCardProblems(15, topShape, bottomShape, blockNums, f)...)

	// A16
	blockNums = map[int][]*Block{
		1: {f.Yellow_smallhook, f.Green_flash, f.Blue_lighter},
		3: {f.Green_L, f.Red_stool, f.Green_flash},
		5: {f.Red_smallhook, f.Red_stool, f.Green_bighook},
		8: {f.Green_L, f.Green_flash, f.Blue_bighook}}
	problems = append(problems, createEasyCardProblems(16, topShape, bottomShape, blockNums, f)...)

	// A17
	topShape = NewArray2dFromData([][]int8{{-1, 0}, {0, 0}, {0, 0}, {0, -1}, {0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Yellow_gate, f.Blue_bighook, f.Green_L},
		3: {f.Yellow_hello, f.Green_L, f.Blue_lighter},
		5: {f.Yellow_smallhook, f.Red_stool, f.Yellow_hello},
		8: {f.Blue_lighter, f.Green_flash, f.Yellow_smallhook}}
	problems = append(problems, createEasyCardProblems(17, topShape, bottomShape, blockNums, f)...)

	// A18
	blockNums = map[int][]*Block{
		1: {f.Blue_lighter, f.Red_bighook, f.Green_L},
		3: {f.Yellow_hello, f.Red_flash, f.Blue_lighter},
		5: {f.Blue_lighter, f.Yellow_smallhook, f.Blue_flash},
		8: {f.Green_L, f.Red_bighook, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(18, topShape, bottomShape, blockNums, f)...)

	// A19
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, -1}, {-1, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Red_smallhook, f.Green_flash, f.Red_stool},
		3: {f.Blue_bighook, f.Green_L, f.Red_stool},
		5: {f.Red_stool, f.Green_L, f.Yellow_bighook},
		8: {f.Green_flash, f.Red_smallhook, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(19, topShape, bottomShape, blockNums, f)...)

	// A20
	blockNums = map[int][]*Block{
		1: {f.Red_stool, f.Yellow_hello, f.Red_smallhook},
		3: {f.Red_smallhook, f.Blue_bighook, f.Red_stool},
		5: {f.Red_flash, f.Red_stool, f.Yellow_bighook},
		8: {f.Red_smallhook, f.Blue_lighter, f.Blue_flash}}
	problems = append(problems, createEasyCardProblems(20, topShape, bottomShape, blockNums, f)...)

	// A21
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, 0}, {-1, 0, -1}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {-1, 0, 0}})
	blockNums = map[int][]*Block{
		1: {f.Blue_lighter, f.Yellow_bighook, f.Yellow_smallhook},
		3: {f.Blue_lighter, f.Red_stool, f.Green_L},
		5: {f.Green_bighook, f.Red_smallhook, f.Red_stool},
		8: {f.Red_smallhook, f.Blue_lighter, f.Yellow_hello}}
	problems = append(problems, createEasyCardProblems(21, topShape, bottomShape, blockNums, f)...)

	// A22
	blockNums = map[int][]*Block{
		1: {f.Blue_bighook, f.Red_stool, f.Green_L},
		3: {f.Yellow_bighook, f.Red_stool, f.Green_T},
		5: {f.Blue_flash, f.Red_smallhook, f.Blue_bighook},
		8: {f.Blue_bighook, f.Green_L, f.Yellow_hello}}
	problems = append(problems, createEasyCardProblems(22, topShape, bottomShape, blockNums, f)...)

	// A23
	blockNums = map[int][]*Block{
		1: {f.Blue_lighter, f.Red_bighook, f.Red_smallhook},
		3: {f.Blue_flash, f.Red_smallhook, f.Blue_lighter},
		5: {f.Yellow_smallhook, f.Red_stool, f.Blue_bighook},
		8: {f.Red_stool, f.Green_L, f.Green_bighook}}
	problems = append(problems, createEasyCardProblems(23, topShape, bottomShape, blockNums, f)...)

	// A24
	blockNums = map[int][]*Block{
		1: {f.Red_stool, f.Green_bighook, f.Green_L},
		3: {f.Yellow_hello, f.Green_L, f.Red_stool},
		5: {f.Yellow_smallhook, f.Yellow_bighook, f.Blue_lighter},
		8: {f.Green_L, f.Yellow_bighook, f.Green_flash}}
	problems = append(problems, createEasyCardProblems(24, topShape, bottomShape, blockNums, f)...)

	// A25
	topShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, -1}})
	blockNums = map[int][]*Block{
		1: {f.Red_smallhook, f.Yellow_bighook, f.Blue_v},
		3: {f.Blue_v, f.Red_smallhook, f.Yellow_hello},
		5: {f.Yellow_smallhook, f.Green_flash, f.Blue_lighter},
		8: {f.Red_smallhook, f.Green_bighook, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(25, topShape, bottomShape, blockNums, f)...)

	// A26
	blockNums = map[int][]*Block{
		1: {f.Blue_v, f.Blue_bighook, f.Yellow_smallhook},
		3: {f.Red_flash, f.Green_L, f.Red_smallhook},
		5: {f.Blue_flash, f.Red_stool, f.Green_L},
		8: {f.Green_flash, f.Green_bighook, f.Green_L}}
	problems = append(problems, createEasyCardProblems(26, topShape, bottomShape, blockNums, f)...)

	// A27
	blockNums = map[int][]*Block{
		1: {f.Blue_v, f.Green_bighook, f.Red_smallhook},
		3: {f.Blue_flash, f.Green_T, f.Blue_v},
		5: {f.Yellow_smallhook, f.Blue_lighter, f.Yellow_hello},
		8: {f.Green_flash, f.Red_stool, f.Green_L}}
	problems = append(problems, createEasyCardProblems(27, topShape, bottomShape, blockNums, f)...)

	// A28
	blockNums = map[int][]*Block{
		1: {f.Blue_v, f.Yellow_smallhook, f.Green_bighook},
		3: {f.Green_T, f.Green_flash, f.Blue_v},
		5: {f.Yellow_smallhook, f.Blue_bighook, f.Red_stool},
		8: {f.Blue_bighook, f.Green_L, f.Blue_flash}}
	problems = append(problems, createEasyCardProblems(28, topShape, bottomShape, blockNums, f)...)

	// A29
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Blue_bighook, f.Red_smallhook, f.Blue_v},
		3: {f.Blue_lighter, f.Yellow_smallhook, f.Blue_v},
		5: {f.Green_bighook, f.Red_stool, f.Green_L},
		8: {f.Green_bighook, f.Blue_lighter, f.Green_L}}
	problems = append(problems, createEasyCardProblems(29, topShape, bottomShape, blockNums, f)...)

	// A30
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {-1, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Green_flash, f.Red_stool, f.Red_flash},
		3: {f.Green_flash, f.Red_stool, f.Green_L},
		5: {f.Red_stool, f.Blue_bighook, f.Red_smallhook},
		8: {f.Blue_flash, f.Yellow_smallhook, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(30, topShape, bottomShape, blockNums, f)...)

	// A31
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Red_stool, f.Yellow_smallhook, f.Blue_v},
		3: {f.Green_bighook, f.Blue_v, f.Green_L},
		5: {f.Yellow_hello, f.Green_flash, f.Yellow_smallhook},
		8: {f.Red_smallhook, f.Blue_flash, f.Red_stool}}
	problems = append(problems, createEasyCardProblems(31, topShape, bottomShape, blockNums, f)...)

	// A32
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	blockNums = map[int][]*Block{
		1: {f.Red_stool, f.Green_flash, f.Red_flash},
		3: {f.Blue_bighook, f.Red_smallhook, f.Red_stool},
		5: {f.Green_flash, f.Red_stool, f.Green_L},
		8: {f.Green_L, f.Blue_lighter, f.Green_flash}}
	problems = append(problems, createEasyCardProblems(32, topShape, bottomShape, blockNums, f)...)

	// A33
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {0, 0, -1}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int][]*Block{
		1: {f.Blue_v, f.Green_bighook, f.Yellow_smallhook},
		3: {f.Blue_v, f.Green_L, f.Green_flash},
		5: {f.Green_bighook, f.Green_L, f.Red_stool},
		8: {f.Yellow_hello, f.Yellow_gate, f.Yellow_smallhook}}
	problems = append(problems, createEasyCardProblems(33, topShape, bottomShape, blockNums, f)...)

	// A34
	blockNums = map[int][]*Block{
		1: {f.Red_stool, f.Blue_v, f.Red_smallhook},
		3: {f.Yellow_hello, f.Red_flash, f.Blue_v},
		5: {f.Yellow_smallhook, f.Red_stool, f.Blue_bighook},
		8: {f.Yellow_hello, f.Blue_bighook, f.Green_L}}
	problems = append(problems, createEasyCardProblems(34, topShape, bottomShape, blockNums, f)...)

	// A35
	topShape = NewArray2dFromData([][]int8{{-1, 0}, {-1, 0}, {0, 0}, {0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int][]*Block{
		1: {f.Yellow_bighook, f.Blue_v, f.Green_L},
		3: {f.Red_flash, f.Green_L, f.Red_smallhook},
		5: {f.Red_smallhook, f.Yellow_bighook, f.Blue_lighter},
		8: {f.Blue_flash, f.Yellow_gate, f.Red_smallhook}}
	problems = append(problems, createEasyCardProblems(35, topShape, bottomShape, blockNums, f)...)

	// A36
	blockNums = map[int][]*Block{
		1: {f.Green_L, f.Yellow_hello, f.Blue_v},
		3: {f.Blue_bighook, f.Green_L, f.Blue_v},
		5: {f.Red_smallhook, f.Blue_flash, f.Yellow_hello},
		8: {f.Red_stool, f.Red_smallhook, f.Blue_bighook}}
	problems = append(problems, createEasyCardProblems(36, topShape, bottomShape, blockNums, f)...)

	return problems
}

// Creates all the problems from the original Ubongo game with the difficulty 'Easy'
// Returns a slice with 144 elements
func createAllDifficultProblems(f *BlockFactory) []*Problem {
	problems := make([]*Problem, 0)

	// B1
	topShape := NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	bottomShape := NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums := map[int][]*Block{
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
	problems = append(problems, createDifficultCardProblems(1, topShape, bottomShape, blockNums, f)...)

	// B2
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0, 0}, {-1, 0, 0, 0}, {-1, 0, 0, -1}, {0, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, -1, -1}})
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
	problems = append(problems, createDifficultCardProblems(2, topShape, bottomShape, blockNums, f)...)

	// B3
	topShape = NewArray2dFromData([][]int8{{0, -1}, {0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
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
	problems = append(problems, createDifficultCardProblems(3, topShape, bottomShape, blockNums, f)...)

	// B4
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, -1}, {0, 0, 0}})
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
	problems = append(problems, createDifficultCardProblems(4, topShape, bottomShape, blockNums, f)...)

	// B5
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0, 0}, {-1, 0, 0, -1}, {0, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Blue_v, f.Red_smallhook, f.Yellow_hello, f.Green_L},
		2:  {f.Blue_v, f.Red_smallhook, f.Yellow_hello, f.Red_flash},
		3:  {f.Green_T, f.Green_L, f.Green_flash, f.Blue_v},
		4:  {f.Yellow_bighook, f.Green_L, f.Red_flash, f.Blue_v},
		5:  {f.Blue_v, f.Red_smallhook, f.Blue_bighook, f.Green_L},
		6:  {f.Blue_bighook, f.Yellow_smallhook, f.Blue_v, f.Red_flash},
		7:  {f.Blue_v, f.Green_L, f.Red_smallhook, f.Green_bighook},
		8:  {f.Green_flash, f.Green_L, f.Blue_v, f.Yellow_smallhook},
		9:  {f.Green_L, f.Blue_v, f.Yellow_smallhook, f.Blue_lighter},
		10: {f.Blue_v, f.Green_bighook, f.Red_smallhook, f.Green_T}}
	problems = append(problems, createDifficultCardProblems(5, topShape, bottomShape, blockNums, f)...)

	// B6
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Blue_v, f.Red_smallhook, f.Red_stool, f.Red_flash},
		2:  {f.Red_flash, f.Green_bighook, f.Yellow_smallhook, f.Blue_v},
		3:  {f.Yellow_smallhook, f.Yellow_hello, f.Red_smallhook, f.Blue_v},
		4:  {f.Blue_v, f.Green_L, f.Blue_flash, f.Red_smallhook},
		5:  {f.Blue_v, f.Green_L, f.Blue_flash, f.Yellow_smallhook},
		6:  {f.Red_smallhook, f.Blue_v, f.Red_flash, f.Yellow_hello},
		7:  {f.Blue_v, f.Green_T, f.Blue_lighter, f.Green_L},
		8:  {f.Green_L, f.Green_flash, f.Green_T, f.Blue_v},
		9:  {f.Green_bighook, f.Red_smallhook, f.Yellow_smallhook, f.Blue_v},
		10: {f.Red_flash, f.Blue_v, f.Green_L, f.Green_bighook}}
	problems = append(problems, createDifficultCardProblems(6, topShape, bottomShape, blockNums, f)...)

	// B7
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1, -1}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Red_smallhook, f.Green_T, f.Blue_v, f.Green_flash},
		2:  {f.Blue_v, f.Red_smallhook, f.Blue_bighook, f.Green_L},
		3:  {f.Green_L, f.Blue_v, f.Yellow_bighook, f.Red_smallhook},
		4:  {f.Green_L, f.Blue_v, f.Yellow_smallhook, f.Blue_flash},
		5:  {f.Yellow_bighook, f.Yellow_smallhook, f.Red_smallhook, f.Blue_v},
		6:  {f.Yellow_smallhook, f.Blue_flash, f.Blue_v, f.Green_T},
		7:  {f.Yellow_bighook, f.Blue_v, f.Red_flash, f.Yellow_smallhook},
		8:  {f.Green_bighook, f.Blue_v, f.Yellow_smallhook, f.Green_L},
		9:  {f.Red_smallhook, f.Blue_v, f.Green_flash, f.Green_L},
		10: {f.Blue_v, f.Green_T, f.Red_smallhook, f.Blue_bighook}}
	problems = append(problems, createDifficultCardProblems(7, topShape, bottomShape, blockNums, f)...)

	// B8
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1, -1}, {-1, 0, 0, 0}, {-1, 0, 0, -1}, {0, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {0, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Red_stool, f.Green_L, f.Yellow_smallhook, f.Blue_v},
		2:  {f.Green_L, f.Yellow_smallhook, f.Blue_v, f.Blue_flash},
		3:  {f.Red_smallhook, f.Red_flash, f.Blue_v, f.Green_bighook},
		4:  {f.Red_smallhook, f.Red_stool, f.Blue_v, f.Red_flash},
		5:  {f.Green_bighook, f.Blue_v, f.Green_L, f.Red_smallhook},
		6:  {f.Blue_v, f.Green_T, f.Red_smallhook, f.Yellow_bighook},
		7:  {f.Red_bighook, f.Yellow_smallhook, f.Green_L, f.Blue_v},
		8:  {f.Green_L, f.Red_smallhook, f.Yellow_hello, f.Blue_v},
		9:  {f.Green_L, f.Red_smallhook, f.Yellow_bighook, f.Blue_v},
		10: {f.Blue_bighook, f.Green_L, f.Blue_v, f.Yellow_smallhook}}
	problems = append(problems, createDifficultCardProblems(8, topShape, bottomShape, blockNums, f)...)

	// B9
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {0, 0, 0}, {0, 0, 0}, {0, -1, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Green_L, f.Blue_v, f.Green_T, f.Yellow_bighook},
		2:  {f.Green_L, f.Blue_v, f.Red_smallhook, f.Blue_lighter},
		3:  {f.Red_smallhook, f.Yellow_smallhook, f.Blue_lighter, f.Blue_v},
		4:  {f.Yellow_bighook, f.Blue_v, f.Green_T, f.Red_flash},
		5:  {f.Blue_lighter, f.Blue_v, f.Green_L, f.Yellow_smallhook},
		6:  {f.Blue_lighter, f.Blue_v, f.Green_flash, f.Blue_bighook},
		7:  {f.Red_bighook, f.Red_stool, f.Green_L, f.Yellow_smallhook},
		8:  {f.Red_bighook, f.Blue_v, f.Green_bighook, f.Blue_lighter},
		9:  {f.Blue_flash, f.Red_smallhook, f.Yellow_smallhook, f.Blue_lighter},
		10: {f.Green_L, f.Blue_lighter, f.Blue_bighook, f.Yellow_smallhook}}
	problems = append(problems, createDifficultCardProblems(9, topShape, bottomShape, blockNums, f)...)

	// B10
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {-1, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Yellow_smallhook, f.Blue_lighter, f.Blue_v, f.Green_L},
		2:  {f.Red_flash, f.Blue_v, f.Red_smallhook, f.Blue_lighter},
		3:  {f.Blue_v, f.Red_flash, f.Green_flash, f.Green_T},
		4:  {f.Red_smallhook, f.Red_flash, f.Green_flash, f.Blue_v},
		5:  {f.Blue_v, f.Yellow_smallhook, f.Green_L, f.Green_flash},
		6:  {f.Green_L, f.Yellow_smallhook, f.Blue_flash, f.Yellow_bighook},
		7:  {f.Red_stool, f.Blue_lighter, f.Green_flash, f.Blue_v},
		8:  {f.Blue_flash, f.Blue_bighook, f.Green_L, f.Yellow_smallhook},
		9:  {f.Red_flash, f.Green_L, f.Yellow_bighook, f.Blue_bighook},
		10: {f.Blue_v, f.Green_flash, f.Red_bighook, f.Blue_lighter}}
	problems = append(problems, createDifficultCardProblems(10, topShape, bottomShape, blockNums, f)...)

	// B11
	topShape = NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {0, 0, -1}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Green_L, f.Red_smallhook, f.Blue_v, f.Red_stool},
		2:  {f.Blue_v, f.Green_L, f.Red_flash, f.Blue_bighook},
		3:  {f.Green_L, f.Red_smallhook, f.Blue_v, f.Blue_lighter},
		4:  {f.Yellow_smallhook, f.Green_bighook, f.Blue_v, f.Green_L},
		5:  {f.Yellow_gate, f.Green_L, f.Blue_v, f.Red_smallhook},
		6:  {f.Yellow_hello, f.Red_flash, f.Blue_lighter, f.Red_smallhook},
		7:  {f.Yellow_hello, f.Yellow_gate, f.Red_flash, f.Red_smallhook},
		8:  {f.Blue_bighook, f.Red_smallhook, f.Yellow_smallhook, f.Blue_lighter},
		9:  {f.Green_L, f.Green_T, f.Yellow_gate, f.Blue_bighook},
		10: {f.Red_smallhook, f.Green_L, f.Red_bighook, f.Green_flash}}
	problems = append(problems, createDifficultCardProblems(11, topShape, bottomShape, blockNums, f)...)

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
	problems = append(problems, createDifficultCardProblems(12, topShape, bottomShape, blockNums, f)...)

	// 13
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {0, 0, 0}, {0, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Yellow_bighook, f.Red_smallhook, f.Blue_bighook, f.Yellow_smallhook},
		2:  {f.Yellow_smallhook, f.Red_bighook, f.Red_flash, f.Yellow_gate},
		3:  {f.Green_L, f.Green_flash, f.Yellow_smallhook, f.Blue_bighook},
		4:  {f.Yellow_hello, f.Yellow_gate, f.Yellow_bighook, f.Blue_v},
		5:  {f.Blue_v, f.Red_stool, f.Red_bighook, f.Yellow_hello},
		6:  {f.Green_L, f.Blue_bighook, f.Red_bighook, f.Yellow_smallhook},
		7:  {f.Red_stool, f.Red_smallhook, f.Green_T, f.Blue_bighook},
		8:  {f.Red_smallhook, f.Blue_lighter, f.Blue_flash, f.Green_L},
		9:  {f.Blue_v, f.Blue_bighook, f.Yellow_gate, f.Red_stool},
		10: {f.Yellow_hello, f.Red_flash, f.Red_smallhook, f.Red_stool}}
	problems = append(problems, createDifficultCardProblems(13, topShape, bottomShape, blockNums, f)...)

	// B14
	topShape = NewArray2dFromData([][]int8{{0, -1, -1, -1}, {0, 0, -1, -1}, {0, 0, 0, -1}, {-1, 0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Blue_lighter, f.Yellow_smallhook, f.Yellow_hello, f.Green_L},
		2:  {f.Yellow_gate, f.Red_smallhook, f.Red_flash, f.Yellow_bighook},
		3:  {f.Green_bighook, f.Blue_lighter, f.Green_T, f.Red_smallhook},
		4:  {f.Red_flash, f.Green_L, f.Yellow_hello, f.Blue_lighter},
		5:  {f.Green_flash, f.Green_L, f.Yellow_smallhook, f.Blue_flash},
		6:  {f.Green_L, f.Green_bighook, f.Yellow_bighook, f.Red_smallhook},
		7:  {f.Yellow_smallhook, f.Red_smallhook, f.Green_flash, f.Red_stool},
		8:  {f.Blue_bighook, f.Red_smallhook, f.Blue_lighter, f.Yellow_smallhook},
		9:  {f.Yellow_gate, f.Red_stool, f.Green_T, f.Yellow_smallhook},
		10: {f.Green_flash, f.Red_smallhook, f.Red_flash, f.Yellow_hello}}
	problems = append(problems, createDifficultCardProblems(14, topShape, bottomShape, blockNums, f)...)

	// B15
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {0, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Yellow_hello, f.Green_bighook, f.Blue_lighter, f.Blue_v},
		2:  {f.Yellow_smallhook, f.Blue_lighter, f.Red_smallhook, f.Yellow_hello},
		3:  {f.Green_L, f.Red_smallhook, f.Yellow_gate, f.Blue_flash},
		4:  {f.Blue_lighter, f.Green_L, f.Green_T, f.Blue_bighook},
		5:  {f.Red_smallhook, f.Green_bighook, f.Yellow_smallhook, f.Red_stool},
		6:  {f.Green_bighook, f.Red_smallhook, f.Red_stool, f.Green_L},
		7:  {f.Green_L, f.Red_stool, f.Blue_lighter, f.Red_smallhook},
		8:  {f.Blue_flash, f.Red_flash, f.Green_L, f.Red_stool},
		9:  {f.Red_bighook, f.Green_bighook, f.Green_L, f.Red_smallhook},
		10: {f.Green_L, f.Green_flash, f.Green_bighook, f.Red_smallhook}}
	problems = append(problems, createDifficultCardProblems(15, topShape, bottomShape, blockNums, f)...)

	// B16
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {0, 0, -1}, {0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {0, 0, -1}, {0, 0, -1}, {0, -1, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Green_L, f.Blue_bighook, f.Red_stool, f.Yellow_smallhook},
		2:  {f.Red_smallhook, f.Green_L, f.Blue_lighter, f.Red_stool},
		3:  {f.Yellow_hello, f.Red_stool, f.Blue_v, f.Green_bighook},
		4:  {f.Blue_lighter, f.Red_smallhook, f.Green_L, f.Blue_bighook},
		5:  {f.Yellow_smallhook, f.Red_stool, f.Red_flash, f.Yellow_hello},
		6:  {f.Green_flash, f.Yellow_smallhook, f.Green_L, f.Red_bighook},
		7:  {f.Green_flash, f.Green_T, f.Blue_lighter, f.Green_L},
		8:  {f.Green_bighook, f.Green_L, f.Yellow_smallhook, f.Green_flash},
		9:  {f.Yellow_bighook, f.Blue_lighter, f.Red_smallhook, f.Green_L},
		10: {f.Red_stool, f.Blue_v, f.Green_bighook, f.Blue_bighook}}
	problems = append(problems, createDifficultCardProblems(16, topShape, bottomShape, blockNums, f)...)

	// B17
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Blue_lighter, f.Blue_bighook, f.Green_bighook, f.Blue_v},
		2:  {f.Red_stool, f.Green_L, f.Green_bighook, f.Yellow_smallhook},
		3:  {f.Green_T, f.Green_L, f.Red_stool, f.Green_bighook},
		4:  {f.Yellow_gate, f.Green_flash, f.Blue_v, f.Green_bighook},
		5:  {f.Yellow_smallhook, f.Green_L, f.Green_flash, f.Red_bighook},
		6:  {f.Yellow_smallhook, f.Yellow_hello, f.Blue_lighter, f.Green_L},
		7:  {f.Green_flash, f.Yellow_smallhook, f.Blue_bighook, f.Green_L},
		8:  {f.Blue_lighter, f.Yellow_smallhook, f.Yellow_bighook, f.Green_T},
		9:  {f.Blue_flash, f.Red_smallhook, f.Red_bighook, f.Green_L},
		10: {f.Green_L, f.Blue_lighter, f.Yellow_smallhook, f.Green_flash}}
	problems = append(problems, createDifficultCardProblems(17, topShape, bottomShape, blockNums, f)...)

	// B18
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0, 0}, {-1, 0, 0, -1}, {0, 0, 0, 0}, {-1, -1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, 0, 0, -1}, {-1, -1, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Red_stool, f.Red_smallhook, f.Yellow_smallhook, f.Blue_flash},
		2:  {f.Yellow_smallhook, f.Blue_bighook, f.Blue_flash, f.Red_smallhook},
		3:  {f.Yellow_hello, f.Red_bighook, f.Red_smallhook, f.Yellow_smallhook},
		4:  {f.Red_stool, f.Red_bighook, f.Blue_v, f.Green_bighook},
		5:  {f.Blue_flash, f.Red_stool, f.Blue_bighook, f.Blue_v},
		6:  {f.Yellow_bighook, f.Red_smallhook, f.Green_L, f.Blue_flash},
		7:  {f.Red_smallhook, f.Green_L, f.Yellow_bighook, f.Red_stool},
		8:  {f.Blue_flash, f.Yellow_smallhook, f.Blue_lighter, f.Green_L},
		9:  {f.Green_L, f.Blue_bighook, f.Yellow_smallhook, f.Green_flash},
		10: {f.Green_L, f.Green_flash, f.Red_flash, f.Yellow_bighook}}
	problems = append(problems, createDifficultCardProblems(18, topShape, bottomShape, blockNums, f)...)

	// B19
	topShape = NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {0, 0, -1}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Red_stool, f.Green_bighook, f.Yellow_smallhook, f.Red_smallhook},
		2:  {f.Blue_v, f.Green_flash, f.Red_stool, f.Blue_bighook},
		3:  {f.Yellow_hello, f.Red_stool, f.Green_bighook, f.Blue_v},
		4:  {f.Red_stool, f.Blue_bighook, f.Green_L, f.Red_smallhook},
		5:  {f.Red_smallhook, f.Blue_flash, f.Red_stool, f.Red_flash},
		6:  {f.Green_L, f.Blue_bighook, f.Green_flash, f.Green_T},
		7:  {f.Green_L, f.Yellow_hello, f.Blue_lighter, f.Red_smallhook},
		8:  {f.Blue_lighter, f.Green_bighook, f.Red_smallhook, f.Yellow_smallhook},
		9:  {f.Blue_bighook, f.Blue_lighter, f.Green_L, f.Red_smallhook},
		10: {f.Blue_bighook, f.Red_stool, f.Green_bighook, f.Blue_v}}
	problems = append(problems, createDifficultCardProblems(19, topShape, bottomShape, blockNums, f)...)

	// B20
	topShape = NewArray2dFromData([][]int8{{0, 0, 0, 0}, {-1, 0, 0, -1}, {-1, -1, 0, 0}, {-1, -1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, -1}, {0, 0, -1}, {-1, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Green_L, f.Red_smallhook, f.Red_stool, f.Green_flash},
		2:  {f.Yellow_bighook, f.Blue_flash, f.Red_smallhook, f.Yellow_smallhook},
		3:  {f.Green_L, f.Yellow_smallhook, f.Blue_flash, f.Red_stool},
		4:  {f.Red_smallhook, f.Green_L, f.Blue_lighter, f.Red_stool},
		5:  {f.Blue_bighook, f.Green_flash, f.Blue_v, f.Red_stool},
		6:  {f.Green_bighook, f.Blue_flash, f.Yellow_smallhook, f.Green_L},
		7:  {f.Green_L, f.Red_stool, f.Red_flash, f.Blue_bighook},
		8:  {f.Green_T, f.Green_bighook, f.Blue_flash, f.Red_smallhook},
		9:  {f.Green_bighook, f.Yellow_smallhook, f.Green_L, f.Red_stool},
		10: {f.Green_L, f.Blue_bighook, f.Blue_lighter, f.Yellow_smallhook}}
	problems = append(problems, createDifficultCardProblems(20, topShape, bottomShape, blockNums, f)...)

	// B21
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, -1}, {-1, 0, 0}, {0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0, -1}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, 0, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Blue_flash, f.Green_bighook, f.Red_smallhook, f.Green_L},
		2:  {f.Red_stool, f.Red_smallhook, f.Red_flash, f.Green_bighook},
		3:  {f.Red_smallhook, f.Green_L, f.Green_flash, f.Red_stool},
		4:  {f.Blue_lighter, f.Red_bighook, f.Green_L, f.Red_smallhook},
		5:  {f.Green_L, f.Red_stool, f.Red_flash, f.Blue_flash},
		6:  {f.Red_smallhook, f.Blue_lighter, f.Green_L, f.Yellow_hello},
		7:  {f.Yellow_bighook, f.Green_L, f.Blue_flash, f.Red_smallhook},
		8:  {f.Green_L, f.Yellow_hello, f.Green_T, f.Blue_flash},
		9:  {f.Red_stool, f.Red_smallhook, f.Yellow_hello, f.Yellow_smallhook},
		10: {f.Red_bighook, f.Yellow_hello, f.Blue_v, f.Red_stool}}
	problems = append(problems, createDifficultCardProblems(21, topShape, bottomShape, blockNums, f)...)

	// B22
	topShape = NewArray2dFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, -1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Red_stool, f.Blue_bighook, f.Blue_v, f.Blue_flash},
		2:  {f.Yellow_smallhook, f.Red_stool, f.Green_L, f.Red_bighook},
		3:  {f.Green_L, f.Blue_flash, f.Yellow_smallhook, f.Green_bighook},
		4:  {f.Yellow_smallhook, f.Green_L, f.Red_bighook, f.Blue_flash},
		5:  {f.Yellow_bighook, f.Blue_v, f.Blue_lighter, f.Green_bighook},
		6:  {f.Yellow_bighook, f.Blue_v, f.Blue_lighter, f.Red_stool},
		7:  {f.Blue_bighook, f.Blue_lighter, f.Red_bighook, f.Blue_v},
		8:  {f.Green_T, f.Red_stool, f.Yellow_smallhook, f.Blue_bighook},
		9:  {f.Green_bighook, f.Blue_v, f.Red_stool, f.Yellow_gate},
		10: {f.Red_stool, f.Green_flash, f.Yellow_smallhook, f.Green_T}}
	problems = append(problems, createDifficultCardProblems(22, topShape, bottomShape, blockNums, f)...)

	// B23
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, -1, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Red_stool, f.Green_L, f.Green_bighook, f.Red_flash},
		2:  {f.Red_stool, f.Yellow_gate, f.Yellow_smallhook, f.Green_L},
		3:  {f.Blue_v, f.Blue_lighter, f.Blue_bighook, f.Red_stool},
		4:  {f.Blue_v, f.Red_stool, f.Yellow_hello, f.Yellow_bighook},
		5:  {f.Red_smallhook, f.Yellow_smallhook, f.Green_bighook, f.Blue_flash},
		6:  {f.Blue_flash, f.Yellow_gate, f.Blue_v, f.Green_bighook},
		7:  {f.Green_L, f.Yellow_smallhook, f.Yellow_gate, f.Blue_lighter},
		8:  {f.Blue_v, f.Yellow_hello, f.Blue_bighook, f.Green_bighook},
		9:  {f.Blue_lighter, f.Yellow_smallhook, f.Red_flash, f.Yellow_hello},
		10: {f.Yellow_hello, f.Green_flash, f.Green_L, f.Yellow_smallhook}}
	problems = append(problems, createDifficultCardProblems(23, topShape, bottomShape, blockNums, f)...)

	// B24
	topShape = NewArray2dFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0, 0}, {-1, 0, 0, -1}, {-1, 0, 0, -1}, {-1, -1, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Blue_bighook, f.Green_L, f.Red_stool, f.Red_smallhook},
		2:  {f.Yellow_smallhook, f.Blue_flash, f.Green_L, f.Red_bighook},
		3:  {f.Yellow_hello, f.Red_bighook, f.Blue_bighook, f.Blue_v},
		4:  {f.Green_L, f.Yellow_smallhook, f.Red_stool, f.Green_bighook},
		5:  {f.Yellow_smallhook, f.Green_L, f.Green_flash, f.Yellow_bighook},
		6:  {f.Green_L, f.Red_smallhook, f.Yellow_hello, f.Red_bighook},
		7:  {f.Blue_lighter, f.Green_flash, f.Blue_v, f.Red_stool},
		8:  {f.Blue_lighter, f.Red_stool, f.Green_L, f.Red_smallhook},
		9:  {f.Green_L, f.Red_smallhook, f.Yellow_bighook, f.Blue_bighook},
		10: {f.Green_bighook, f.Red_smallhook, f.Red_bighook, f.Green_L}}
	problems = append(problems, createDifficultCardProblems(24, topShape, bottomShape, blockNums, f)...)

	// B25
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Blue_v, f.Red_bighook, f.Red_stool, f.Blue_lighter},
		2:  {f.Blue_bighook, f.Red_stool, f.Green_T, f.Yellow_smallhook},
		3:  {f.Green_T, f.Yellow_smallhook, f.Blue_lighter, f.Green_bighook},
		4:  {f.Yellow_smallhook, f.Blue_flash, f.Green_L, f.Yellow_gate},
		5:  {f.Green_bighook, f.Red_smallhook, f.Blue_lighter, f.Green_L},
		6:  {f.Yellow_smallhook, f.Blue_lighter, f.Green_L, f.Blue_flash},
		7:  {f.Blue_bighook, f.Yellow_smallhook, f.Blue_lighter, f.Green_L},
		8:  {f.Red_smallhook, f.Blue_lighter, f.Yellow_hello, f.Green_T},
		9:  {f.Yellow_hello, f.Yellow_smallhook, f.Green_L, f.Red_bighook},
		10: {f.Blue_bighook, f.Green_L, f.Yellow_hello, f.Red_smallhook}}
	problems = append(problems, createDifficultCardProblems(25, topShape, bottomShape, blockNums, f)...)

	// B26
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}, {-1, -1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {-1, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Yellow_smallhook, f.Yellow_hello, f.Green_L, f.Yellow_bighook},
		2:  {f.Yellow_bighook, f.Blue_v, f.Red_stool, f.Yellow_hello},
		3:  {f.Green_L, f.Red_stool, f.Blue_bighook, f.Red_smallhook},
		4:  {f.Green_bighook, f.Red_smallhook, f.Green_flash, f.Green_L},
		5:  {f.Blue_bighook, f.Green_bighook, f.Red_smallhook, f.Yellow_smallhook},
		6:  {f.Red_stool, f.Green_L, f.Green_flash, f.Yellow_smallhook},
		7:  {f.Red_bighook, f.Yellow_hello, f.Blue_v, f.Green_flash},
		8:  {f.Yellow_hello, f.Green_L, f.Red_flash, f.Green_flash},
		9:  {f.Green_T, f.Yellow_bighook, f.Red_stool, f.Green_L},
		10: {f.Green_flash, f.Yellow_hello, f.Yellow_smallhook, f.Red_smallhook}}
	problems = append(problems, createDifficultCardProblems(26, topShape, bottomShape, blockNums, f)...)

	// B27
	topShape = NewArray2dFromData([][]int8{{0, 0, 0, -1}, {-1, 0, 0, 0}, {-1, -1, 0, 0}, {-1, -1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, 0, 0, 0}, {-1, -1, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Yellow_smallhook, f.Green_flash, f.Red_stool, f.Green_T},
		2:  {f.Blue_lighter, f.Green_L, f.Blue_bighook, f.Yellow_smallhook},
		3:  {f.Green_L, f.Green_flash, f.Red_smallhook, f.Blue_lighter},
		4:  {f.Yellow_gate, f.Blue_lighter, f.Green_bighook, f.Blue_v},
		5:  {f.Yellow_hello, f.Red_smallhook, f.Blue_lighter, f.Red_flash},
		6:  {f.Red_stool, f.Blue_v, f.Blue_bighook, f.Yellow_hello},
		7:  {f.Red_smallhook, f.Green_bighook, f.Green_L, f.Yellow_gate},
		8:  {f.Yellow_bighook, f.Green_L, f.Green_flash, f.Yellow_smallhook},
		9:  {f.Red_smallhook, f.Yellow_gate, f.Yellow_smallhook, f.Green_bighook},
		10: {f.Blue_v, f.Blue_bighook, f.Blue_lighter, f.Yellow_bighook}}
	problems = append(problems, createDifficultCardProblems(27, topShape, bottomShape, blockNums, f)...)

	// B28
	topShape = NewArray2dFromData([][]int8{{0, 0, -1, -1}, {0, 0, 0, 0}, {-1, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {0, 0, -1}, {0, 0, -1}, {-1, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Green_T, f.Red_stool, f.Red_smallhook, f.Green_flash},
		2:  {f.Blue_v, f.Yellow_hello, f.Red_stool, f.Blue_lighter},
		3:  {f.Blue_flash, f.Red_stool, f.Green_flash, f.Blue_v},
		4:  {f.Yellow_smallhook, f.Red_smallhook, f.Red_stool, f.Yellow_hello},
		5:  {f.Green_T, f.Red_stool, f.Blue_flash, f.Yellow_smallhook},
		6:  {f.Blue_flash, f.Green_L, f.Yellow_smallhook, f.Yellow_gate},
		7:  {f.Green_L, f.Yellow_smallhook, f.Yellow_bighook, f.Blue_lighter},
		8:  {f.Green_L, f.Yellow_bighook, f.Red_smallhook, f.Blue_bighook},
		9:  {f.Blue_lighter, f.Blue_v, f.Red_stool, f.Blue_bighook},
		10: {f.Green_bighook, f.Blue_lighter, f.Green_flash, f.Blue_v}}
	problems = append(problems, createDifficultCardProblems(28, topShape, bottomShape, blockNums, f)...)

	// B29
	topShape = NewArray2dFromData([][]int8{{-1, 0, 0, -1}, {0, 0, 0, 0}, {-1, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Yellow_smallhook, f.Blue_lighter, f.Blue_bighook, f.Red_flash},
		2:  {f.Red_stool, f.Blue_flash, f.Green_T, f.Yellow_smallhook},
		3:  {f.Green_T, f.Yellow_bighook, f.Yellow_smallhook, f.Green_bighook},
		4:  {f.Blue_v, f.Green_flash, f.Red_stool, f.Yellow_hello},
		5:  {f.Green_flash, f.Red_smallhook, f.Red_bighook, f.Yellow_smallhook},
		6:  {f.Green_L, f.Blue_bighook, f.Red_smallhook, f.Blue_lighter},
		7:  {f.Red_stool, f.Green_L, f.Blue_bighook, f.Red_smallhook},
		8:  {f.Blue_lighter, f.Blue_bighook, f.Green_T, f.Yellow_smallhook},
		9:  {f.Green_bighook, f.Blue_lighter, f.Blue_v, f.Yellow_gate},
		10: {f.Green_flash, f.Red_bighook, f.Green_L, f.Red_smallhook}}
	problems = append(problems, createDifficultCardProblems(29, topShape, bottomShape, blockNums, f)...)

	// B30
	topShape = NewArray2dFromData([][]int8{{-1, 0, 0, -1}, {0, 0, 0, -1}, {-1, -1, 0, 0}, {-1, -1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, 0, 0, 0}, {-1, 0, -1, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Yellow_smallhook, f.Green_L, f.Red_stool, f.Blue_bighook},
		2:  {f.Blue_bighook, f.Red_stool, f.Red_smallhook, f.Green_L},
		3:  {f.Red_stool, f.Blue_v, f.Blue_bighook, f.Green_bighook},
		4:  {f.Red_smallhook, f.Green_L, f.Red_stool, f.Green_bighook},
		5:  {f.Red_stool, f.Red_bighook, f.Yellow_smallhook, f.Green_L},
		6:  {f.Blue_v, f.Red_stool, f.Red_bighook, f.Blue_lighter},
		7:  {f.Green_L, f.Red_smallhook, f.Red_stool, f.Blue_lighter},
		8:  {f.Green_flash, f.Yellow_bighook, f.Blue_v, f.Green_bighook},
		9:  {f.Blue_flash, f.Blue_v, f.Blue_lighter, f.Blue_bighook},
		10: {f.Yellow_smallhook, f.Blue_bighook, f.Green_L, f.Yellow_hello}}
	problems = append(problems, createDifficultCardProblems(30, topShape, bottomShape, blockNums, f)...)

	// B31
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, -1}, {-1, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0, 0}, {0, 0, 0, -1}, {0, 0, 0, -1}, {0, -1, -1, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Green_bighook, f.Blue_lighter, f.Red_stool, f.Blue_v},
		2:  {f.Red_stool, f.Green_L, f.Yellow_hello, f.Red_flash},
		3:  {f.Green_flash, f.Yellow_smallhook, f.Red_stool, f.Red_smallhook},
		4:  {f.Green_bighook, f.Red_smallhook, f.Red_stool, f.Green_L},
		5:  {f.Yellow_smallhook, f.Green_L, f.Yellow_hello, f.Red_stool},
		6:  {f.Blue_bighook, f.Green_flash, f.Green_L, f.Yellow_smallhook},
		7:  {f.Red_flash, f.Red_smallhook, f.Blue_lighter, f.Yellow_bighook},
		8:  {f.Green_L, f.Red_flash, f.Red_stool, f.Yellow_hello},
		9:  {f.Green_L, f.Yellow_smallhook, f.Red_bighook, f.Green_flash},
		10: {f.Green_L, f.Blue_flash, f.Yellow_smallhook, f.Green_bighook}}
	problems = append(problems, createDifficultCardProblems(31, topShape, bottomShape, blockNums, f)...)

	// B32
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {0, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0, -1}, {0, 0, 0, -1}, {-1, 0, 0, 0}, {-1, 0, -1, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Green_L, f.Yellow_smallhook, f.Blue_v, f.Green_bighook},
		2:  {f.Blue_v, f.Green_T, f.Blue_bighook, f.Green_L},
		3:  {f.Blue_v, f.Red_bighook, f.Yellow_smallhook, f.Red_flash},
		4:  {f.Green_L, f.Red_flash, f.Yellow_hello, f.Blue_v},
		5:  {f.Blue_bighook, f.Blue_v, f.Red_smallhook, f.Green_L},
		6:  {f.Red_flash, f.Red_smallhook, f.Red_stool, f.Red_bighook},
		7:  {f.Blue_lighter, f.Red_bighook, f.Blue_bighook, f.Blue_v},
		8:  {f.Yellow_smallhook, f.Blue_lighter, f.Yellow_bighook, f.Red_flash},
		9:  {f.Green_bighook, f.Yellow_smallhook, f.Green_L, f.Red_bighook},
		10: {f.Green_L, f.Yellow_smallhook, f.Green_flash, f.Yellow_hello}}
	problems = append(problems, createDifficultCardProblems(32, topShape, bottomShape, blockNums, f)...)

	// B33
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, 0, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Red_stool, f.Red_bighook, f.Blue_lighter, f.Green_bighook},
		2:  {f.Yellow_hello, f.Yellow_gate, f.Blue_lighter, f.Green_bighook},
		3:  {f.Blue_lighter, f.Blue_flash, f.Yellow_gate, f.Yellow_bighook},
		4:  {f.Green_bighook, f.Yellow_bighook, f.Red_stool, f.Yellow_gate},
		5:  {f.Yellow_hello, f.Blue_lighter, f.Green_flash, f.Yellow_gate},
		6:  {f.Red_bighook, f.Blue_v, f.Green_flash, f.Blue_lighter},
		7:  {f.Yellow_hello, f.Red_smallhook, f.Green_L, f.Red_stool},
		8:  {f.Red_bighook, f.Blue_bighook, f.Green_T, f.Green_L},
		9:  {f.Yellow_smallhook, f.Blue_bighook, f.Blue_lighter, f.Green_L},
		10: {f.Yellow_bighook, f.Green_flash, f.Green_L, f.Red_flash}}
	problems = append(problems, createDifficultCardProblems(33, topShape, bottomShape, blockNums, f)...)

	// B34
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {0, 0, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {0, 0, 0}, {0, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Red_bighook, f.Yellow_hello, f.Yellow_bighook, f.Green_flash},
		2:  {f.Red_bighook, f.Green_bighook, f.Red_stool, f.Green_flash},
		3:  {f.Red_bighook, f.Yellow_bighook, f.Blue_lighter, f.Red_stool},
		4:  {f.Blue_lighter, f.Red_bighook, f.Green_flash, f.Red_stool},
		5:  {f.Blue_bighook, f.Red_stool, f.Green_bighook, f.Blue_lighter},
		6:  {f.Green_bighook, f.Red_stool, f.Blue_v, f.Blue_lighter},
		7:  {f.Yellow_hello, f.Red_smallhook, f.Blue_bighook, f.Green_L},
		8:  {f.Green_bighook, f.Red_bighook, f.Green_T, f.Red_smallhook},
		9:  {f.Blue_flash, f.Red_stool, f.Blue_v, f.Green_flash},
		10: {f.Green_L, f.Yellow_hello, f.Red_smallhook, f.Red_stool}}
	problems = append(problems, createDifficultCardProblems(34, topShape, bottomShape, blockNums, f)...)

	// B35
	topShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {0, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int][]*Block{
		1:  {f.Blue_bighook, f.Red_stool, f.Green_flash, f.Green_bighook},
		2:  {f.Yellow_hello, f.Green_flash, f.Red_stool, f.Yellow_bighook},
		3:  {f.Red_stool, f.Green_bighook, f.Blue_lighter, f.Yellow_hello},
		4:  {f.Yellow_bighook, f.Green_bighook, f.Yellow_hello, f.Red_stool},
		5:  {f.Yellow_gate, f.Red_stool, f.Blue_lighter, f.Blue_bighook},
		6:  {f.Green_flash, f.Blue_v, f.Red_bighook, f.Green_bighook},
		7:  {f.Green_L, f.Red_stool, f.Red_bighook, f.Red_flash},
		8:  {f.Blue_lighter, f.Red_stool, f.Green_L, f.Red_smallhook},
		9:  {f.Green_flash, f.Yellow_smallhook, f.Red_bighook, f.Green_L},
		10: {f.Green_L, f.Red_smallhook, f.Blue_lighter, f.Blue_flash}}
	problems = append(problems, createDifficultCardProblems(35, topShape, bottomShape, blockNums, f)...)

	// B36
	topShape = NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, 0}, {0, 0, 0}, {-1, 0, -1}})
	blockNums = map[int][]*Block{
		1:  {f.Yellow_bighook, f.Yellow_gate, f.Yellow_hello, f.Red_stool},
		2:  {f.Blue_bighook, f.Red_stool, f.Yellow_gate, f.Blue_flash},
		3:  {f.Green_bighook, f.Blue_bighook, f.Yellow_hello, f.Blue_flash},
		4:  {f.Yellow_gate, f.Blue_bighook, f.Yellow_hello, f.Green_flash},
		5:  {f.Yellow_hello, f.Red_stool, f.Yellow_bighook, f.Green_bighook},
		6:  {f.Green_L, f.Red_smallhook, f.Blue_bighook, f.Red_stool},
		7:  {f.Yellow_smallhook, f.Green_L, f.Red_bighook, f.Green_bighook},
		8:  {f.Green_flash, f.Yellow_smallhook, f.Green_L, f.Yellow_hello},
		9:  {f.Green_L, f.Red_smallhook, f.Red_stool, f.Yellow_hello},
		10: {f.Blue_v, f.Yellow_hello, f.Green_flash, f.Blue_flash}}
	problems = append(problems, createDifficultCardProblems(36, topShape, bottomShape, blockNums, f)...)

	return problems
}

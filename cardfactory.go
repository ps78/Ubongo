package main

import (
	"sync"
)

// ******************************************************************
// Public elements
// ******************************************************************

// Represents the singleton card factory. Get instance with GetCardFactory()
type CardFactory struct {
	// Contains all cards in map with 3 keys: [Difficulty][cardNumber][DiceNumber]
	Cards map[UbongoDifficulty](map[int]*Card)
}

// Returns the singleton instance of the problem factory
func GetCardFactory() *CardFactory {
	// Create the singleton instance
	onceCardFactorySingleton.Do(func() {
		bf := GetBlockFactory()

		f := new(CardFactory)

		f.Cards = make(map[UbongoDifficulty]map[int]*Card)

		// create all standard easy and difficult problems in a flat slice
		allCards := createAllEasyCards(bf)
		allCards = append(allCards, createAllDifficultCards(bf)...)

		// insert all problems in the 3-level map f.Problems[difficulty][cardNum][DiceNum]
		for _, p := range allCards {
			if _, ok := f.Cards[p.Difficulty]; !ok {
				f.Cards[p.Difficulty] = make(map[int]*Card)
			}
			f.Cards[p.Difficulty][p.CardNumber] = p
		}

		cardFactoryInstance = f
	})
	return cardFactoryInstance
}

// Returns the problem with the given parameters if it exists, nil otherwise
func (f *CardFactory) Get(difficulty UbongoDifficulty, cardNumber int) *Card {
	if _, okDiff := f.Cards[difficulty]; okDiff {
		if _, okCard := f.Cards[difficulty][cardNumber]; okCard {
			return f.Cards[difficulty][cardNumber]
		}
	}
	return nil
}

// Returns all cards as a slice
func (f *CardFactory) GetByAnimal(difficulty UbongoDifficulty, animal UbongoAnimal) []*Card {
	result := make([]*Card, 0)
	for _, card := range f.Cards[difficulty] {
		if card.Animal == animal {
			result = append(result, card)
		}
	}
	return result
}

// Returns all cards as a slice
func (f *CardFactory) GetAll() []*Card {
	result := make([]*Card, 0)
	for _, diffV := range f.Cards {
		for _, numV := range diffV {
			result = append(result, numV)
		}
	}
	return result
}

// ******************************************************************
// Private elements
// ******************************************************************

// used to create a thread-safe singleton instance of a CardFactory
var onceCardFactorySingleton sync.Once

// the singleton
var cardFactoryInstance *CardFactory

// animalByCardNum is used to assign animal to a card number
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

// Creates all easy cards of a specific card (as many as there are keys in the blocks-map)
func createEasyCard(cardNum int, topShape, bottomShape *Array2d, blocks map[int]*Blockset, f *BlockFactory) *Card {
	problems := make(map[int]*Problem)

	for dice, blockset := range blocks {
		var shape *Array2d
		if dice <= 4 {
			shape = topShape
		} else {
			shape = bottomShape
		}
		problems[dice] = NewProblem(shape, 2, blockset)
	}

	return NewCard(cardNum, Easy, animalByCardNum[cardNum], problems)
}

// Creates all difficult cards of a specific card (as many as there are keys in the blocks-map)
func createDifficultCard(cardNum int, topShape, bottomShape *Array2d, blocks map[int]*Blockset, f *BlockFactory) *Card {
	problems := make(map[int]*Problem)

	for dice, blockset := range blocks {
		var shape *Array2d
		if dice <= 5 {
			shape = topShape
		} else {
			shape = bottomShape
		}
		problems[dice] = NewProblem(shape, 2, blockset)
	}

	return NewCard(cardNum, Difficult, animalByCardNum[cardNum], problems)
}

// Creates all the problems from the original Ubongo game with the difficulty 'Easy'
// Returns a slice with 144 elements
func createAllEasyCards(f *BlockFactory) []*Card {
	cards := make([]*Card, 0)

	// A1
	topShape := NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, -1, 0}, {-1, 0, 0}, {0, 0, 0}})
	bottomShape := NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, -1, -1}})
	blockNums := map[int]*Blockset{
		1: NewBlockset(f.Red_bighook, f.Green_L, f.Blue_lighter),
		3: NewBlockset(f.Blue_lighter, f.Green_bighook, f.Red_smallhook),
		5: NewBlockset(f.Yellow_bighook, f.Blue_bighook, f.Red_smallhook),
		8: NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Green_flash)}
	cards = append(cards, createEasyCard(1, topShape, bottomShape, blockNums, f))

	// A2
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_stool, f.Green_L, f.Yellow_hello),
		3: NewBlockset(f.Blue_lighter, f.Red_flash, f.Yellow_hello),
		5: NewBlockset(f.Green_L, f.Green_bighook, f.Red_bighook),
		8: NewBlockset(f.Blue_flash, f.Red_stool, f.Green_L)}
	cards = append(cards, createEasyCard(2, topShape, bottomShape, blockNums, f))

	// A3
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_lighter, f.Yellow_smallhook, f.Blue_bighook),
		3: NewBlockset(f.Red_bighook, f.Green_L, f.Red_stool),
		5: NewBlockset(f.Blue_lighter, f.Red_smallhook, f.Red_stool),
		8: NewBlockset(f.Blue_bighook, f.Green_L, f.Blue_lighter)}
	cards = append(cards, createEasyCard(3, topShape, bottomShape, blockNums, f))

	// A4
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Green_L, f.Blue_lighter, f.Yellow_bighook),
		3: NewBlockset(f.Yellow_hello, f.Green_L, f.Blue_lighter),
		5: NewBlockset(f.Green_L, f.Red_bighook, f.Blue_bighook),
		8: NewBlockset(f.Yellow_bighook, f.Green_L, f.Blue_bighook)}
	cards = append(cards, createEasyCard(4, topShape, bottomShape, blockNums, f))

	// A5
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, -1, 0}, {0, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {0, 0, -1}, {0, -1, -1}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Yellow_bighook, f.Red_stool, f.Green_L),
		3: NewBlockset(f.Yellow_smallhook, f.Blue_bighook, f.Red_bighook),
		5: NewBlockset(f.Red_stool, f.Yellow_hello, f.Red_smallhook),
		8: NewBlockset(f.Yellow_gate, f.Green_L, f.Red_stool)}
	cards = append(cards, createEasyCard(5, topShape, bottomShape, blockNums, f))

	// A6
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_bighook, f.Red_stool, f.Green_L),
		3: NewBlockset(f.Yellow_bighook, f.Green_bighook, f.Red_smallhook),
		5: NewBlockset(f.Yellow_gate, f.Blue_bighook, f.Red_smallhook),
		8: NewBlockset(f.Blue_lighter, f.Green_T, f.Yellow_hello)}
	cards = append(cards, createEasyCard(6, topShape, bottomShape, blockNums, f))

	// A7
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_flash, f.Green_L, f.Blue_lighter),
		3: NewBlockset(f.Green_T, f.Blue_lighter, f.Green_flash),
		5: NewBlockset(f.Yellow_hello, f.Blue_lighter, f.Yellow_smallhook),
		8: NewBlockset(f.Yellow_hello, f.Green_L, f.Blue_lighter)}
	cards = append(cards, createEasyCard(7, topShape, bottomShape, blockNums, f))

	// A8
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_lighter, f.Red_smallhook, f.Blue_bighook),
		3: NewBlockset(f.Green_bighook, f.Blue_lighter, f.Yellow_smallhook),
		5: NewBlockset(f.Yellow_bighook, f.Green_L, f.Yellow_gate),
		8: NewBlockset(f.Yellow_gate, f.Green_bighook, f.Yellow_smallhook)}
	cards = append(cards, createEasyCard(8, topShape, bottomShape, blockNums, f))

	// A9
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_stool, f.Red_bighook, f.Red_flash),
		3: NewBlockset(f.Green_L, f.Green_flash, f.Blue_lighter),
		5: NewBlockset(f.Green_T, f.Red_stool, f.Blue_flash),
		8: NewBlockset(f.Blue_bighook, f.Yellow_smallhook, f.Red_bighook)}
	cards = append(cards, createEasyCard(9, topShape, bottomShape, blockNums, f))

	// A10
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_smallhook, f.Green_bighook, f.Red_stool),
		3: NewBlockset(f.Red_stool, f.Yellow_smallhook, f.Blue_flash),
		5: NewBlockset(f.Yellow_smallhook, f.Blue_flash, f.Blue_bighook),
		8: NewBlockset(f.Red_smallhook, f.Yellow_bighook, f.Red_stool)}
	cards = append(cards, createEasyCard(10, topShape, bottomShape, blockNums, f))

	// A11
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_lighter, f.Green_L, f.Blue_flash),
		3: NewBlockset(f.Green_bighook, f.Green_flash, f.Yellow_smallhook),
		5: NewBlockset(f.Green_bighook, f.Red_smallhook, f.Red_stool),
		8: NewBlockset(f.Red_smallhook, f.Blue_bighook, f.Red_stool)}
	cards = append(cards, createEasyCard(11, topShape, bottomShape, blockNums, f))

	// A12
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Yellow_hello),
		3: NewBlockset(f.Green_bighook, f.Red_smallhook, f.Blue_lighter),
		5: NewBlockset(f.Green_bighook, f.Green_flash, f.Red_smallhook),
		8: NewBlockset(f.Yellow_smallhook, f.Green_bighook, f.Red_stool)}
	cards = append(cards, createEasyCard(12, topShape, bottomShape, blockNums, f))

	// A13
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {0, 0, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_bighook, f.Yellow_hello, f.Yellow_smallhook),
		3: NewBlockset(f.Blue_lighter, f.Red_smallhook, f.Blue_flash),
		5: NewBlockset(f.Blue_flash, f.Green_L, f.Red_stool),
		8: NewBlockset(f.Green_flash, f.Blue_lighter, f.Green_L)}
	cards = append(cards, createEasyCard(13, topShape, bottomShape, blockNums, f))

	// A14
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Yellow_gate, f.Red_smallhook, f.Red_bighook),
		3: NewBlockset(f.Red_smallhook, f.Blue_bighook, f.Red_stool),
		5: NewBlockset(f.Blue_lighter, f.Red_flash, f.Blue_flash),
		8: NewBlockset(f.Green_bighook, f.Green_T, f.Red_stool)}
	cards = append(cards, createEasyCard(14, topShape, bottomShape, blockNums, f))

	// A15
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Green_bighook),
		3: NewBlockset(f.Green_flash, f.Red_smallhook, f.Blue_lighter),
		5: NewBlockset(f.Green_L, f.Red_stool, f.Green_flash),
		8: NewBlockset(f.Green_bighook, f.Green_L, f.Blue_flash)}
	cards = append(cards, createEasyCard(15, topShape, bottomShape, blockNums, f))

	// A16
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Yellow_smallhook, f.Green_flash, f.Blue_lighter),
		3: NewBlockset(f.Green_L, f.Red_stool, f.Green_flash),
		5: NewBlockset(f.Red_smallhook, f.Red_stool, f.Green_bighook),
		8: NewBlockset(f.Green_L, f.Green_flash, f.Blue_bighook)}
	cards = append(cards, createEasyCard(16, topShape, bottomShape, blockNums, f))

	// A17
	topShape = NewArray2dFromData([][]int8{{-1, 0}, {0, 0}, {0, 0}, {0, -1}, {0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Yellow_gate, f.Blue_bighook, f.Green_L),
		3: NewBlockset(f.Yellow_hello, f.Green_L, f.Blue_lighter),
		5: NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Yellow_hello),
		8: NewBlockset(f.Blue_lighter, f.Green_flash, f.Yellow_smallhook)}
	cards = append(cards, createEasyCard(17, topShape, bottomShape, blockNums, f))

	// A18
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_lighter, f.Red_bighook, f.Green_L),
		3: NewBlockset(f.Yellow_hello, f.Red_flash, f.Blue_lighter),
		5: NewBlockset(f.Blue_lighter, f.Yellow_smallhook, f.Blue_flash),
		8: NewBlockset(f.Green_L, f.Red_bighook, f.Red_stool)}
	cards = append(cards, createEasyCard(18, topShape, bottomShape, blockNums, f))

	// A19
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, -1}, {-1, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_smallhook, f.Green_flash, f.Red_stool),
		3: NewBlockset(f.Blue_bighook, f.Green_L, f.Red_stool),
		5: NewBlockset(f.Red_stool, f.Green_L, f.Yellow_bighook),
		8: NewBlockset(f.Green_flash, f.Red_smallhook, f.Red_stool)}
	cards = append(cards, createEasyCard(19, topShape, bottomShape, blockNums, f))

	// A20
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_stool, f.Yellow_hello, f.Red_smallhook),
		3: NewBlockset(f.Red_smallhook, f.Blue_bighook, f.Red_stool),
		5: NewBlockset(f.Red_flash, f.Red_stool, f.Yellow_bighook),
		8: NewBlockset(f.Red_smallhook, f.Blue_lighter, f.Blue_flash)}
	cards = append(cards, createEasyCard(20, topShape, bottomShape, blockNums, f))

	// A21
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, 0}, {-1, 0, -1}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {-1, 0, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_lighter, f.Yellow_bighook, f.Yellow_smallhook),
		3: NewBlockset(f.Blue_lighter, f.Red_stool, f.Green_L),
		5: NewBlockset(f.Green_bighook, f.Red_smallhook, f.Red_stool),
		8: NewBlockset(f.Red_smallhook, f.Blue_lighter, f.Yellow_hello)}
	cards = append(cards, createEasyCard(21, topShape, bottomShape, blockNums, f))

	// A22
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_bighook, f.Red_stool, f.Green_L),
		3: NewBlockset(f.Yellow_bighook, f.Red_stool, f.Green_T),
		5: NewBlockset(f.Blue_flash, f.Red_smallhook, f.Blue_bighook),
		8: NewBlockset(f.Blue_bighook, f.Green_L, f.Yellow_hello)}
	cards = append(cards, createEasyCard(22, topShape, bottomShape, blockNums, f))

	// A23
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_lighter, f.Red_bighook, f.Red_smallhook),
		3: NewBlockset(f.Blue_flash, f.Red_smallhook, f.Blue_lighter),
		5: NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Blue_bighook),
		8: NewBlockset(f.Red_stool, f.Green_L, f.Green_bighook)}
	cards = append(cards, createEasyCard(23, topShape, bottomShape, blockNums, f))

	// A24
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_stool, f.Green_bighook, f.Green_L),
		3: NewBlockset(f.Yellow_hello, f.Green_L, f.Red_stool),
		5: NewBlockset(f.Yellow_smallhook, f.Yellow_bighook, f.Blue_lighter),
		8: NewBlockset(f.Green_L, f.Yellow_bighook, f.Green_flash)}
	cards = append(cards, createEasyCard(24, topShape, bottomShape, blockNums, f))

	// A25
	topShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, -1}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_smallhook, f.Yellow_bighook, f.Blue_v),
		3: NewBlockset(f.Blue_v, f.Red_smallhook, f.Yellow_hello),
		5: NewBlockset(f.Yellow_smallhook, f.Green_flash, f.Blue_lighter),
		8: NewBlockset(f.Red_smallhook, f.Green_bighook, f.Red_stool)}
	cards = append(cards, createEasyCard(25, topShape, bottomShape, blockNums, f))

	// A26
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_v, f.Blue_bighook, f.Yellow_smallhook),
		3: NewBlockset(f.Red_flash, f.Green_L, f.Red_smallhook),
		5: NewBlockset(f.Blue_flash, f.Red_stool, f.Green_L),
		8: NewBlockset(f.Green_flash, f.Green_bighook, f.Green_L)}
	cards = append(cards, createEasyCard(26, topShape, bottomShape, blockNums, f))

	// A27
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_v, f.Green_bighook, f.Red_smallhook),
		3: NewBlockset(f.Blue_flash, f.Green_T, f.Blue_v),
		5: NewBlockset(f.Yellow_smallhook, f.Blue_lighter, f.Yellow_hello),
		8: NewBlockset(f.Green_flash, f.Red_stool, f.Green_L)}
	cards = append(cards, createEasyCard(27, topShape, bottomShape, blockNums, f))

	// A28
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_v, f.Yellow_smallhook, f.Green_bighook),
		3: NewBlockset(f.Green_T, f.Green_flash, f.Blue_v),
		5: NewBlockset(f.Yellow_smallhook, f.Blue_bighook, f.Red_stool),
		8: NewBlockset(f.Blue_bighook, f.Green_L, f.Blue_flash)}
	cards = append(cards, createEasyCard(28, topShape, bottomShape, blockNums, f))

	// A29
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_bighook, f.Red_smallhook, f.Blue_v),
		3: NewBlockset(f.Blue_lighter, f.Yellow_smallhook, f.Blue_v),
		5: NewBlockset(f.Green_bighook, f.Red_stool, f.Green_L),
		8: NewBlockset(f.Green_bighook, f.Blue_lighter, f.Green_L)}
	cards = append(cards, createEasyCard(29, topShape, bottomShape, blockNums, f))

	// A30
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {-1, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Green_flash, f.Red_stool, f.Red_flash),
		3: NewBlockset(f.Green_flash, f.Red_stool, f.Green_L),
		5: NewBlockset(f.Red_stool, f.Blue_bighook, f.Red_smallhook),
		8: NewBlockset(f.Blue_flash, f.Yellow_smallhook, f.Red_stool)}
	cards = append(cards, createEasyCard(30, topShape, bottomShape, blockNums, f))

	// A31
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_stool, f.Yellow_smallhook, f.Blue_v),
		3: NewBlockset(f.Green_bighook, f.Blue_v, f.Green_L),
		5: NewBlockset(f.Yellow_hello, f.Green_flash, f.Yellow_smallhook),
		8: NewBlockset(f.Red_smallhook, f.Blue_flash, f.Red_stool)}
	cards = append(cards, createEasyCard(31, topShape, bottomShape, blockNums, f))

	// A32
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_stool, f.Green_flash, f.Red_flash),
		3: NewBlockset(f.Blue_bighook, f.Red_smallhook, f.Red_stool),
		5: NewBlockset(f.Green_flash, f.Red_stool, f.Green_L),
		8: NewBlockset(f.Green_L, f.Blue_lighter, f.Green_flash)}
	cards = append(cards, createEasyCard(32, topShape, bottomShape, blockNums, f))

	// A33
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {0, 0, -1}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Blue_v, f.Green_bighook, f.Yellow_smallhook),
		3: NewBlockset(f.Blue_v, f.Green_L, f.Green_flash),
		5: NewBlockset(f.Green_bighook, f.Green_L, f.Red_stool),
		8: NewBlockset(f.Yellow_hello, f.Yellow_gate, f.Yellow_smallhook)}
	cards = append(cards, createEasyCard(33, topShape, bottomShape, blockNums, f))

	// A34
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Red_stool, f.Blue_v, f.Red_smallhook),
		3: NewBlockset(f.Yellow_hello, f.Red_flash, f.Blue_v),
		5: NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Blue_bighook),
		8: NewBlockset(f.Yellow_hello, f.Blue_bighook, f.Green_L)}
	cards = append(cards, createEasyCard(34, topShape, bottomShape, blockNums, f))

	// A35
	topShape = NewArray2dFromData([][]int8{{-1, 0}, {-1, 0}, {0, 0}, {0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Yellow_bighook, f.Blue_v, f.Green_L),
		3: NewBlockset(f.Red_flash, f.Green_L, f.Red_smallhook),
		5: NewBlockset(f.Red_smallhook, f.Yellow_bighook, f.Blue_lighter),
		8: NewBlockset(f.Blue_flash, f.Yellow_gate, f.Red_smallhook)}
	cards = append(cards, createEasyCard(35, topShape, bottomShape, blockNums, f))

	// A36
	blockNums = map[int]*Blockset{
		1: NewBlockset(f.Green_L, f.Yellow_hello, f.Blue_v),
		3: NewBlockset(f.Blue_bighook, f.Green_L, f.Blue_v),
		5: NewBlockset(f.Red_smallhook, f.Blue_flash, f.Yellow_hello),
		8: NewBlockset(f.Red_stool, f.Red_smallhook, f.Blue_bighook)}
	cards = append(cards, createEasyCard(36, topShape, bottomShape, blockNums, f))

	return cards
}

// Creates all the problems from the original Ubongo game with the difficulty 'Easy'
// Returns a slice with 144 elements
func createAllDifficultCards(f *BlockFactory) []*Card {
	cards := make([]*Card, 0)

	// B1
	topShape := NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	bottomShape := NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums := map[int]*Blockset{
		1:  NewBlockset(f.Red_smallhook, f.Blue_lighter, f.Green_L, f.Blue_v),
		2:  NewBlockset(f.Red_bighook, f.Red_smallhook, f.Blue_v, f.Red_flash),
		3:  NewBlockset(f.Blue_v, f.Blue_bighook, f.Green_T, f.Green_L),
		4:  NewBlockset(f.Green_flash, f.Green_T, f.Blue_v, f.Green_L),
		5:  NewBlockset(f.Red_flash, f.Blue_v, f.Green_L, f.Blue_bighook),
		6:  NewBlockset(f.Green_bighook, f.Yellow_smallhook, f.Green_L, f.Blue_v),
		7:  NewBlockset(f.Yellow_bighook, f.Red_smallhook, f.Blue_v, f.Green_T),
		8:  NewBlockset(f.Blue_flash, f.Green_L, f.Blue_v, f.Red_smallhook),
		9:  NewBlockset(f.Red_smallhook, f.Green_L, f.Red_bighook, f.Blue_v),
		10: NewBlockset(f.Red_bighook, f.Yellow_smallhook, f.Blue_v, f.Red_flash)}
	cards = append(cards, createDifficultCard(1, topShape, bottomShape, blockNums, f))

	// B2
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0, 0}, {-1, 0, 0, 0}, {-1, 0, 0, -1}, {0, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, -1, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_bighook, f.Yellow_smallhook, f.Green_flash, f.Red_smallhook),
		2:  NewBlockset(f.Red_stool, f.Blue_flash, f.Green_flash, f.Blue_v),
		3:  NewBlockset(f.Yellow_smallhook, f.Green_bighook, f.Green_flash, f.Red_flash),
		4:  NewBlockset(f.Yellow_smallhook, f.Green_bighook, f.Green_flash, f.Red_flash),
		5:  NewBlockset(f.Red_stool, f.Green_bighook, f.Green_L, f.Red_smallhook),
		6:  NewBlockset(f.Yellow_smallhook, f.Red_smallhook, f.Blue_v, f.Blue_bighook),
		7:  NewBlockset(f.Red_smallhook, f.Red_bighook, f.Green_L, f.Blue_v),
		8:  NewBlockset(f.Blue_v, f.Green_L, f.Blue_flash, f.Yellow_smallhook),
		9:  NewBlockset(f.Blue_v, f.Blue_bighook, f.Green_L, f.Red_flash),
		10: NewBlockset(f.Green_bighook, f.Red_smallhook, f.Blue_v, f.Yellow_smallhook)}
	cards = append(cards, createDifficultCard(2, topShape, bottomShape, blockNums, f))

	// B3
	topShape = NewArray2dFromData([][]int8{{0, -1}, {0, 0}, {0, 0}, {0, 0}, {-1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_flash, f.Blue_v, f.Green_L, f.Blue_flash),
		2:  NewBlockset(f.Green_T, f.Red_smallhook, f.Green_flash, f.Blue_v),
		3:  NewBlockset(f.Yellow_hello, f.Green_L, f.Green_T, f.Blue_v),
		4:  NewBlockset(f.Blue_v, f.Red_flash, f.Blue_bighook, f.Green_L),
		5:  NewBlockset(f.Red_smallhook, f.Yellow_hello, f.Yellow_smallhook, f.Blue_v),
		6:  NewBlockset(f.Green_L, f.Blue_v, f.Yellow_smallhook, f.Green_flash),
		7:  NewBlockset(f.Green_L, f.Red_flash, f.Green_bighook, f.Blue_v),
		8:  NewBlockset(f.Red_flash, f.Green_L, f.Blue_v, f.Yellow_hello),
		9:  NewBlockset(f.Green_flash, f.Red_smallhook, f.Red_flash, f.Blue_v),
		10: NewBlockset(f.Red_smallhook, f.Yellow_bighook, f.Blue_v, f.Red_flash)}
	cards = append(cards, createDifficultCard(3, topShape, bottomShape, blockNums, f))

	// B4
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, -1}, {0, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_L, f.Blue_v, f.Green_bighook, f.Red_smallhook),
		2:  NewBlockset(f.Red_smallhook, f.Blue_bighook, f.Red_flash, f.Blue_v),
		3:  NewBlockset(f.Green_bighook, f.Blue_v, f.Green_L, f.Red_flash),
		4:  NewBlockset(f.Yellow_bighook, f.Red_smallhook, f.Blue_v, f.Green_T),
		5:  NewBlockset(f.Green_L, f.Blue_v, f.Yellow_smallhook, f.Green_bighook),
		6:  NewBlockset(f.Blue_v, f.Green_T, f.Green_L, f.Yellow_bighook),
		7:  NewBlockset(f.Red_smallhook, f.Yellow_smallhook, f.Blue_v, f.Blue_lighter),
		8:  NewBlockset(f.Green_L, f.Green_flash, f.Green_T, f.Blue_v),
		9:  NewBlockset(f.Red_smallhook, f.Yellow_smallhook, f.Blue_v, f.Yellow_hello),
		10: NewBlockset(f.Blue_v, f.Green_T, f.Blue_flash, f.Red_smallhook)}
	cards = append(cards, createDifficultCard(4, topShape, bottomShape, blockNums, f))

	// B5
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0, 0}, {-1, 0, 0, -1}, {0, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Blue_v, f.Red_smallhook, f.Yellow_hello, f.Green_L),
		2:  NewBlockset(f.Blue_v, f.Red_smallhook, f.Yellow_hello, f.Red_flash),
		3:  NewBlockset(f.Green_T, f.Green_L, f.Green_flash, f.Blue_v),
		4:  NewBlockset(f.Yellow_bighook, f.Green_L, f.Red_flash, f.Blue_v),
		5:  NewBlockset(f.Blue_v, f.Red_smallhook, f.Blue_bighook, f.Green_L),
		6:  NewBlockset(f.Blue_bighook, f.Yellow_smallhook, f.Blue_v, f.Red_flash),
		7:  NewBlockset(f.Blue_v, f.Green_L, f.Red_smallhook, f.Green_bighook),
		8:  NewBlockset(f.Green_flash, f.Green_L, f.Blue_v, f.Yellow_smallhook),
		9:  NewBlockset(f.Green_L, f.Blue_v, f.Yellow_smallhook, f.Blue_lighter),
		10: NewBlockset(f.Blue_v, f.Green_bighook, f.Red_smallhook, f.Green_T)}
	cards = append(cards, createDifficultCard(5, topShape, bottomShape, blockNums, f))

	// B6
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Blue_v, f.Red_smallhook, f.Red_stool, f.Red_flash),
		2:  NewBlockset(f.Red_flash, f.Green_bighook, f.Yellow_smallhook, f.Blue_v),
		3:  NewBlockset(f.Yellow_smallhook, f.Yellow_hello, f.Red_smallhook, f.Blue_v),
		4:  NewBlockset(f.Blue_v, f.Green_L, f.Blue_flash, f.Red_smallhook),
		5:  NewBlockset(f.Blue_v, f.Green_L, f.Blue_flash, f.Yellow_smallhook),
		6:  NewBlockset(f.Red_smallhook, f.Blue_v, f.Red_flash, f.Yellow_hello),
		7:  NewBlockset(f.Blue_v, f.Green_T, f.Blue_lighter, f.Green_L),
		8:  NewBlockset(f.Green_L, f.Green_flash, f.Green_T, f.Blue_v),
		9:  NewBlockset(f.Green_bighook, f.Red_smallhook, f.Yellow_smallhook, f.Blue_v),
		10: NewBlockset(f.Red_flash, f.Blue_v, f.Green_L, f.Green_bighook)}
	cards = append(cards, createDifficultCard(6, topShape, bottomShape, blockNums, f))

	// B7
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1, -1}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_smallhook, f.Green_T, f.Blue_v, f.Green_flash),
		2:  NewBlockset(f.Blue_v, f.Red_smallhook, f.Blue_bighook, f.Green_L),
		3:  NewBlockset(f.Green_L, f.Blue_v, f.Yellow_bighook, f.Red_smallhook),
		4:  NewBlockset(f.Green_L, f.Blue_v, f.Yellow_smallhook, f.Blue_flash),
		5:  NewBlockset(f.Yellow_bighook, f.Yellow_smallhook, f.Red_smallhook, f.Blue_v),
		6:  NewBlockset(f.Yellow_smallhook, f.Blue_flash, f.Blue_v, f.Green_T),
		7:  NewBlockset(f.Yellow_bighook, f.Blue_v, f.Red_flash, f.Yellow_smallhook),
		8:  NewBlockset(f.Green_bighook, f.Blue_v, f.Yellow_smallhook, f.Green_L),
		9:  NewBlockset(f.Red_smallhook, f.Blue_v, f.Green_flash, f.Green_L),
		10: NewBlockset(f.Blue_v, f.Green_T, f.Red_smallhook, f.Blue_bighook)}
	cards = append(cards, createDifficultCard(7, topShape, bottomShape, blockNums, f))

	// B8
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1, -1}, {-1, 0, 0, 0}, {-1, 0, 0, -1}, {0, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {0, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_stool, f.Green_L, f.Yellow_smallhook, f.Blue_v),
		2:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Blue_v, f.Blue_flash),
		3:  NewBlockset(f.Red_smallhook, f.Red_flash, f.Blue_v, f.Green_bighook),
		4:  NewBlockset(f.Red_smallhook, f.Red_stool, f.Blue_v, f.Red_flash),
		5:  NewBlockset(f.Green_bighook, f.Blue_v, f.Green_L, f.Red_smallhook),
		6:  NewBlockset(f.Blue_v, f.Green_T, f.Red_smallhook, f.Yellow_bighook),
		7:  NewBlockset(f.Red_bighook, f.Yellow_smallhook, f.Green_L, f.Blue_v),
		8:  NewBlockset(f.Green_L, f.Red_smallhook, f.Yellow_hello, f.Blue_v),
		9:  NewBlockset(f.Green_L, f.Red_smallhook, f.Yellow_bighook, f.Blue_v),
		10: NewBlockset(f.Blue_bighook, f.Green_L, f.Blue_v, f.Yellow_smallhook)}
	cards = append(cards, createDifficultCard(8, topShape, bottomShape, blockNums, f))

	// B9
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {0, 0, 0}, {0, 0, 0}, {0, -1, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_L, f.Blue_v, f.Green_T, f.Yellow_bighook),
		2:  NewBlockset(f.Green_L, f.Blue_v, f.Red_smallhook, f.Blue_lighter),
		3:  NewBlockset(f.Red_smallhook, f.Yellow_smallhook, f.Blue_lighter, f.Blue_v),
		4:  NewBlockset(f.Yellow_bighook, f.Blue_v, f.Green_T, f.Red_flash),
		5:  NewBlockset(f.Blue_lighter, f.Blue_v, f.Green_L, f.Yellow_smallhook),
		6:  NewBlockset(f.Blue_lighter, f.Blue_v, f.Green_flash, f.Blue_bighook),
		7:  NewBlockset(f.Red_bighook, f.Red_stool, f.Green_L, f.Yellow_smallhook),
		8:  NewBlockset(f.Red_bighook, f.Blue_v, f.Green_bighook, f.Blue_lighter),
		9:  NewBlockset(f.Blue_flash, f.Red_smallhook, f.Yellow_smallhook, f.Blue_lighter),
		10: NewBlockset(f.Green_L, f.Blue_lighter, f.Blue_bighook, f.Yellow_smallhook)}
	cards = append(cards, createDifficultCard(9, topShape, bottomShape, blockNums, f))

	// B10
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {-1, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Yellow_smallhook, f.Blue_lighter, f.Blue_v, f.Green_L),
		2:  NewBlockset(f.Red_flash, f.Blue_v, f.Red_smallhook, f.Blue_lighter),
		3:  NewBlockset(f.Blue_v, f.Red_flash, f.Green_flash, f.Green_T),
		4:  NewBlockset(f.Red_smallhook, f.Red_flash, f.Green_flash, f.Blue_v),
		5:  NewBlockset(f.Blue_v, f.Yellow_smallhook, f.Green_L, f.Green_flash),
		6:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Blue_flash, f.Yellow_bighook),
		7:  NewBlockset(f.Red_stool, f.Blue_lighter, f.Green_flash, f.Blue_v),
		8:  NewBlockset(f.Blue_flash, f.Blue_bighook, f.Green_L, f.Yellow_smallhook),
		9:  NewBlockset(f.Red_flash, f.Green_L, f.Yellow_bighook, f.Blue_bighook),
		10: NewBlockset(f.Blue_v, f.Green_flash, f.Red_bighook, f.Blue_lighter)}
	cards = append(cards, createDifficultCard(10, topShape, bottomShape, blockNums, f))

	// B11
	topShape = NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {0, 0, -1}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_L, f.Red_smallhook, f.Blue_v, f.Red_stool),
		2:  NewBlockset(f.Blue_v, f.Green_L, f.Red_flash, f.Blue_bighook),
		3:  NewBlockset(f.Green_L, f.Red_smallhook, f.Blue_v, f.Blue_lighter),
		4:  NewBlockset(f.Yellow_smallhook, f.Green_bighook, f.Blue_v, f.Green_L),
		5:  NewBlockset(f.Yellow_gate, f.Green_L, f.Blue_v, f.Red_smallhook),
		6:  NewBlockset(f.Yellow_hello, f.Red_flash, f.Blue_lighter, f.Red_smallhook),
		7:  NewBlockset(f.Yellow_hello, f.Yellow_gate, f.Red_flash, f.Red_smallhook),
		8:  NewBlockset(f.Blue_bighook, f.Red_smallhook, f.Yellow_smallhook, f.Blue_lighter),
		9:  NewBlockset(f.Green_L, f.Green_T, f.Yellow_gate, f.Blue_bighook),
		10: NewBlockset(f.Red_smallhook, f.Green_L, f.Red_bighook, f.Green_flash)}
	cards = append(cards, createDifficultCard(11, topShape, bottomShape, blockNums, f))

	// B12
	topShape = NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {0, 0, 0}, {0, 0, -1}, {-1, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_stool, f.Green_L, f.Blue_v, f.Red_flash),
		2:  NewBlockset(f.Red_stool, f.Yellow_smallhook, f.Green_T, f.Blue_v),
		3:  NewBlockset(f.Red_stool, f.Yellow_smallhook, f.Blue_v, f.Green_L),
		4:  NewBlockset(f.Red_bighook, f.Red_smallhook, f.Blue_v, f.Green_T),
		5:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Blue_flash, f.Blue_v),
		6:  NewBlockset(f.Yellow_smallhook, f.Green_L, f.Red_stool, f.Green_bighook),
		7:  NewBlockset(f.Blue_v, f.Yellow_hello, f.Yellow_gate, f.Blue_flash),
		8:  NewBlockset(f.Blue_flash, f.Green_L, f.Green_flash, f.Red_smallhook),
		9:  NewBlockset(f.Red_stool, f.Red_flash, f.Yellow_bighook, f.Red_smallhook),
		10: NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Red_smallhook, f.Green_bighook)}
	cards = append(cards, createDifficultCard(12, topShape, bottomShape, blockNums, f))

	// 13
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, 0}, {0, 0, 0}, {0, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1}, {0, 0}, {0, 0}, {0, 0}, {0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Yellow_bighook, f.Red_smallhook, f.Blue_bighook, f.Yellow_smallhook),
		2:  NewBlockset(f.Yellow_smallhook, f.Red_bighook, f.Red_flash, f.Yellow_gate),
		3:  NewBlockset(f.Green_L, f.Green_flash, f.Yellow_smallhook, f.Blue_bighook),
		4:  NewBlockset(f.Yellow_hello, f.Yellow_gate, f.Yellow_bighook, f.Blue_v),
		5:  NewBlockset(f.Blue_v, f.Red_stool, f.Red_bighook, f.Yellow_hello),
		6:  NewBlockset(f.Green_L, f.Blue_bighook, f.Red_bighook, f.Yellow_smallhook),
		7:  NewBlockset(f.Red_stool, f.Red_smallhook, f.Green_T, f.Blue_bighook),
		8:  NewBlockset(f.Red_smallhook, f.Blue_lighter, f.Blue_flash, f.Green_L),
		9:  NewBlockset(f.Blue_v, f.Blue_bighook, f.Yellow_gate, f.Red_stool),
		10: NewBlockset(f.Yellow_hello, f.Red_flash, f.Red_smallhook, f.Red_stool)}
	cards = append(cards, createDifficultCard(13, topShape, bottomShape, blockNums, f))

	// B14
	topShape = NewArray2dFromData([][]int8{{0, -1, -1, -1}, {0, 0, -1, -1}, {0, 0, 0, -1}, {-1, 0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Blue_lighter, f.Yellow_smallhook, f.Yellow_hello, f.Green_L),
		2:  NewBlockset(f.Yellow_gate, f.Red_smallhook, f.Red_flash, f.Yellow_bighook),
		3:  NewBlockset(f.Green_bighook, f.Blue_lighter, f.Green_T, f.Red_smallhook),
		4:  NewBlockset(f.Red_flash, f.Green_L, f.Yellow_hello, f.Blue_lighter),
		5:  NewBlockset(f.Green_flash, f.Green_L, f.Yellow_smallhook, f.Blue_flash),
		6:  NewBlockset(f.Green_L, f.Green_bighook, f.Yellow_bighook, f.Red_smallhook),
		7:  NewBlockset(f.Yellow_smallhook, f.Red_smallhook, f.Green_flash, f.Red_stool),
		8:  NewBlockset(f.Blue_bighook, f.Red_smallhook, f.Blue_lighter, f.Yellow_smallhook),
		9:  NewBlockset(f.Yellow_gate, f.Red_stool, f.Green_T, f.Yellow_smallhook),
		10: NewBlockset(f.Green_flash, f.Red_smallhook, f.Red_flash, f.Yellow_hello)}
	cards = append(cards, createDifficultCard(14, topShape, bottomShape, blockNums, f))

	// B15
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {0, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Yellow_hello, f.Green_bighook, f.Blue_lighter, f.Blue_v),
		2:  NewBlockset(f.Yellow_smallhook, f.Blue_lighter, f.Red_smallhook, f.Yellow_hello),
		3:  NewBlockset(f.Green_L, f.Red_smallhook, f.Yellow_gate, f.Blue_flash),
		4:  NewBlockset(f.Blue_lighter, f.Green_L, f.Green_T, f.Blue_bighook),
		5:  NewBlockset(f.Red_smallhook, f.Green_bighook, f.Yellow_smallhook, f.Red_stool),
		6:  NewBlockset(f.Green_bighook, f.Red_smallhook, f.Red_stool, f.Green_L),
		7:  NewBlockset(f.Green_L, f.Red_stool, f.Blue_lighter, f.Red_smallhook),
		8:  NewBlockset(f.Blue_flash, f.Red_flash, f.Green_L, f.Red_stool),
		9:  NewBlockset(f.Red_bighook, f.Green_bighook, f.Green_L, f.Red_smallhook),
		10: NewBlockset(f.Green_L, f.Green_flash, f.Green_bighook, f.Red_smallhook)}
	cards = append(cards, createDifficultCard(15, topShape, bottomShape, blockNums, f))

	// B16
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {0, 0, -1}, {0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {0, 0, -1}, {0, 0, -1}, {0, -1, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_L, f.Blue_bighook, f.Red_stool, f.Yellow_smallhook),
		2:  NewBlockset(f.Red_smallhook, f.Green_L, f.Blue_lighter, f.Red_stool),
		3:  NewBlockset(f.Yellow_hello, f.Red_stool, f.Blue_v, f.Green_bighook),
		4:  NewBlockset(f.Blue_lighter, f.Red_smallhook, f.Green_L, f.Blue_bighook),
		5:  NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Red_flash, f.Yellow_hello),
		6:  NewBlockset(f.Green_flash, f.Yellow_smallhook, f.Green_L, f.Red_bighook),
		7:  NewBlockset(f.Green_flash, f.Green_T, f.Blue_lighter, f.Green_L),
		8:  NewBlockset(f.Green_bighook, f.Green_L, f.Yellow_smallhook, f.Green_flash),
		9:  NewBlockset(f.Yellow_bighook, f.Blue_lighter, f.Red_smallhook, f.Green_L),
		10: NewBlockset(f.Red_stool, f.Blue_v, f.Green_bighook, f.Blue_bighook)}
	cards = append(cards, createDifficultCard(16, topShape, bottomShape, blockNums, f))

	// B17
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {-1, 0, 0}, {0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Blue_lighter, f.Blue_bighook, f.Green_bighook, f.Blue_v),
		2:  NewBlockset(f.Red_stool, f.Green_L, f.Green_bighook, f.Yellow_smallhook),
		3:  NewBlockset(f.Green_T, f.Green_L, f.Red_stool, f.Green_bighook),
		4:  NewBlockset(f.Yellow_gate, f.Green_flash, f.Blue_v, f.Green_bighook),
		5:  NewBlockset(f.Yellow_smallhook, f.Green_L, f.Green_flash, f.Red_bighook),
		6:  NewBlockset(f.Yellow_smallhook, f.Yellow_hello, f.Blue_lighter, f.Green_L),
		7:  NewBlockset(f.Green_flash, f.Yellow_smallhook, f.Blue_bighook, f.Green_L),
		8:  NewBlockset(f.Blue_lighter, f.Yellow_smallhook, f.Yellow_bighook, f.Green_T),
		9:  NewBlockset(f.Blue_flash, f.Red_smallhook, f.Red_bighook, f.Green_L),
		10: NewBlockset(f.Green_L, f.Blue_lighter, f.Yellow_smallhook, f.Green_flash)}
	cards = append(cards, createDifficultCard(17, topShape, bottomShape, blockNums, f))

	// B18
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0, 0}, {-1, 0, 0, -1}, {0, 0, 0, 0}, {-1, -1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, 0, 0, -1}, {-1, -1, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_stool, f.Red_smallhook, f.Yellow_smallhook, f.Blue_flash),
		2:  NewBlockset(f.Yellow_smallhook, f.Blue_bighook, f.Blue_flash, f.Red_smallhook),
		3:  NewBlockset(f.Yellow_hello, f.Red_bighook, f.Red_smallhook, f.Yellow_smallhook),
		4:  NewBlockset(f.Red_stool, f.Red_bighook, f.Blue_v, f.Green_bighook),
		5:  NewBlockset(f.Blue_flash, f.Red_stool, f.Blue_bighook, f.Blue_v),
		6:  NewBlockset(f.Yellow_bighook, f.Red_smallhook, f.Green_L, f.Blue_flash),
		7:  NewBlockset(f.Red_smallhook, f.Green_L, f.Yellow_bighook, f.Red_stool),
		8:  NewBlockset(f.Blue_flash, f.Yellow_smallhook, f.Blue_lighter, f.Green_L),
		9:  NewBlockset(f.Green_L, f.Blue_bighook, f.Yellow_smallhook, f.Green_flash),
		10: NewBlockset(f.Green_L, f.Green_flash, f.Red_flash, f.Yellow_bighook)}
	cards = append(cards, createDifficultCard(18, topShape, bottomShape, blockNums, f))

	// B19
	topShape = NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {0, 0, -1}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {-1, 0, 0}, {0, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_stool, f.Green_bighook, f.Yellow_smallhook, f.Red_smallhook),
		2:  NewBlockset(f.Blue_v, f.Green_flash, f.Red_stool, f.Blue_bighook),
		3:  NewBlockset(f.Yellow_hello, f.Red_stool, f.Green_bighook, f.Blue_v),
		4:  NewBlockset(f.Red_stool, f.Blue_bighook, f.Green_L, f.Red_smallhook),
		5:  NewBlockset(f.Red_smallhook, f.Blue_flash, f.Red_stool, f.Red_flash),
		6:  NewBlockset(f.Green_L, f.Blue_bighook, f.Green_flash, f.Green_T),
		7:  NewBlockset(f.Green_L, f.Yellow_hello, f.Blue_lighter, f.Red_smallhook),
		8:  NewBlockset(f.Blue_lighter, f.Green_bighook, f.Red_smallhook, f.Yellow_smallhook),
		9:  NewBlockset(f.Blue_bighook, f.Blue_lighter, f.Green_L, f.Red_smallhook),
		10: NewBlockset(f.Blue_bighook, f.Red_stool, f.Green_bighook, f.Blue_v)}
	cards = append(cards, createDifficultCard(19, topShape, bottomShape, blockNums, f))

	// B20
	topShape = NewArray2dFromData([][]int8{{0, 0, 0, 0}, {-1, 0, 0, -1}, {-1, -1, 0, 0}, {-1, -1, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, -1}, {0, 0, -1}, {-1, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_L, f.Red_smallhook, f.Red_stool, f.Green_flash),
		2:  NewBlockset(f.Yellow_bighook, f.Blue_flash, f.Red_smallhook, f.Yellow_smallhook),
		3:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Blue_flash, f.Red_stool),
		4:  NewBlockset(f.Red_smallhook, f.Green_L, f.Blue_lighter, f.Red_stool),
		5:  NewBlockset(f.Blue_bighook, f.Green_flash, f.Blue_v, f.Red_stool),
		6:  NewBlockset(f.Green_bighook, f.Blue_flash, f.Yellow_smallhook, f.Green_L),
		7:  NewBlockset(f.Green_L, f.Red_stool, f.Red_flash, f.Blue_bighook),
		8:  NewBlockset(f.Green_T, f.Green_bighook, f.Blue_flash, f.Red_smallhook),
		9:  NewBlockset(f.Green_bighook, f.Yellow_smallhook, f.Green_L, f.Red_stool),
		10: NewBlockset(f.Green_L, f.Blue_bighook, f.Blue_lighter, f.Yellow_smallhook)}
	cards = append(cards, createDifficultCard(20, topShape, bottomShape, blockNums, f))

	// B21
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, -1}, {-1, 0, 0}, {0, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0, -1}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, 0, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Blue_flash, f.Green_bighook, f.Red_smallhook, f.Green_L),
		2:  NewBlockset(f.Red_stool, f.Red_smallhook, f.Red_flash, f.Green_bighook),
		3:  NewBlockset(f.Red_smallhook, f.Green_L, f.Green_flash, f.Red_stool),
		4:  NewBlockset(f.Blue_lighter, f.Red_bighook, f.Green_L, f.Red_smallhook),
		5:  NewBlockset(f.Green_L, f.Red_stool, f.Red_flash, f.Blue_flash),
		6:  NewBlockset(f.Red_smallhook, f.Blue_lighter, f.Green_L, f.Yellow_hello),
		7:  NewBlockset(f.Yellow_bighook, f.Green_L, f.Blue_flash, f.Red_smallhook),
		8:  NewBlockset(f.Green_L, f.Yellow_hello, f.Green_T, f.Blue_flash),
		9:  NewBlockset(f.Red_stool, f.Red_smallhook, f.Yellow_hello, f.Yellow_smallhook),
		10: NewBlockset(f.Red_bighook, f.Yellow_hello, f.Blue_v, f.Red_stool)}
	cards = append(cards, createDifficultCard(21, topShape, bottomShape, blockNums, f))

	// B22
	topShape = NewArray2dFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, -1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_stool, f.Blue_bighook, f.Blue_v, f.Blue_flash),
		2:  NewBlockset(f.Yellow_smallhook, f.Red_stool, f.Green_L, f.Red_bighook),
		3:  NewBlockset(f.Green_L, f.Blue_flash, f.Yellow_smallhook, f.Green_bighook),
		4:  NewBlockset(f.Yellow_smallhook, f.Green_L, f.Red_bighook, f.Blue_flash),
		5:  NewBlockset(f.Yellow_bighook, f.Blue_v, f.Blue_lighter, f.Green_bighook),
		6:  NewBlockset(f.Yellow_bighook, f.Blue_v, f.Blue_lighter, f.Red_stool),
		7:  NewBlockset(f.Blue_bighook, f.Blue_lighter, f.Red_bighook, f.Blue_v),
		8:  NewBlockset(f.Green_T, f.Red_stool, f.Yellow_smallhook, f.Blue_bighook),
		9:  NewBlockset(f.Green_bighook, f.Blue_v, f.Red_stool, f.Yellow_gate),
		10: NewBlockset(f.Red_stool, f.Green_flash, f.Yellow_smallhook, f.Green_T)}
	cards = append(cards, createDifficultCard(22, topShape, bottomShape, blockNums, f))

	// B23
	topShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, -1, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, -1, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_stool, f.Green_L, f.Green_bighook, f.Red_flash),
		2:  NewBlockset(f.Red_stool, f.Yellow_gate, f.Yellow_smallhook, f.Green_L),
		3:  NewBlockset(f.Blue_v, f.Blue_lighter, f.Blue_bighook, f.Red_stool),
		4:  NewBlockset(f.Blue_v, f.Red_stool, f.Yellow_hello, f.Yellow_bighook),
		5:  NewBlockset(f.Red_smallhook, f.Yellow_smallhook, f.Green_bighook, f.Blue_flash),
		6:  NewBlockset(f.Blue_flash, f.Yellow_gate, f.Blue_v, f.Green_bighook),
		7:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Yellow_gate, f.Blue_lighter),
		8:  NewBlockset(f.Blue_v, f.Yellow_hello, f.Blue_bighook, f.Green_bighook),
		9:  NewBlockset(f.Blue_lighter, f.Yellow_smallhook, f.Red_flash, f.Yellow_hello),
		10: NewBlockset(f.Yellow_hello, f.Green_flash, f.Green_L, f.Yellow_smallhook)}
	cards = append(cards, createDifficultCard(23, topShape, bottomShape, blockNums, f))

	// B24
	topShape = NewArray2dFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, 0}, {0, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0, 0}, {-1, 0, 0, -1}, {-1, 0, 0, -1}, {-1, -1, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Blue_bighook, f.Green_L, f.Red_stool, f.Red_smallhook),
		2:  NewBlockset(f.Yellow_smallhook, f.Blue_flash, f.Green_L, f.Red_bighook),
		3:  NewBlockset(f.Yellow_hello, f.Red_bighook, f.Blue_bighook, f.Blue_v),
		4:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Red_stool, f.Green_bighook),
		5:  NewBlockset(f.Yellow_smallhook, f.Green_L, f.Green_flash, f.Yellow_bighook),
		6:  NewBlockset(f.Green_L, f.Red_smallhook, f.Yellow_hello, f.Red_bighook),
		7:  NewBlockset(f.Blue_lighter, f.Green_flash, f.Blue_v, f.Red_stool),
		8:  NewBlockset(f.Blue_lighter, f.Red_stool, f.Green_L, f.Red_smallhook),
		9:  NewBlockset(f.Green_L, f.Red_smallhook, f.Yellow_bighook, f.Blue_bighook),
		10: NewBlockset(f.Green_bighook, f.Red_smallhook, f.Red_bighook, f.Green_L)}
	cards = append(cards, createDifficultCard(24, topShape, bottomShape, blockNums, f))

	// B25
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, -1}, {0, 0, 0}, {-1, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Blue_v, f.Red_bighook, f.Red_stool, f.Blue_lighter),
		2:  NewBlockset(f.Blue_bighook, f.Red_stool, f.Green_T, f.Yellow_smallhook),
		3:  NewBlockset(f.Green_T, f.Yellow_smallhook, f.Blue_lighter, f.Green_bighook),
		4:  NewBlockset(f.Yellow_smallhook, f.Blue_flash, f.Green_L, f.Yellow_gate),
		5:  NewBlockset(f.Green_bighook, f.Red_smallhook, f.Blue_lighter, f.Green_L),
		6:  NewBlockset(f.Yellow_smallhook, f.Blue_lighter, f.Green_L, f.Blue_flash),
		7:  NewBlockset(f.Blue_bighook, f.Yellow_smallhook, f.Blue_lighter, f.Green_L),
		8:  NewBlockset(f.Red_smallhook, f.Blue_lighter, f.Yellow_hello, f.Green_T),
		9:  NewBlockset(f.Yellow_hello, f.Yellow_smallhook, f.Green_L, f.Red_bighook),
		10: NewBlockset(f.Blue_bighook, f.Green_L, f.Yellow_hello, f.Red_smallhook)}
	cards = append(cards, createDifficultCard(25, topShape, bottomShape, blockNums, f))

	// B26
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}, {-1, -1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {-1, 0, 0}, {-1, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Yellow_smallhook, f.Yellow_hello, f.Green_L, f.Yellow_bighook),
		2:  NewBlockset(f.Yellow_bighook, f.Blue_v, f.Red_stool, f.Yellow_hello),
		3:  NewBlockset(f.Green_L, f.Red_stool, f.Blue_bighook, f.Red_smallhook),
		4:  NewBlockset(f.Green_bighook, f.Red_smallhook, f.Green_flash, f.Green_L),
		5:  NewBlockset(f.Blue_bighook, f.Green_bighook, f.Red_smallhook, f.Yellow_smallhook),
		6:  NewBlockset(f.Red_stool, f.Green_L, f.Green_flash, f.Yellow_smallhook),
		7:  NewBlockset(f.Red_bighook, f.Yellow_hello, f.Blue_v, f.Green_flash),
		8:  NewBlockset(f.Yellow_hello, f.Green_L, f.Red_flash, f.Green_flash),
		9:  NewBlockset(f.Green_T, f.Yellow_bighook, f.Red_stool, f.Green_L),
		10: NewBlockset(f.Green_flash, f.Yellow_hello, f.Yellow_smallhook, f.Red_smallhook)}
	cards = append(cards, createDifficultCard(26, topShape, bottomShape, blockNums, f))

	// B27
	topShape = NewArray2dFromData([][]int8{{0, 0, 0, -1}, {-1, 0, 0, 0}, {-1, -1, 0, 0}, {-1, -1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, 0, 0, 0}, {-1, -1, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Yellow_smallhook, f.Green_flash, f.Red_stool, f.Green_T),
		2:  NewBlockset(f.Blue_lighter, f.Green_L, f.Blue_bighook, f.Yellow_smallhook),
		3:  NewBlockset(f.Green_L, f.Green_flash, f.Red_smallhook, f.Blue_lighter),
		4:  NewBlockset(f.Yellow_gate, f.Blue_lighter, f.Green_bighook, f.Blue_v),
		5:  NewBlockset(f.Yellow_hello, f.Red_smallhook, f.Blue_lighter, f.Red_flash),
		6:  NewBlockset(f.Red_stool, f.Blue_v, f.Blue_bighook, f.Yellow_hello),
		7:  NewBlockset(f.Red_smallhook, f.Green_bighook, f.Green_L, f.Yellow_gate),
		8:  NewBlockset(f.Yellow_bighook, f.Green_L, f.Green_flash, f.Yellow_smallhook),
		9:  NewBlockset(f.Red_smallhook, f.Yellow_gate, f.Yellow_smallhook, f.Green_bighook),
		10: NewBlockset(f.Blue_v, f.Blue_bighook, f.Blue_lighter, f.Yellow_bighook)}
	cards = append(cards, createDifficultCard(27, topShape, bottomShape, blockNums, f))

	// B28
	topShape = NewArray2dFromData([][]int8{{0, 0, -1, -1}, {0, 0, 0, 0}, {-1, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {0, 0, -1}, {0, 0, -1}, {-1, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_T, f.Red_stool, f.Red_smallhook, f.Green_flash),
		2:  NewBlockset(f.Blue_v, f.Yellow_hello, f.Red_stool, f.Blue_lighter),
		3:  NewBlockset(f.Blue_flash, f.Red_stool, f.Green_flash, f.Blue_v),
		4:  NewBlockset(f.Yellow_smallhook, f.Red_smallhook, f.Red_stool, f.Yellow_hello),
		5:  NewBlockset(f.Green_T, f.Red_stool, f.Blue_flash, f.Yellow_smallhook),
		6:  NewBlockset(f.Blue_flash, f.Green_L, f.Yellow_smallhook, f.Yellow_gate),
		7:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Yellow_bighook, f.Blue_lighter),
		8:  NewBlockset(f.Green_L, f.Yellow_bighook, f.Red_smallhook, f.Blue_bighook),
		9:  NewBlockset(f.Blue_lighter, f.Blue_v, f.Red_stool, f.Blue_bighook),
		10: NewBlockset(f.Green_bighook, f.Blue_lighter, f.Green_flash, f.Blue_v)}
	cards = append(cards, createDifficultCard(28, topShape, bottomShape, blockNums, f))

	// B29
	topShape = NewArray2dFromData([][]int8{{-1, 0, 0, -1}, {0, 0, 0, 0}, {-1, 0, 0, -1}, {-1, 0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Yellow_smallhook, f.Blue_lighter, f.Blue_bighook, f.Red_flash),
		2:  NewBlockset(f.Red_stool, f.Blue_flash, f.Green_T, f.Yellow_smallhook),
		3:  NewBlockset(f.Green_T, f.Yellow_bighook, f.Yellow_smallhook, f.Green_bighook),
		4:  NewBlockset(f.Blue_v, f.Green_flash, f.Red_stool, f.Yellow_hello),
		5:  NewBlockset(f.Green_flash, f.Red_smallhook, f.Red_bighook, f.Yellow_smallhook),
		6:  NewBlockset(f.Green_L, f.Blue_bighook, f.Red_smallhook, f.Blue_lighter),
		7:  NewBlockset(f.Red_stool, f.Green_L, f.Blue_bighook, f.Red_smallhook),
		8:  NewBlockset(f.Blue_lighter, f.Blue_bighook, f.Green_T, f.Yellow_smallhook),
		9:  NewBlockset(f.Green_bighook, f.Blue_lighter, f.Blue_v, f.Yellow_gate),
		10: NewBlockset(f.Green_flash, f.Red_bighook, f.Green_L, f.Red_smallhook)}
	cards = append(cards, createDifficultCard(29, topShape, bottomShape, blockNums, f))

	// B30
	topShape = NewArray2dFromData([][]int8{{-1, 0, 0, -1}, {0, 0, 0, -1}, {-1, -1, 0, 0}, {-1, -1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, 0, -1}, {-1, 0, 0, -1}, {-1, 0, 0, 0}, {-1, 0, -1, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Yellow_smallhook, f.Green_L, f.Red_stool, f.Blue_bighook),
		2:  NewBlockset(f.Blue_bighook, f.Red_stool, f.Red_smallhook, f.Green_L),
		3:  NewBlockset(f.Red_stool, f.Blue_v, f.Blue_bighook, f.Green_bighook),
		4:  NewBlockset(f.Red_smallhook, f.Green_L, f.Red_stool, f.Green_bighook),
		5:  NewBlockset(f.Red_stool, f.Red_bighook, f.Yellow_smallhook, f.Green_L),
		6:  NewBlockset(f.Blue_v, f.Red_stool, f.Red_bighook, f.Blue_lighter),
		7:  NewBlockset(f.Green_L, f.Red_smallhook, f.Red_stool, f.Blue_lighter),
		8:  NewBlockset(f.Green_flash, f.Yellow_bighook, f.Blue_v, f.Green_bighook),
		9:  NewBlockset(f.Blue_flash, f.Blue_v, f.Blue_lighter, f.Blue_bighook),
		10: NewBlockset(f.Yellow_smallhook, f.Blue_bighook, f.Green_L, f.Yellow_hello)}
	cards = append(cards, createDifficultCard(30, topShape, bottomShape, blockNums, f))

	// B31
	topShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, -1}, {0, 0, -1}, {-1, 0, 0}, {-1, 0, 0}})
	bottomShape = NewArray2dFromData([][]int8{{-1, -1, 0, 0}, {0, 0, 0, -1}, {0, 0, 0, -1}, {0, -1, -1, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_bighook, f.Blue_lighter, f.Red_stool, f.Blue_v),
		2:  NewBlockset(f.Red_stool, f.Green_L, f.Yellow_hello, f.Red_flash),
		3:  NewBlockset(f.Green_flash, f.Yellow_smallhook, f.Red_stool, f.Red_smallhook),
		4:  NewBlockset(f.Green_bighook, f.Red_smallhook, f.Red_stool, f.Green_L),
		5:  NewBlockset(f.Yellow_smallhook, f.Green_L, f.Yellow_hello, f.Red_stool),
		6:  NewBlockset(f.Blue_bighook, f.Green_flash, f.Green_L, f.Yellow_smallhook),
		7:  NewBlockset(f.Red_flash, f.Red_smallhook, f.Blue_lighter, f.Yellow_bighook),
		8:  NewBlockset(f.Green_L, f.Red_flash, f.Red_stool, f.Yellow_hello),
		9:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Red_bighook, f.Green_flash),
		10: NewBlockset(f.Green_L, f.Blue_flash, f.Yellow_smallhook, f.Green_bighook)}
	cards = append(cards, createDifficultCard(31, topShape, bottomShape, blockNums, f))

	// B32
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {0, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, 0, -1}, {0, 0, 0, -1}, {-1, 0, 0, 0}, {-1, 0, -1, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Green_L, f.Yellow_smallhook, f.Blue_v, f.Green_bighook),
		2:  NewBlockset(f.Blue_v, f.Green_T, f.Blue_bighook, f.Green_L),
		3:  NewBlockset(f.Blue_v, f.Red_bighook, f.Yellow_smallhook, f.Red_flash),
		4:  NewBlockset(f.Green_L, f.Red_flash, f.Yellow_hello, f.Blue_v),
		5:  NewBlockset(f.Blue_bighook, f.Blue_v, f.Red_smallhook, f.Green_L),
		6:  NewBlockset(f.Red_flash, f.Red_smallhook, f.Red_stool, f.Red_bighook),
		7:  NewBlockset(f.Blue_lighter, f.Red_bighook, f.Blue_bighook, f.Blue_v),
		8:  NewBlockset(f.Yellow_smallhook, f.Blue_lighter, f.Yellow_bighook, f.Red_flash),
		9:  NewBlockset(f.Green_bighook, f.Yellow_smallhook, f.Green_L, f.Red_bighook),
		10: NewBlockset(f.Green_L, f.Yellow_smallhook, f.Green_flash, f.Yellow_hello)}
	cards = append(cards, createDifficultCard(32, topShape, bottomShape, blockNums, f))

	// B33
	topShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {-1, 0, 0}, {0, 0, 0}, {0, 0, 0}, {0, -1, -1}})
	bottomShape = NewArray2dFromData([][]int8{{-1, 0, -1}, {0, 0, -1}, {0, 0, 0}, {0, 0, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_stool, f.Red_bighook, f.Blue_lighter, f.Green_bighook),
		2:  NewBlockset(f.Yellow_hello, f.Yellow_gate, f.Blue_lighter, f.Green_bighook),
		3:  NewBlockset(f.Blue_lighter, f.Blue_flash, f.Yellow_gate, f.Yellow_bighook),
		4:  NewBlockset(f.Green_bighook, f.Yellow_bighook, f.Red_stool, f.Yellow_gate),
		5:  NewBlockset(f.Yellow_hello, f.Blue_lighter, f.Green_flash, f.Yellow_gate),
		6:  NewBlockset(f.Red_bighook, f.Blue_v, f.Green_flash, f.Blue_lighter),
		7:  NewBlockset(f.Yellow_hello, f.Red_smallhook, f.Green_L, f.Red_stool),
		8:  NewBlockset(f.Red_bighook, f.Blue_bighook, f.Green_T, f.Green_L),
		9:  NewBlockset(f.Yellow_smallhook, f.Blue_bighook, f.Blue_lighter, f.Green_L),
		10: NewBlockset(f.Yellow_bighook, f.Green_flash, f.Green_L, f.Red_flash)}
	cards = append(cards, createDifficultCard(33, topShape, bottomShape, blockNums, f))

	// B34
	topShape = NewArray2dFromData([][]int8{{-1, -1, 0}, {0, 0, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, -1, -1}, {0, 0, 0}, {0, 0, 0}, {0, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Red_bighook, f.Yellow_hello, f.Yellow_bighook, f.Green_flash),
		2:  NewBlockset(f.Red_bighook, f.Green_bighook, f.Red_stool, f.Green_flash),
		3:  NewBlockset(f.Red_bighook, f.Yellow_bighook, f.Blue_lighter, f.Red_stool),
		4:  NewBlockset(f.Blue_lighter, f.Red_bighook, f.Green_flash, f.Red_stool),
		5:  NewBlockset(f.Blue_bighook, f.Red_stool, f.Green_bighook, f.Blue_lighter),
		6:  NewBlockset(f.Green_bighook, f.Red_stool, f.Blue_v, f.Blue_lighter),
		7:  NewBlockset(f.Yellow_hello, f.Red_smallhook, f.Blue_bighook, f.Green_L),
		8:  NewBlockset(f.Green_bighook, f.Red_bighook, f.Green_T, f.Red_smallhook),
		9:  NewBlockset(f.Blue_flash, f.Red_stool, f.Blue_v, f.Green_flash),
		10: NewBlockset(f.Green_L, f.Yellow_hello, f.Red_smallhook, f.Red_stool)}
	cards = append(cards, createDifficultCard(34, topShape, bottomShape, blockNums, f))

	// B35
	topShape = NewArray2dFromData([][]int8{{0, 0, 0}, {-1, 0, 0}, {0, 0, 0}, {0, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, 0}, {0, 0, 0}, {-1, -1, 0}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Blue_bighook, f.Red_stool, f.Green_flash, f.Green_bighook),
		2:  NewBlockset(f.Yellow_hello, f.Green_flash, f.Red_stool, f.Yellow_bighook),
		3:  NewBlockset(f.Red_stool, f.Green_bighook, f.Blue_lighter, f.Yellow_hello),
		4:  NewBlockset(f.Yellow_bighook, f.Green_bighook, f.Yellow_hello, f.Red_stool),
		5:  NewBlockset(f.Yellow_gate, f.Red_stool, f.Blue_lighter, f.Blue_bighook),
		6:  NewBlockset(f.Green_flash, f.Blue_v, f.Red_bighook, f.Green_bighook),
		7:  NewBlockset(f.Green_L, f.Red_stool, f.Red_bighook, f.Red_flash),
		8:  NewBlockset(f.Blue_lighter, f.Red_stool, f.Green_L, f.Red_smallhook),
		9:  NewBlockset(f.Green_flash, f.Yellow_smallhook, f.Red_bighook, f.Green_L),
		10: NewBlockset(f.Green_L, f.Red_smallhook, f.Blue_lighter, f.Blue_flash)}
	cards = append(cards, createDifficultCard(35, topShape, bottomShape, blockNums, f))

	// B36
	topShape = NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {0, 0, 0}, {0, 0, -1}})
	bottomShape = NewArray2dFromData([][]int8{{0, 0, -1}, {0, 0, 0}, {0, 0, 0}, {-1, 0, -1}})
	blockNums = map[int]*Blockset{
		1:  NewBlockset(f.Yellow_bighook, f.Yellow_gate, f.Yellow_hello, f.Red_stool),
		2:  NewBlockset(f.Blue_bighook, f.Red_stool, f.Yellow_gate, f.Blue_flash),
		3:  NewBlockset(f.Green_bighook, f.Blue_bighook, f.Yellow_hello, f.Blue_flash),
		4:  NewBlockset(f.Yellow_gate, f.Blue_bighook, f.Yellow_hello, f.Green_flash),
		5:  NewBlockset(f.Yellow_hello, f.Red_stool, f.Yellow_bighook, f.Green_bighook),
		6:  NewBlockset(f.Green_L, f.Red_smallhook, f.Blue_bighook, f.Red_stool),
		7:  NewBlockset(f.Yellow_smallhook, f.Green_L, f.Red_bighook, f.Green_bighook),
		8:  NewBlockset(f.Green_flash, f.Yellow_smallhook, f.Green_L, f.Yellow_hello),
		9:  NewBlockset(f.Green_L, f.Red_smallhook, f.Red_stool, f.Yellow_hello),
		10: NewBlockset(f.Blue_v, f.Yellow_hello, f.Green_flash, f.Blue_flash)}
	cards = append(cards, createDifficultCard(36, topShape, bottomShape, blockNums, f))

	return cards
}

/*
*******************************************************************

This file contains all block-related types and functions:

- Block struct: defines a block
- MakeBlockXX functions: create specific block types

*******************************************************************
*/
package main

import (
	"fmt"
)

// BlockColor is the color of a block in the original game
type BlockColor int8

const (
	Blue BlockColor = iota
	Red
	Yellow
	Green
)

// Returns a string representation for the BlockColor enum
func (s BlockColor) String() string {
	switch s {
	case Blue:
		return "Blue"
	case Red:
		return "Red"
	case Yellow:
		return "Yellow"
	case Green:
		return "Green"
	}
	return "Unknown"
}

// block Array3D can only have the following values:
const IS_BLOCK int8 = 1
const IS_NOT_BLOCK int8 = 0

// Block represents a single Ubongo block including all
// of it's possible rotations in space
type Block struct {
	// Color of the block as in the original game
	Color BlockColor

	// Number of the block, 1-16 for the blocks of the original game
	Number int

	// Shapes is an array of all rotations of the block.
	// Dimensions mean:
	// - 1st index: shape enumeration, base shape is the first, the rest are rotations
	// - 2nd index: x-dimension (horizontal to the right)
	// - 3rd index: y-dimension (up in the 2D-base plane)
	// - 4th index: z-dimension (up into the 3rd dimension)
	Shapes []*Array3d

	// Volume is the number of unit cubes the block consists of.
	// All blocks of the original game consist of 3, 4 or 5 unit cubes
	Volume int
}

// Returns a string representation of the block
func (b Block) String() string {
	return fmt.Sprintf("Block %d: %s (volume %d, %d orientations)",
		b.Number, b.Color, b.Volume, len(b.Shapes))
}

// ****************************************************************************
// Factory functions for the 16 block types of the original game
// ****************************************************************************

type BlockFactory struct {
	blocks         map[int]*Block
	MinBlockNumber int
	MaxBlockNumber int
}

func NewBlockFactory() *BlockFactory {
	return &BlockFactory{
		blocks:         map[int]*Block{},
		MinBlockNumber: 1,
		MaxBlockNumber: 16}
}

// Returns the block with the given number
func (f *BlockFactory) Get(blockNumber int) *Block {
	if blockNumber < f.MinBlockNumber || blockNumber > f.MaxBlockNumber {
		return nil
	}
	if block, ok := f.blocks[blockNumber]; ok {
		return block
	} else {
		funcs := map[int]BlockFactoryFunc{
			1:  NewBlock1,
			2:  NewBlock2,
			3:  NewBlock3,
			4:  NewBlock4,
			5:  NewBlock5,
			6:  NewBlock6,
			7:  NewBlock7,
			8:  NewBlock8,
			9:  NewBlock9,
			10: NewBlock10,
			11: NewBlock11,
			12: NewBlock12,
			13: NewBlock13,
			14: NewBlock14,
			15: NewBlock15,
			16: NewBlock16}

		if blockFunc, ok := funcs[blockNumber]; ok {
			block = blockFunc()
			f.blocks[blockNumber] = block
			return block
		} else {
			return nil
		}
	}
}

// Returns an array with all blocks
func (f *BlockFactory) GetAll() []*Block {
	a := make([]*Block, 0)
	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		if block := f.Get(i); block != nil {
			a = append(a, block)
		}
	}
	return a
}

// declare a global map that can be used to create blocks by its number
var NewBlock map[int]BlockFactoryFunc = map[int]BlockFactoryFunc{
	8: NewBlock8}

// A function that creates a block
type BlockFactoryFunc func() *Block

func (base *Array3d) CreateRotations() []*Array3d {
	arr := make([]*Array3d, 0)

	addIfNotInList := func(lst []*Array3d, el *Array3d) []*Array3d {
		for _, a := range lst {
			if a.IsEqual(el) {
				return lst
			}
		}
		return append(lst, el)
	}

	// the following code generates all possible rotations about 90° along the x, y, z axis for an
	// object in space, in the general case. Some rotations might be identical due to symmetries
	// of the object and will be eliminated
	arr = append(arr, base)
	arr = addIfNotInList(arr, base.RotateZ())
	arr = addIfNotInList(arr, base.RotateZ2())
	arr = addIfNotInList(arr, base.RotateZ3())

	arr = addIfNotInList(arr, base.RotateX())
	arr = addIfNotInList(arr, base.RotateX().RotateZ())
	arr = addIfNotInList(arr, base.RotateX().RotateZ2())
	arr = addIfNotInList(arr, base.RotateX().RotateZ3())

	arr = addIfNotInList(arr, base.RotateX2())
	arr = addIfNotInList(arr, base.RotateX2().RotateZ())
	arr = addIfNotInList(arr, base.RotateX2().RotateZ2())
	arr = addIfNotInList(arr, base.RotateX2().RotateZ3())

	arr = addIfNotInList(arr, base.RotateX3())
	arr = addIfNotInList(arr, base.RotateX3().RotateZ())
	arr = addIfNotInList(arr, base.RotateX3().RotateZ2())
	arr = addIfNotInList(arr, base.RotateX3().RotateZ3())

	arr = addIfNotInList(arr, base.RotateY())
	arr = addIfNotInList(arr, base.RotateY().RotateZ())
	arr = addIfNotInList(arr, base.RotateY().RotateZ2())
	arr = addIfNotInList(arr, base.RotateY().RotateZ3())

	arr = addIfNotInList(arr, base.RotateY3())
	arr = addIfNotInList(arr, base.RotateY3().RotateZ())
	arr = addIfNotInList(arr, base.RotateY3().RotateZ2())
	arr = addIfNotInList(arr, base.RotateY3().RotateZ3())

	return arr
}

func NewBlock1() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {1, 0}, {0, 0}}})
	return &Block{
		Number: 1,
		Color:  Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock2() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 0}, {1, 0}, {1, 1}}, {{1, 0}, {0, 0}, {0, 0}}})
	return &Block{
		Number: 2,
		Color:  Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock3() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 0}, {0, 0}}, {{1, 1}, {0, 1}}})
	return &Block{
		Number: 3,
		Color:  Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock4() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}, {1, 1}}})
	return &Block{
		Number: 4,
		Color:  Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock5() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {0, 0}, {0, 0}}, {{1, 0}, {1, 0}, {1, 0}}})
	return &Block{
		Number: 5,
		Color:  Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock6() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 0}, {1, 1}, {0, 1}}, {{1, 0}, {0, 0}, {0, 0}}})
	return &Block{
		Number: 6,
		Color:  Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock7() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1}, {1}, {1}}, {{0}, {1}, {1}}})
	return &Block{
		Number: 7,
		Color:  Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock8() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{0, 1}, {1, 1}}})
	return &Block{
		Number: 8,
		Color:  Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock9() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}}, {{1, 0}, {1, 0}}})
	return &Block{
		Number: 9,
		Color:  Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock10() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}}, {{0, 0}, {1, 0}}})
	return &Block{
		Number: 10,
		Color:  Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock11() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {0, 0}, {1, 0}}})
	return &Block{
		Number: 11,
		Color:  Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock12() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1}, {1}, {0}}, {{0}, {1}, {1}}})
	return &Block{
		Number: 12,
		Color:  Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock13() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}, {0, 0}}, {{0, 0}, {1, 0}, {1, 0}}})
	return &Block{
		Number: 13,
		Color:  Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock14() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 0}, {1, 0}, {1, 0}}, {{1, 1}, {0, 0}, {0, 0}}})
	return &Block{
		Number: 14,
		Color:  Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock15() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1}, {1}, {1}}, {{0}, {1}, {0}}})
	return &Block{
		Number: 15,
		Color:  Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock16() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1}, {1}, {1}}, {{1}, {0}, {0}}})
	return &Block{
		Number: 16,
		Color:  Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

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

// Block represents a single Ubongo block including all
// of it's possible rotations in space
type Block struct {
	// Color of the block as in the original game
	Color BlockColor

	// Number of the block, 1-16 for the blocks of the original game
	Number int

	// Name is some human-readable representation of the block shape
	Name string

	// Shapes is an array of all rotations of the block.
	// Dimensions mean:
	// - 1st index: shape enumeration, base shape is the first, the rest are rotations
	// - 2nd index: x-dimension (horizontal to the right)
	// - 3rd index: y-dimension (up in the 2D-base plane)
	// - 4th index: z-dimension (up into the 3rd dimension)
	Shapes []Array3d

	// NumShapes is the total number of all rotations including the base shape
	// i.e. the lenght of the first dimension of Shapes
	// This can be max 24 theoretically, but is usually less due to symmetries
	NumShapes int

	// Volume is the number of unit cubes the block consists of.
	// All blocks of the original game consist of 3, 4 or 5 unit cubes
	Volume int
}

// Returns a string representation of the block
func (b Block) String() string {
	return fmt.Sprintf("Block %d: %s %s (volume %d, %d orientations)",
		b.Number, b.Color, b.Name, b.Volume, len(b.Shapes))
}

// ****************************************************************************
// Factory functions for the 16 block types of the original game
// ****************************************************************************

type BlockFactory struct {
	blocks map[int]*Block
}

func NewBlockFactory() *BlockFactory {
	return &BlockFactory{blocks: map[int]*Block{}}
}

// Returns the block with the given number
func (f *BlockFactory) Get(blockNumber int) *Block {
	if block, ok := f.blocks[blockNumber]; ok {
		return block
	} else {
		funcs := map[int]BlockFactoryFunc{
			8: NewBlock8}

		block = funcs[blockNumber]()
		f.blocks[blockNumber] = block
		return block
	}
}

// Returns an array with all blocks
func (f *BlockFactory) GetAll() []*Block {
	a := make([]*Block, 1)
	for i := 8; i <= 8; i++ {
		a[0] = f.Get(i)
	}
	return a
}

// declare a global map that can be used to create blocks by its number
var NewBlock map[int]BlockFactoryFunc = map[int]BlockFactoryFunc{
	8: NewBlock8}

// A function that creates a block
type BlockFactoryFunc func() *Block

// NewBlock8 creates the blue small angle-shaped block
func NewBlock8() *Block {
	var b *Block = new(Block)

	b.Number = 8
	b.Color = Blue
	b.Name = "Angle"

	b.Shapes = []Array3d{
		{{{0, 1}, {1, 1}}},
		{{{1, 0}, {1, 1}}},
		{{{1, 1}, {1, 0}}},
		{{{1, 1}, {0, 1}}},
		{{{1}, {1}}, {{0}, {1}}},
		{{{1, 1}}, {{1, 0}}},
		{{{1}, {1}}, {{1}, {0}}},
		{{{1, 1}}, {{0, 1}}},
		{{{1, 0}}, {{1, 1}}},
		{{{1}, {0}}, {{1}, {1}}},
		{{{0, 1}}, {{1, 1}}},
		{{{0}, {1}}, {{1}, {1}}}}

	b.Volume = CountValues3D(b.Shapes[0], 1)

	return b
}

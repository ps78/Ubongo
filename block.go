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

// BlockShape is a 3-dimensional array representing the shape of a block, where
// 1 inidicates the presence of a unit cube and 0 absence of one
type BlockShape = [][][]int8

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
	Shapes []BlockShape

	// NumOrientations is the total number of all rotations including the base shape
	// i.e. the lenght of the first dimension of Shapes
	// This can be max 24 theoretically, but is usually less due to symmetries
	NumOrientations int

	// Volume is the number of unit cubes the block consists of.
	// All blocks of the original game consist of 3, 4 or 5 unit cubes
	Volume int
}

// Returns a string representation of the block
func (b Block) String() string {
	return fmt.Sprintf("Block %d: %s %s (volume %d, %d orientations)",
		b.Number, b.Color, b.Name, b.Volume, len(b.Shapes))
}

// GetBlockVolume calculates the volume of the given block shape in unit cubes
func GetBlockVolume(shape BlockShape) int {
	var volume int = 0
	for x, b := range shape {
		for y, c := range b {
			for z := range c {
				if shape[x][y][z] == 1 {
					volume++
				}
			}
		}
	}
	return volume
}

// GetBoundBoxSize returns the dimensions of the given volume
// which correspond to the size of the bounding box
func GetBoundingBoxFromBlockShape(shape BlockShape) BoundingBox {
	xdim := len(shape)
	ydim := len(shape[0])
	zdim := len(shape[0][0])
	return BoundingBox{xdim, ydim, zdim}
}

// ****************************************************************************
// Factory functions for the 16 block types of the original game
// ****************************************************************************

// A function that creates a block
type BlockFactoryFunc func() *Block

// MakeBlock08 creates the blue small angle-shaped block
func MakeBlock08() *Block {
	var b *Block = new(Block)

	b.Number = 8
	b.Color = Blue
	b.Name = "Angle"

	b.Shapes = []BlockShape{
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

	b.Volume = GetBlockVolume(b.Shapes[0])

	return b
}

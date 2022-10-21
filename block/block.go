/*
*******************************************************************

This file contains all block-related types and functions:

- Block struct: defines a block
- functions that create the 16 blocks of the game

*******************************************************************
*/
package block

import (
	"fmt"
	"image/color"
	"ubongo/base/array3d"
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

func (c BlockColor) ToRGBA() color.RGBA {
	switch c {
	case Green:
		return color.RGBA{20, 255, 50, 0}
	case Blue:
		return color.RGBA{0, 20, 150, 0}
	case Red:
		return color.RGBA{200, 20, 20, 0}
	case Yellow:
		return color.RGBA{200, 200, 0, 0}
	}
	return color.RGBA{255, 255, 255, 0}
}

// Block represents a single Ubongo block including all
// of it's possible rotations in space
type Block struct {
	// Color of the block as in the original game
	Color BlockColor

	// Number of the block, 1-16 for the blocks of the original game
	Number int

	// Easily recognizable name of the block, unique together with the color
	Name string

	// Shapes is an array of all rotations of the block
	Shapes []*array3d.A

	// Volume is the number of unit cubes the block consists of.
	// All blocks of the original game consist of 3, 4 or 5 unit cubes
	Volume int
}

// Returns a string representation of the block
func (b *Block) String() string {
	return fmt.Sprintf("Block %d: %s %s (volume %d, %d orientations)",
		b.Number, b.Color, b.Name, b.Volume, len(b.Shapes))
}

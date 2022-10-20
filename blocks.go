/*
*******************************************************************

This file contains all block-related types and functions:

- Block struct: defines a block
- functions that create the 16 blocks of the game

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

	// Easily recognizable name of the block, unique together with the color
	Name string

	// Shapes is an array of all rotations of the block
	Shapes []*Array3d

	// Volume is the number of unit cubes the block consists of.
	// All blocks of the original game consist of 3, 4 or 5 unit cubes
	Volume int
}

// Returns a string representation of the block
func (b *Block) String() string {
	return fmt.Sprintf("Block %d: %s %s (volume %d, %d orientations)",
		b.Number, b.Color, b.Name, b.Volume, len(b.Shapes))
}

// Creates all 90° rotations of the base 3d array along the x, y and z axis
// A maximum of 24 arrays are returned, but identical rotations are removed,
// hence the number can be smaller (depending on symmetries of the base array)
func (base *Array3d) createRotations() []*Array3d {
	arr := make([]*Array3d, 0)

	// helper function that adds el to lst if it is not already in lst
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

/*******************************************************************************
 * Block-creator functions for the 16 blocks of the game
 *******************************************************************************/

func NewBlock1() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {1, 0}, {0, 0}}})
	return &Block{
		Number: 1,
		Name:   "hello",
		Color:  Yellow,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock2() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 0}, {1, 0}, {1, 1}}, {{1, 0}, {0, 0}, {0, 0}}})
	return &Block{
		Number: 2,
		Name:   "big hook",
		Color:  Yellow,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock3() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 0}, {0, 0}}, {{1, 1}, {0, 1}}})
	return &Block{
		Number: 3,
		Name:   "small hook",
		Color:  Yellow,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock4() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}, {1, 1}}})
	return &Block{
		Number: 4,
		Name:   "gate",
		Color:  Yellow,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock5() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {0, 0}, {0, 0}}, {{1, 0}, {1, 0}, {1, 0}}})
	return &Block{
		Number: 5,
		Name:   "big hook",
		Color:  Blue,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock6() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 0}, {1, 1}, {0, 1}}, {{1, 0}, {0, 0}, {0, 0}}})
	return &Block{
		Number: 6,
		Name:   "flash",
		Color:  Blue,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock7() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1}, {1}, {1}}, {{0}, {1}, {1}}})
	return &Block{
		Number: 7,
		Name:   "lighter",
		Color:  Blue,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock8() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{0, 1}, {1, 1}}})
	return &Block{
		Number: 8,
		Name:   "v",
		Color:  Blue,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock9() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}}, {{1, 0}, {1, 0}}})
	return &Block{
		Number: 9,
		Name:   "stool",
		Color:  Red,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock10() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}}, {{0, 0}, {1, 0}}})
	return &Block{
		Number: 10,
		Name:   "small hook",
		Color:  Red,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock11() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {0, 0}, {1, 0}}})
	return &Block{
		Number: 11,
		Name:   "big hook",
		Color:  Red,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock12() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1}, {1}, {0}}, {{0}, {1}, {1}}})
	return &Block{
		Number: 12,
		Name:   "flash",
		Color:  Red,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock13() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 1}, {1, 0}, {0, 0}}, {{0, 0}, {1, 0}, {1, 0}}})
	return &Block{
		Number: 13,
		Name:   "flash",
		Color:  Green,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock14() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1, 0}, {1, 0}, {1, 0}}, {{1, 1}, {0, 0}, {0, 0}}})
	return &Block{
		Number: 14,
		Name:   "big hook",
		Color:  Green,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock15() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1}, {1}, {1}}, {{0}, {1}, {0}}})
	return &Block{
		Number: 15,
		Name:   "T",
		Color:  Green,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

func NewBlock16() *Block {
	baseShape := NewArray3dFromData([][][]int8{{{1}, {1}, {1}}, {{1}, {0}, {0}}})
	return &Block{
		Number: 16,
		Name:   "L",
		Color:  Green,
		Shapes: baseShape.createRotations(),
		Volume: baseShape.Count(1)}
}

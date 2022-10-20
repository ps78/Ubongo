package main

import (
	"fmt"
	"ubongo/utils"
)

// GameSolution represents one solution to a specific game
type GameSolution struct {
	// Array of block references, as many as blocks are required by the game
	Blocks []*Block

	// For each block of the Blocks array, the ShapeIndex array references the
	// specific shape of the block to be used. I.e.
	ShapeIndex []int

	// Shifts defines the translation of the Shape with the identical index
	// relative to the origin, in 'array-units', i.e. 1 equals one mini-cube
	Shifts []utils.Vector
}

// Creates an instance of GameSolution
func NewGameSolution(blocks []*Block, shapeIndex []int, shifts []utils.Vector) *GameSolution {
	sol := GameSolution{}

	sol.Blocks = make([]*Block, len(blocks))
	copy(sol.Blocks, blocks)

	sol.ShapeIndex = make([]int, len(shapeIndex))
	copy(sol.ShapeIndex, shapeIndex)

	sol.Shifts = make([]utils.Vector, len(shifts))
	copy(sol.Shifts, shifts)

	return &sol
}

// Returns a multi-line string representing the GameSolution
func (gs *GameSolution) String() string {
	result := "GameSolution\n\t"
	for i := 0; i < len(gs.Blocks); i++ {
		result += fmt.Sprintf("<#%s %s (v%d) Shape #%d %s Shift %s>",
			gs.Blocks[i].Color, gs.Blocks[i].Name, gs.Blocks[i].Volume, gs.ShapeIndex[i], gs.Blocks[i].Shapes[gs.ShapeIndex[i]], gs.Shifts[i])
		if i < len(gs.Blocks)-1 {
			result += "\n\t"
		}
	}
	return result
}

// GetCenterOfGravity calculates the center of gravity of the given solution
func (gs *GameSolution) GetCenterOfGravity() utils.Vectorf {
	c := utils.Vectorf{}
	var totalVolume float64
	for i, shapeIdx := range gs.ShapeIndex {
		block := gs.Blocks[i]
		blockVolume := float64(block.Volume)
		totalVolume += blockVolume
		c = c.Add(block.Shapes[shapeIdx].GetCenterOfGravity().Add(gs.Shifts[i].AsVectorf()).Mult(blockVolume))
	}
	return c.Div(totalVolume)
}

// Returns the bounding box of the whole game solution
func (gs *GameSolution) GetBoundingBox() utils.Vector {
	bb := utils.Vector{}
	for i, sIdx := range gs.ShapeIndex {
		shape := gs.Blocks[i].Shapes[sIdx]
		dim := shape.GetBoundingBox().Add(gs.Shifts[i])
		for i := 0; i < 3; i++ {
			if dim[i] > bb[i] {
				bb[i] = dim[i]
			}
		}
	}
	return bb
}

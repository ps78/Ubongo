package main

import "fmt"

type GameSolution struct {
	Blocks []*Block
	Shapes []*Array3d
	Shifts []Vector
}

// Creates an instance of GameSolution
func NewGameSolution(blocks []*Block, shapes []*Array3d, shifts []Vector) *GameSolution {
	sol := GameSolution{}

	sol.Blocks = make([]*Block, len(blocks))
	copy(sol.Blocks, blocks)

	sol.Shapes = make([]*Array3d, len(shapes))
	copy(sol.Shapes, shapes)

	sol.Shifts = make([]Vector, len(shifts))
	copy(sol.Shifts, shifts)

	return &sol
}

// Returns a multi-line string representing the GameSolution
func (gs *GameSolution) String() string {
	result := "GameSolution\n\t"
	for i := 0; i < len(gs.Blocks); i++ {
		_, shapeIdx := FindArray3d(gs.Blocks[i].Shapes, gs.Shapes[i])
		result += fmt.Sprintf("<#%d (v%d) Shape #%d %s Shift %s>", gs.Blocks[i].Number, gs.Blocks[i].Volume, shapeIdx, gs.Shapes[i], gs.Shifts[i])
		if i < len(gs.Blocks)-1 {
			result += "\n\t"
		}
	}
	return result
}

func (gs *GameSolution) GetCenterOfGravity() Vectorf {
	c := Vectorf{}
	var totalVolume float64
	for i, shape := range gs.Shapes {
		blockVolume := float64(gs.Blocks[i].Volume)
		totalVolume += blockVolume
		c = c.Add(shape.GetCenterOfGravity().Add(gs.Shifts[i].Float64()).Mult(blockVolume))
	}
	return c.Div(totalVolume)
}

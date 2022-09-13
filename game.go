package main

import (
	"fmt"
)

type Game struct {
	Shape  *Array2d
	Volume *Array3d
}

// the following constants define the allowed values of the shape and volume array values of a game
const EMPTY int8 = 0
const OCCUPIED int8 = 1
const OUTSIDE int8 = -1

type GameSolution struct {
	Blocks []*Block
	Shapes []*Array3d
	Shifts []Vector
}

// Creates a new game, initialized with the given shape and height and an empty volume
func NewGame(shape *Array2d, height int) *Game {
	return &Game{Shape: shape.Clone(), Volume: shape.Extrude(height)}
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

// Returns a nicely formatted string representation of the game
func (g *Game) String() string {
	return fmt.Sprintf("Game (area %d, volume %d, empty %d)",
		g.Shape.Count(0), g.Shape.Count(0)*g.Volume.DimZ, g.Volume.Count(0))
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

// Removes all blocks from a game
func (g *Game) Clear() {
	for x := 0; x < g.Volume.DimX; x++ {
		for y := 0; y < g.Volume.DimY; y++ {
			for z := 0; z < g.Volume.DimZ; z++ {
				g.Volume.Set(x, y, z, g.Shape.Get(x, y))
			}
		}
	}
}

// Creates a copy of the game
func (g *Game) Clone() *Game {
	return &Game{Shape: g.Shape.Clone(), Volume: g.Volume.Clone()}
}

// Tries to add the given block to the game volume
// returns true if successful, false if not
func (g *Game) TryAddBlock(block *Array3d, pos Vector) bool {
	// check overall dimensions
	if pos[0]+block.DimX > g.Volume.DimX ||
		pos[1]+block.DimY > g.Volume.DimY ||
		pos[2]+block.DimZ > g.Volume.DimZ {
		return false
	}

	// step 1: test if it is possible to add block
	for x := 0; x < block.DimX; x++ {
		for y := 0; y < block.DimY; y++ {
			for z := 0; z < block.DimZ; z++ {
				if block.Get(x, y, z) == IS_BLOCK && g.Volume.Get(x+pos[0], y+pos[1], z+pos[2]) != EMPTY {
					return false
				}
			}
		}
	}

	// step 2: actually add the block, the expensive step is the cloning of the object
	v := g.Volume.Clone()
	for x := 0; x < block.DimX; x++ {
		for y := 0; y < block.DimY; y++ {
			for z := 0; z < block.DimZ; z++ {
				if block.Get(x, y, z) == IS_BLOCK {
					v.Set(x+pos[0], y+pos[1], z+pos[2], OCCUPIED)
				}
			}
		}
	}

	// replace the game's volume with the new one containing the block
	g.Volume = v
	return true
}

// Removes the block at the given position from the volume
// This does not check if the block is actually present and
// simply sets all values from 1 to 0
func (g *Game) RemoveBlock(block *Array3d, pos Vector) bool {
	// check overall dimensions
	if pos[0]+block.DimX > g.Volume.DimX ||
		pos[1]+block.DimY > g.Volume.DimY ||
		pos[2]+block.DimZ > g.Volume.DimZ {
		return false
	}

	for x := 0; x < block.DimX; x++ {
		for y := 0; y < block.DimY; y++ {
			for z := 0; z < block.DimZ; z++ {
				// only run following test if the block-cube is solid at the current position
				if block.Get(x, y, z) == IS_BLOCK {
					// if space is occupied by block, set it to 0
					if g.Volume.Get(x+pos[0], y+pos[1], z+pos[2]) == OCCUPIED {
						g.Volume.Set(x+pos[0], y+pos[1], z+pos[2], EMPTY) // mark space as empty
					}
				}
			}
		}
	}
	return true
}

// Finds all solutino for a given game using the set of blocks provided
func (g *Game) Solve(blocks []*Block) []*GameSolution {
	// check the sum of the block volumes, it must match the empty volume of the game to yield a solution
	sum := 0
	for _, b := range blocks {
		sum += b.Volume
	}
	if g.Volume.Count(EMPTY) != sum {
		return []*GameSolution{}
	}

	// working arrays for the recursive solver:
	solutions := make([]*GameSolution, 0)
	shapes := make([]*Array3d, 0)
	shifts := make([]Vector, 0)

	g.recursiveSolver(blocks, 0, &shapes, &shifts, &solutions)

	return solutions
}

var recursiveSolverCount int64 = 0

// Recursive function called by Solve, don't call directly
func (g *Game) recursiveSolver(blocks []*Block, startAtBlockIdx int, shapes *[]*Array3d, shifts *[]Vector, solutions *[]*GameSolution) {
	recursiveSolverCount++
	gameBox := g.Volume.GetBoundingBox()

	// cycle through the remaining blocks
	for blockIdx := startAtBlockIdx; blockIdx < len(blocks); blockIdx++ {
		block := blocks[blockIdx]

		for _, shape := range block.Shapes {
			*shapes = append(*shapes, shape)

			shiftVectors := gameBox.GetShiftVectors(shape.GetBoundingBox())
			for _, shift := range shiftVectors {
				if ok := g.TryAddBlock(shape, shift); ok {

					*shifts = append(*shifts, shift)

					// if this was the last block, stop recursion
					if blockIdx == len(blocks)-1 {
						// check if we have a solution
						if g.Volume.Count(0) == 0 {
							*solutions = append(*solutions, NewGameSolution(blocks, *shapes, *shifts))
						}
						// if it wasn't the last block, continue recursion
					} else {
						g.recursiveSolver(blocks, blockIdx+1, shapes, shifts, solutions)
					}

					*shifts = (*shifts)[:len(*shifts)-1]

					g.RemoveBlock(shape, shift)
				}
			} // end loop over shifts

			*shapes = (*shapes)[:len(*shapes)-1]

		} // end loop over shapes

	} // end loop over blocks
}

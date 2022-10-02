package main

import (
	"fmt"
)

type Game struct {
	Shape  *Array2d
	Volume *Array3d
	Blocks *Blockset
}

// value of an area or volume representing empty space (but part of shape/volume)
const EMPTY int8 = 0

// value of a volume representing a space in a volume occupied by a block
const OCCUPIED int8 = 1

// value of an area or volume representing a space outside the area/volume
const OUTSIDE int8 = -1

// The number of blocks of each type in the original Ubongo game
// map[BlockNumer]Count
var UbongoBlockSet map[int]int = map[int]int{
	1:  2, // yellow hello
	2:  2, // yellow bighook
	3:  3, // yellow smallhook
	4:  2, // yellow gate
	5:  2, // blue bighook
	6:  2, // blue flash
	7:  3, // blue lighter
	8:  4, // blue v
	9:  3, // red stool
	10: 3, // red smallhook
	11: 2, // red bighook
	12: 2, // red flash
	13: 2, // green flash
	14: 2, // green bighook
	15: 2, // green T
	16: 4, // green L
}

// Creates a new game, initialized with the given shape and height and an empty volume
func NewGame(p *Problem) *Game {
	return &Game{
		Shape:  p.Shape.Clone(),
		Volume: p.Shape.Extrude(p.Height),
		Blocks: p.Blocks.Clone()}
}

// Returns a nicely formatted string representation of the game
func (g *Game) String() string {
	return fmt.Sprintf("Game (area %d, volume %d, empty %d)",
		g.Shape.Count(0), g.Shape.Count(0)*g.Volume.DimZ, g.Volume.Count(0))
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
	return &Game{
		Shape:  g.Shape.Clone(),
		Volume: g.Volume.Clone(),
		Blocks: g.Blocks.Clone()}
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
func (g *Game) Solve() []*GameSolution {
	// check the sum of the block volumes, it must match the empty volume of the game to yield a solution
	if g.Volume.Count(EMPTY) != g.Blocks.Volume() {
		return []*GameSolution{}
	}

	// working arrays for the recursive solver:
	solutions := make([]*GameSolution, 0)
	shapes := make([]*Array3d, 0)
	shifts := make([]Vector, 0)

	g.recursiveSolver(0, &shapes, &shifts, &solutions)

	return solutions
}

// Recursive function called by Solve, don't call directly
func (g *Game) recursiveSolver(blockIdx int, shapes *[]*Array3d, shifts *[]Vector, solutions *[]*GameSolution) {
	gameBox := g.Volume.GetBoundingBox()

	block := g.Blocks.At(blockIdx)

	for _, shape := range block.Shapes {
		*shapes = append(*shapes, shape)

		shiftVectors := gameBox.GetShiftVectors(shape.GetBoundingBox())
		for _, shift := range shiftVectors {
			if ok := g.TryAddBlock(shape, shift); ok {

				*shifts = append(*shifts, shift)

				// if this was the last block, stop recursion
				if blockIdx == g.Blocks.Count-1 {
					// check if we have a solution
					if g.Volume.Count(0) == 0 {
						*solutions = append(*solutions, NewGameSolution(g.Blocks.AsSlice(), *shapes, *shifts))
					}
					// if it wasn't the last block, continue recursion
				} else {
					g.recursiveSolver(blockIdx+1, shapes, shifts, solutions)
				}

				*shifts = (*shifts)[:len(*shifts)-1]

				g.RemoveBlock(shape, shift)
			}
		} // end loop over shifts

		*shapes = (*shapes)[:len(*shapes)-1]

	} // end loop over shapes
}

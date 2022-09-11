package main

import (
	"fmt"
)

type Game struct {
	Shape  *Array2d
	Volume *Array3d
}

type GameSolution struct {
	Blocks []*Block
	Shapes []*Array3d
	Shifts []Vector
}

// Creates a new game, initialized with the given shape and height and an empty volume
func NewGame(shape *Array2d, height int) *Game {
	g := new(Game)
	g.Shape = shape.Clone()
	g.Volume = shape.Extrude(height)
	return g
}

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
func (g *Game) TryAddBlock(block *Array3d, pos Vector) (bool, *Game) {
	// check overall dimensions
	if pos[0]+block.DimX > g.Volume.DimX ||
		pos[1]+block.DimY > g.Volume.DimY ||
		pos[2]+block.DimZ > g.Volume.DimZ {
		return false, g
	}

	v := g.Volume.Clone()
	for x := 0; x < block.DimX; x++ {
		for y := 0; y < block.DimY; y++ {
			for z := 0; z < block.DimZ; z++ {
				// only run following test if the block-cube is solid at the current position
				if block.Get(x, y, z) == 1 {
					// if space is part of volume and empty -> ok
					if v.Get(x+pos[0], y+pos[1], z+pos[2]) == 0 {
						v.Set(x+pos[0], y+pos[1], z+pos[2], 1) // mark space as occupied
					} else {
						return false, g // otherwise abort
					}
				}
			}
		}
	}

	// replace the game's volume with the new one containing the block
	return true, &Game{Shape: g.Shape, Volume: v}
}

func (g *Game) Solve(blocks []*Block) {
	sum := 0
	for _, b := range blocks {
		sum += b.Volume
	}
	if g.Volume.Count(0) != sum {
		fmt.Printf("Game has no solution, volume of blocks does not match volume of game")
		return
	}

	solutions := make([]GameSolution, 0)
	shapes := make([]*Array3d, 0)
	shifts := make([]Vector, 0)

	g.recursiveSolver(blocks, 0, &shapes, &shifts, &solutions)

	fmt.Printf("Found %d solutions\n", len(solutions))
	for _, sol := range solutions {
		fmt.Println(sol.String())
	}

	fmt.Printf("Block #%d\n", blocks[0].Number)
	for i, shape := range blocks[0].Shapes {
		fmt.Printf("%d - %s\n", i, shape)
	}
}

func (g *Game) recursiveSolver(blocks []*Block, currentBlockIdx int, shapes *[]*Array3d, shifts *[]Vector, solutions *[]GameSolution) {
	gameBox := g.Volume.GetBoundingBox()

	for _, block := range blocks[currentBlockIdx:] {

		for _, shape := range block.Shapes {
			sh := gameBox.GetShiftVectors(shape.GetBoundingBox())
			for _, shift := range sh {
				if ok, newGame := g.TryAddBlock(shape, shift); ok {
					// add shape+shift to the stack
					*shapes = append(*shapes, shape)
					*shifts = append(*shifts, shift)

					// if this was the last block, stop
					if currentBlockIdx == len(blocks)-1 {
						// check if we have a solution
						if newGame.Volume.Count(0) == 0 {
							*solutions = append(*solutions, *NewGameSolution(blocks, *shapes, *shifts))
						}
						// if it wasn't the last block, continue recursion
					} else {
						newGame.recursiveSolver(blocks, currentBlockIdx+1, shapes, shifts, solutions)
					}

					// remove shape+shift from the stack
					*shapes = (*shapes)[:len(*shapes)-1]
					*shifts = (*shifts)[:len(*shifts)-1]
				}
			}
		}
	}
}

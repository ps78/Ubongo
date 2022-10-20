package main

import (
	"fmt"
	"ubongo/utils"
	"ubongo/utils/array2d"
)

// Problem represents a single Ubongo problem to solve
type Problem struct {
	// Shape is the 2D shape of the puzzle, first is the index X-direction (horizontal, to the right),
	// the second index is the Y-direction (up)
	Shape *array2d.A // -1=not part of volume, 0=empty, 1=occupied by a block

	// Height of the volume to fill with the blocks. This is always 2 for the original game
	Height int

	// The area of the problem in unit squares
	Area int

	// Bounding box of the problem volume
	BoundingBox utils.Vector

	// Blocks is an array of the blocks to be used to fill the volume
	Blocks *Blockset
}

// Returns a string representation of the problem
func (p Problem) String() string {
	return fmt.Sprintf("Problem: %d blocks, area %d, height %d", p.Blocks.Count, p.Area, p.Height)
}

// Creates a problem instance
func NewProblem(shape *array2d.A, height int, blocks *Blockset) *Problem {
	var p *Problem = new(Problem)

	p.Shape = shape.Clone()
	p.Blocks = blocks.Clone()
	p.Height = height
	p.Area = p.Shape.Count(0)
	p.BoundingBox = utils.Vector{p.Shape.DimX, p.Shape.DimY, p.Height}

	return p
}

// IsEqual returns true of o contains the same data as p
func (p *Problem) IsEqual(o *Problem) bool {
	if o == nil {
		return false
	} else {
		return p.Area == o.Area &&
			p.Height == o.Height &&
			p.BoundingBox == o.BoundingBox &&
			p.Shape.IsEqual(o.Shape) &&
			p.Blocks.IsEqual(o.Blocks)
	}
}

// Clone creates a deep copy of a problem
func (p *Problem) Clone() *Problem {
	var n *Problem = new(Problem)

	n.Shape = p.Shape.Clone()
	n.Blocks = p.Blocks.Clone()
	n.Height = p.Height
	n.Area = p.Area
	n.BoundingBox = p.BoundingBox

	return n
}

// GenerateProblems creates numProblems new problems based on the given
// parameters (height, shape, blockCount)
func GenerateProblems(bf *BlockFactory, shape *array2d.A, height, blockCount, numProblems int) []*Problem {
	multiplier := 5 // we generate more problems than requested, as some might not have a solution
	results := make([]*Problem, 0)

	// generate random blocksets, more than we need, as not all might be solvable
	sets := GenerateBlocksets(bf, shape.Count(EMPTY)*height, blockCount, multiplier*numProblems)

	for i := range sets {
		p := NewProblem(shape, height, sets[i])
		g := NewGame(p)
		solutions := g.Solve()

		if len(solutions) > 0 {
			results = append(results, p)
		}

		if len(results) >= numProblems {
			break
		}
	}
	return results
}

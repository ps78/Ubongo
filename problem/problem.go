// Package problem contains the type P(roblem) and its methods
package problem

import (
	"fmt"
	"ubongo/base/array2d"
	"ubongo/base/vector"
	"ubongo/blockset"
)

// P represents a single Ubongo problem to solve
type P struct {
	// Shape is the 2D shape of the puzzle, first is the index X-direction (horizontal, to the right),
	// the second index is the Y-direction (up)
	Shape *array2d.A // -1=not part of volume, 0=empty, 1=occupied by a block

	// Height of the volume to fill with the blocks. This is always 2 for the original game
	Height int

	// The area of the problem in unit squares
	Area int

	// Bounding box of the problem volume
	BoundingBox vector.V

	// Blocks is an array of the blocks to be used to fill the volume
	Blocks *blockset.S
}

// String returns a string representation of the problem
func (p *P) String() string {
	if p == nil {
		return "(nil)"
	} else {
		return fmt.Sprintf("Problem: %d blocks, area %d, height %d", p.Blocks.Count, p.Area, p.Height)
	}
}

// New creates a problem instance
func New(shape *array2d.A, height int, blocks *blockset.S) *P {
	var p *P = new(P)

	p.Shape = shape.Clone()
	p.Blocks = blocks.Clone()
	p.Height = height
	p.Area = p.Shape.Count(0)
	p.BoundingBox = vector.V{p.Shape.DimX, p.Shape.DimY, p.Height}

	return p
}

// Equals returns true of o contains the same data as p
func (p *P) Equals(o *P) bool {
	if o == nil {
		return false
	} else {
		return p.Area == o.Area &&
			p.Height == o.Height &&
			p.BoundingBox == o.BoundingBox &&
			p.Shape.Equals(o.Shape) &&
			p.Blocks.Equals(o.Blocks)
	}
}

// Clone creates a deep copy of a problem
func (p *P) Clone() *P {
	if p == nil {
		return nil
	} else {
		var n *P = new(P)

		n.Shape = p.Shape.Clone()
		n.Blocks = p.Blocks.Clone()
		n.Height = p.Height
		n.Area = p.Area
		n.BoundingBox = p.BoundingBox

		return n
	}
}

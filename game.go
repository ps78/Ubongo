package main

type Game struct {
	Shape  *Array2d
	Volume *Array3d
}

// Creates a new game, initialized with the given shape and height and an empty volume
func NewGame(shape *Array2d, height int) *Game {
	g := new(Game)
	g.Shape = shape.Clone()
	g.Volume = shape.Extrude(height)
	return g
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

func (g Game) Clone() *Game {
	return &Game{Shape: g.Shape.Clone(), Volume: g.Volume.Clone()}
}

// Tries to add the given block to the game volume
// returns true if successful, false if not
func (g *Game) TryAddBlock(block Array3d, pos Vector) bool {
	// check overall dimensions
	if pos[0]+block.DimX > g.Volume.DimX ||
		pos[1]+block.DimY > g.Volume.DimY ||
		pos[2]+block.DimZ > g.Volume.DimZ {
		return false
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
						return false // otherwise abort
					}
				}
			}
		}
	}

	// replace the game's volume with the new one containing the block
	g.Volume = v
	return true
}

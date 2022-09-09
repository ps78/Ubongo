package main

type Game struct {
	Shape  Array2d
	Xdim   int
	Ydim   int
	Height int
	Volume Array3d
}

// Creates a new game, initialized with the given shape and height and an empty volume
func NewGame(shape Array2d, height int) *Game {
	g := new(Game)
	g.Shape = Copy2DArray(shape)
	g.Height = height
	g.Xdim = len(shape)
	g.Ydim = len(shape[0])
	g.Volume = Extrude2DArray(g.Shape, g.Height)
	return g
}

// Removes all blocks from a game
func (g *Game) Clear() {
	for x := 0; x < g.Xdim; x++ {
		for y := 0; y < g.Ydim; y++ {
			for z := 0; z < g.Height; z++ {
				g.Volume[x][y][z] = g.Shape[x][y]
			}
		}
	}
}

func (g Game) Clone() *Game {
	clone := new(Game)
	clone.Shape = g.Shape // do not copy this array, just copy the reference
	clone.Height = g.Height
	clone.Xdim = g.Xdim
	clone.Ydim = g.Ydim
	clone.Volume = Copy3DArray(g.Volume)
	return clone
}

// Tries to add the given block to the game volume
// returns true if successful, false if not
func (g *Game) TryAddBlock(block Array3d, pos Vector) bool {
	v := Copy3DArray(g.Volume)

	// check overall dimensions
	xdim := len(block)
	ydim := len(block[0])
	zdim := len(block[0][0])
	if pos[0]+xdim > g.Xdim ||
		pos[1]+ydim > g.Ydim ||
		pos[2]+zdim > g.Height {
		return false
	}

	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {
			for z := 0; z < zdim; z++ {
				// only run following test if the block-cube is solid at the current position
				if block[x][y][z] == 1 {
					// if space is part of volume and empty -> ok
					if v[x+pos[0]][y+pos[1]][z+pos[2]] == 0 {
						v[x+pos[0]][y+pos[1]][z+pos[2]] = 1 // mark space as occupied
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

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
	g.Volume = Make3DArray(g.Xdim, g.Ydim, g.Height)
	return g
}

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

func (game Game) AddBlock(block Block) {

}

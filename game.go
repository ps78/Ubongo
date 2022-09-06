package main

type Game struct {
	Shape  ProblemShape
	Height int
	Volume ProblemVolume
}

func (game Game) Clear() {
	// todo
}

func (game Game) Clone() Game {
	// todo
	return game
}

func (game Game) AddBlock(block Block) {

}

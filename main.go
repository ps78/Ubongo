package main

import (
	"fmt"
	"ubongo/base/array2d"
	"ubongo/blockfactory"
	"ubongo/card"
	"ubongo/cardfactory"
	"ubongo/game"
	"ubongo/graphics"
)

func main() {
	var mode int = 0

	switch mode {
	case 0:
		var a *array2d.A = nil
		fmt.Println(a)
	case 1:
		// create block render images:
		graphics.RenderAll(blockfactory.Get().GetAll(), "./results/images", 500, 500)
	case 2:
		// solve all problems and save stats:
		game.CreateSolutionStatistics(cardfactory.Get(), "./results/solutions.csv")
	case 3:
		// Generate Insane problems
		for _, animal := range card.AllAnimals() {
			game.GenerateCardSet(cardfactory.Get(), blockfactory.Get(), animal, card.Easy, card.Insane, 3, 5, fmt.Sprintf("./results/cards/%s.txt", animal))
		}
	case 4:
		// Visualize a solution
		gs := game.New(cardfactory.Get().Get(card.Difficult, 1).Problems[1]).Solve()[0]
		graphics.Visualize(gs)
	}
}

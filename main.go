package main

import (
	"fmt"
	"ubongo/blockfactory"
	"ubongo/card"
	"ubongo/cardfactory"
	"ubongo/game"
	"ubongo/graphics"
)

func main() {
	var mode int = 1

	switch mode {
	case 1:
		// create block render images:
		graphics.RenderAll(blockfactory.GetBlockFactory().GetAll(), "./results/images", 500, 500)
	case 2:
		// solve all problems and save stats:
		game.CreateSolutionStatistics(cardfactory.GetCardFactory(), "./results/solutions.csv")
	case 3:
		// Generate Insane problems
		for _, animal := range card.AllAnimals() {
			game.GenerateCardSet(cardfactory.GetCardFactory(), blockfactory.GetBlockFactory(), animal, card.Easy, card.Insane, 3, 5, fmt.Sprintf("./results/cards/%s.txt", animal))
		}
	case 4:
		// Visualize a solution
		gs := game.NewGame(cardfactory.GetCardFactory().Get(card.Difficult, 1).Problems[1]).Solve()[0]
		graphics.Visualize(gs)
	}
}

package main

import "fmt"

func main() {
	var mode int = 1

	switch mode {
	case 1:
		// create block render images:
		GetBlockFactory().RenderAll("./images", 500, 500)
	case 2:
		// solve all problems and save stats:
		GetCardFactory().CreateSolutionStatistics("solutions.csv")
	case 3:
		// Generate Insane problems
		for _, animal := range []UbongoAnimal{Elephant, Gazelle, Snake, Gnu, Ostrich, Rhino, Giraffe, Zebra, Warthog} {
			GenerateCardSet(GetCardFactory(), GetBlockFactory(), animal, Easy, Insane, 3, 5, fmt.Sprintf("cards/%s.txt", animal))
		}
	case 4:
		// Visualize a solution
		gs := NewGame(GetCardFactory().Get(Difficult, 1).Problems[1]).Solve()[0]
		gs.Visualize()
	}
}

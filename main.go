package main

import (
	//"fmt"
	//"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func updateImage(win fyne.Window, sol *GameSolution, w, h int, rx, ry, rz float64) {
	img := sol.CreateImage(w, h, rx, ry, rz, 0.1)
	win.SetContent(canvas.NewImageFromImage(img))
}

func main() {
	// create block render images:
	//GetBlockFactory().RenderAll("./images", 500, 500)

	// solve all problems and save stats:
	//GetCardFactory().CreateSolutionStatistics("solutions.csv")

	//fc := GetCardFactory()
	//fb := GetBlockFactory()

	// Generate Insane problems
	bf := GetBlockFactory()
	bc := GetCardFactory()
	type key struct {
		Animal     UbongoAnimal
		CardNumber int
		DiceNumber int
	}
	problems := map[key][]*Problem{}
	animal := Elephant
	for _, card := range bc.GetByAnimal(Easy, animal) {
		for diceNumber := 1; diceNumber <= 10; diceNumber++ {
			var shape *Array2d
			if diceNumber <= 5 {
				shape = card.Problems[1].Shape
			} else {
				shape = card.Problems[8].Shape
			}
			// generate problems
			problems[key{animal, card.CardNumber, diceNumber}] = GenerateProblems(bf, shape, 3, 5, 3)
		}
	}
	// build groups of 4 problems
	//for k, v := range problems {

	//}

	// for each card-group (animal):
	// choose 10 different sets of 4 problems - one on each card - such that
	// the blocks on the four problems are available from the overall blockset

	// Visualize a solution
	/*
		a := app.New()
		w := a.NewWindow("Ubongo")
		imgWidth := 800
		imgHeight := 600
		if len(sols) > 0 {
			sol := sols[1]
			w.Resize(fyne.NewSize(float32(imgWidth), float32(imgHeight)))
			updateImage(w, sol, imgWidth, imgHeight, 0, 0, 0)

			var RX, RY, RZ float64 = 0.0, 0.0, 0.0
			go func() {
				const minFrameTime = 1.0 / 60 // min time to show one frame in seconds
				const speedRx = 0.0           // radians per second
				const speedRy = 0.2           // radians per second
				const speedRz = 0.0           // radians per second
				lastFrame := time.Now()
				for range time.Tick(time.Millisecond) {
					timePassed := float64(time.Since(lastFrame).Seconds())
					if timePassed >= minFrameTime {
						updateImage(w, sol, imgWidth, imgHeight, RX, RY, RZ)
						RX += speedRx * timePassed
						RY += speedRy * timePassed
						RZ += speedRz * timePassed
						lastFrame = time.Now()
					}
				}
			}()

			w.ShowAndRun()
		}
	*/
}

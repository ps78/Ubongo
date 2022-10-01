package main

import (
	//"fmt"
	//"time"

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func updateImage(win fyne.Window, sol *GameSolution, w, h int, rx, ry, rz float64) {
	img := GetSolutionImage(sol, w, h, rx, ry, rz)
	win.SetContent(canvas.NewImageFromImage(img))
}

func main() {
	//fp := GetProblemFactory()
	fb := GetBlockFactory()

	blocksets := CreateBlockSets(fb, 21, 5, 10)
	for _, bs := range blocksets {
		fmt.Println(bs)
	}

	// solve all problems:
	/*
		for card := 1; card <= 36; card++ {
			for dice := 1; dice <= 10; dice++ {
				p := fp.Get(Difficult, card, dice)
				if p != nil {
					g := NewGame(p)

					start := time.Now()
					solutions := g.Solve()
					runtime := time.Since(start)

					//for _, sol := range solutions {
					//	fmt.Println(sol.String())
					//}
					fmt.Printf("%s\tFound %d solutions in %s\n", p, len(solutions), runtime)
				}
			}
		}
	*/

	// Known Insane level problems:
	/*
		p := NewProblem(1, 1, Insane, 3, Elephant,
			fp.Get(Easy, 1, 1).Shape,
			[]*Block{fb.Yellow_bighook, fb.Green_T, fb.Blue_lighter, fb.Blue_v, fb.Red_flash})

			p := NewProblem(1, 5, Insane, 3, Elephant,
				fp.Get(Easy, 1, 5).Shape.Clone(),
				[]*Block{fb.Red_flash, fb.Blue_lighter, fb.Yellow_gate, fb.Green_L, fb.Blue_v})
	*/

	// solve some problem
	/*
		g := NewGame(p)
		start := time.Now()
		sols := g.Solve()
		runtime := time.Since(start)

		fmt.Println(p)
		fmt.Printf("%s\tFound %d solutions in %s\n", p, len(sols), runtime)
		for _, sol := range sols {
			fmt.Println(sol.String())
		}
	*/

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

package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func updateImage(win fyne.Window, sol *GameSolution, w, h int, rx, ry, rz float64) {
	img := GetSolutionImage(sol, w, h, rx, ry, rz)
	win.SetContent(canvas.NewImageFromImage(img))
}

func main() {
	f := GetProblemFactory()

	for card := 17; card <= 36; card++ {
		for dice := 1; dice <= 10; dice++ {
			p := f.Get(Easy, card, dice)
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

	/*
		bf := GetBlockFactory()
		p := f.Get(Difficult, 12, 1).Clone()
		p.Difficulty = Insane
		p.Height = 3
		p.Blocks = []*Block{bf.Blue_v, bf.Yellow_smallhook, bf.Red_flash, bf.Green_T, bf.Red_smallhook, bf.Yellow_gate}
	*/

	p := f.Get(Difficult, 2, 1)
	imgWidth := 800
	imgHeight := 600

	g := NewGame(p)
	sols := g.Solve()

	a := app.New()
	w := a.NewWindow("Ubongo")
	w.Resize(fyne.NewSize(float32(imgWidth), float32(imgHeight)))
	updateImage(w, sols[0], imgWidth, imgHeight, 0, 0, 0)

	var RX, RY, RZ float64 = 0.0, 0.0, 0.0
	go func() {
		const minFrameTime = 1.0 / 60 // min time to show one frame in seconds
		const speedRx = 0.0           // radians per second
		const speedRy = 0.0           // radians per second
		const speedRz = 0.0           // radians per second
		lastFrame := time.Now()
		for range time.Tick(time.Millisecond) {
			timePassed := float64(time.Since(lastFrame).Seconds())
			if timePassed >= minFrameTime {
				updateImage(w, sols[0], imgWidth, imgHeight, RX, RY, RZ)
				RX += speedRx * timePassed
				RY += speedRy * timePassed
				RZ += speedRz * timePassed
				lastFrame = time.Now()
			}
		}
	}()

	w.ShowAndRun()

}

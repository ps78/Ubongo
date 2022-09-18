package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	f := GetProblemFactory()

	for dice := 1; dice <= 10; dice++ {
		p := f.Get(Difficult, 12, dice)
		g := NewGame(p)

		start := time.Now()
		solutions := g.Solve()
		runtime := time.Since(start)

		//for _, sol := range solutions {
		//	fmt.Println(sol.String())
		//}
		fmt.Printf("%s\tFound %d solutions in %s\n", p, len(solutions), runtime)
	}

	bf := GetBlockFactory()
	p := f.Get(Difficult, 12, 1).Clone()
	p.Difficulty = Insane
	p.Height = 3
	p.Blocks = []*Block{bf.Blue_v, bf.Yellow_smallhook, bf.Red_flash, bf.Green_T, bf.Red_smallhook, bf.Yellow_gate}

	g := NewGame(p)
	sols := g.Solve()
	fmt.Printf("Found %d solutions for Insane problem\n", len(sols))

	img := GetSolutionImage(sols[0], 800, 600)

	a := app.New()
	w := a.NewWindow("Ubongo")
	w.Resize(fyne.NewSize(800, 600))
	w.SetContent(canvas.NewImageFromImage(img))

	/*
		hello := widget.NewLabel("Hello Fyne!")
		w.SetContent(container.NewVBox(
			hello,
			widget.NewButton("Hi!", func() {
				hello.SetText("Welcome :)")
			}),
		))
	*/

	w.ShowAndRun()
}

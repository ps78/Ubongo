package main

import (
	"fmt"
	"time"
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

	DrawSolution(sols[0], "solution0.png")
	//s := NewArray3d(3, 3, 3)
	//s.Set(2, 1, 1, IS_BLOCK)
	//DrawToFile(s, Blue, Vector{0, 0, 0}, "block.png")
}

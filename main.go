package main

import (
	"fmt"
	"time"
)

func main() {

	f := NewBlockFactory()

	blocks := []*Block{f.Get(8), f.Get(9), f.Get(12), f.Get(16)}
	area := NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	p := NewProblem("B12", 1, area, blocks)
	fmt.Println(p)

	g := NewGame(p.Shape, p.Height)
	fmt.Println(g)

	start := time.Now()
	solutions := g.Solve(blocks)
	runtime := time.Since(start)

	for _, sol := range solutions {
		fmt.Println(sol.String())
	}
	fmt.Printf("Found %d solutions in %s\n", len(solutions), runtime)
	fmt.Printf("recursive solver function was called %d times\n", recursiveSolverCount)

	fmt.Printf("%s", NewArray2dFromData([][]int8{{0, 1, 0}, {1, 0, 1}}))
}

package main

import (
	"fmt"
)

func main() {
	f := NewBlockFactory()

	blocks := []*Block{f.Get(8), f.Get(9), f.Get(12), f.Get(16)}
	area := NewArray2dFromData([][]int8{{0, -1, 0}, {0, 0, 0}, {-1, 0, 0}, {-1, -1, 0}})
	p := CreateProblem("B12", 1, area, blocks)
	fmt.Println(p)

	g := NewGame(p.Shape, p.Height)
	fmt.Println(g)

	g.Solve(blocks)

	/*
		allBlocks := f.GetAll()
		for _, block := range allBlocks {
			fmt.Println(block)
		}
	*/
}

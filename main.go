package main

import (
	"fmt"
)

func main() {
	b := MakeBlock08()
	p := MakeProblem("B12", 1, Array2d{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}}, []*Block{b})
	fmt.Println(p)
}

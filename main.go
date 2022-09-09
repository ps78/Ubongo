package main

import (
	"fmt"
)

func main() {
	b := CreateBlock08()
	p := CreateProblem("B12", 1, Array2d{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}}, []*Block{b})
	fmt.Println(p)
}

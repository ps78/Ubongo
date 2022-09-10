package main

import (
	"fmt"
)

func main() {
	f := NewBlockFactory()
	b := f.Get(8)
	p := CreateProblem("B12", 1, *NewArray2dFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}}), []*Block{b})
	fmt.Println(p)
}

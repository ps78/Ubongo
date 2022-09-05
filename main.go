package main

import (
	"fmt"
)

func main() {
	b08 := MakeBlock08()
	p := makeProblem("B12", 1, ProblemShape{{1, 1, 0, 0}, {0, 1, 1, 0}, {1, 1, 1, 1}}, []*Block{b08})

	fmt.Println(p)
	fmt.Println(b08)
}

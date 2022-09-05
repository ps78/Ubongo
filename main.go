package main

import (
	"fmt"
)

func main() {
	b := MakeBlock08()
	p := MakeProblem("B12", 1, ProblemShape{{1, 1, 0, 0}, {0, 1, 1, 0}, {1, 1, 1, 1}}, []*Block{b})

	v := p.CreateVolume()
	fmt.Printf("%v", *v)
}

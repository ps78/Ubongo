package main

import (
	"fmt"
)

func main() {
	b := MakeBlock08()
	p := makeProblem("B12", 1, ProblemShape{{1, 1, 0, 0}, {0, 1, 1, 0}, {1, 1, 1, 1}}, []*Block{b})

	fmt.Println(p)
	fmt.Println(b)
	fmt.Println("block bounding box: ", GetBoundingBoxFromBlockShape(b.Shapes[0]))
	fmt.Println("problem bounding box: ", GetBoundingBoxFromProblemShape(p.Shape, p.Height))
}

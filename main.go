package main

import (
	"fmt"
)

type Block struct {
	Name         string
	Shape        [][][]int8
	Orientations [][][][]int8
}

type Volume struct {
	Shape [][][]int8 // 0=not part of volume, 1=empty, 2=occupied by a block
}

/*
Creates a block with the given name/shape, including all orientations

func makeBlock(name string, shape [][][]int8) Block {
}
*/
/*
// Rotates the given shape 90° around the X-Axis
func RotateX(shape [][][]int8) [][][]int8 {

}
*/

func makeBlocks() []Block {
	var blocks = make([]Block, 4)
	blocks[0] = Block{Name: "01", Shape: [][][]int8{{{0, 1}, {1, 1}, {0, 1}}, {{0, 1}, {0, 0}, {0, 0}}}}
	blocks[1] = Block{Name: "02", Shape: [][][]int8{{{1, 1}, {0, 1}, {0, 1}}, {{0, 0}, {0, 0}, {0, 1}}}}
	blocks[2] = Block{Name: "03", Shape: [][][]int8{{{1, 1}, {1, 0}}, {{0, 1}, {0, 0}}}}
	blocks[3] = Block{Name: "04", Shape: [][][]int8{{{1, 1}, {0, 1}, {1, 1}}}}
	return blocks
}

func main() {
	b := makeBlocks()
	v := Volume{Shape: [][][]int8{
		{{0, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 0}, {1, 1, 1, 0, 0}},
		{{0, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 0}, {1, 1, 1, 0, 0}}}}
	fmt.Println(b[0], v)
}

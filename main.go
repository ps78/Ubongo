package main

import (
	"fmt"
)

type Block struct {
	Name  string
	Cubes [][][]bool
}

type Volume struct {
	Shape  [][][]int8 // 0=not part of volume, 1=empty, 2=occupied by a block
	Height int
}

func makeBlocks() []Block {
	var blocks = make([]Block, 4)
	blocks[0] = Block{Name: "01", Cubes: [][][]bool{{{false, true}, {true, true}, {false, true}}, {{false, true}, {false, false}, {false, false}}}}
	blocks[1] = Block{Name: "02", Cubes: [][][]bool{{{true, true}, {false, true}, {false, true}}, {{false, false}, {false, false}, {false, true}}}}
	blocks[2] = Block{Name: "03", Cubes: [][][]bool{{{true, true}, {true, false}}, {{false, true}, {false, false}}}}
	blocks[3] = Block{Name: "04", Cubes: [][][]bool{{{true, true}, {false, true}, {true, true}}}}
	return blocks
}

func main() {
	b := makeBlocks()
	v := Volume{Height: 3, Shape: [][][]int8{{{0, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 0}, {1, 1, 1, 0, 0}}, {{0, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 0}, {1, 1, 1, 0, 0}}}}
	fmt.Println(b[0], v)
}

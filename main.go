package main

import (
	"fmt"
)

type Block struct {
	name  string
	cubes [][][]bool
}

func makeBlocks() []Block {
	var blocks = make([]Block, 3)
	blocks[0] = Block{"L", [][][]bool{{{true, false}, {true, false}, {true, true}}}}
	blocks[1] = Block{"v", [][][]bool{{{true, false}, {true, true}}}}
	blocks[2] = Block{"M", [][][]bool{{{true, false}, {true, true}, {true, false}}}}
	return blocks
}

func main() {
	b := makeBlocks()
	fmt.Println(b)
}

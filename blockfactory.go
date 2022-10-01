/*
*******************************************************************

This file contains the singleton BlockFactory, which is used
to ensure that the blocks of the game are only created once
and can be access thread-safely

*******************************************************************
*/
package main

import (
	"strings"
	"sync"
)

// used to create a thread-safe singleton instance of BlockFactory
var onceBlockFactorySingleton sync.Once

// the singleton
var blockFactoryInstance *BlockFactory

type BlockFactory struct {
	Block          map[int]*Block // access block by number
	MinBlockNumber int
	MaxBlockNumber int
	BlockByVolume  map[int]([]*Block) // access blocks by volume (3, 4 or 5)

	// accessors for blocks by name, for convenience
	Yellow_hello     *Block
	Yellow_bighook   *Block
	Yellow_smallhook *Block
	Yellow_gate      *Block
	Blue_bighook     *Block
	Blue_flash       *Block
	Blue_lighter     *Block
	Blue_v           *Block
	Red_stool        *Block
	Red_smallhook    *Block
	Red_bighook      *Block
	Red_flash        *Block
	Green_flash      *Block
	Green_bighook    *Block
	Green_T          *Block
	Green_L          *Block
}

func GetBlockFactory() *BlockFactory {
	onceBlockFactorySingleton.Do(func() {
		f := new(BlockFactory)
		f.MinBlockNumber = 1
		f.MaxBlockNumber = 16
		f.Block = map[int]*Block{
			1:  NewBlock1(),
			2:  NewBlock2(),
			3:  NewBlock3(),
			4:  NewBlock4(),
			5:  NewBlock5(),
			6:  NewBlock6(),
			7:  NewBlock7(),
			8:  NewBlock8(),
			9:  NewBlock9(),
			10: NewBlock10(),
			11: NewBlock11(),
			12: NewBlock12(),
			13: NewBlock13(),
			14: NewBlock14(),
			15: NewBlock15(),
			16: NewBlock16()}

		f.Yellow_hello = f.Block[1]
		f.Yellow_bighook = f.Block[2]
		f.Yellow_smallhook = f.Block[3]
		f.Yellow_gate = f.Block[4]
		f.Blue_bighook = f.Block[5]
		f.Blue_flash = f.Block[6]
		f.Blue_lighter = f.Block[7]
		f.Blue_v = f.Block[8]
		f.Red_stool = f.Block[9]
		f.Red_smallhook = f.Block[10]
		f.Red_bighook = f.Block[11]
		f.Red_flash = f.Block[12]
		f.Green_flash = f.Block[13]
		f.Green_bighook = f.Block[14]
		f.Green_T = f.Block[15]
		f.Green_L = f.Block[16]

		// add the blocks to the volume-map
		f.BlockByVolume = make(map[int]([]*Block), 0)
		for _, b := range f.Block {
			if _, ok := f.BlockByVolume[b.Volume]; !ok {
				f.BlockByVolume[b.Volume] = make([]*Block, 0)
			}
			f.BlockByVolume[b.Volume] = append(f.BlockByVolume[b.Volume], b)
		}

		blockFactoryInstance = f
	})
	return blockFactoryInstance
}

// Returns the block with the given number, or nil of not found
func (f *BlockFactory) Get(blockNumber int) *Block {
	// get block from cache if possible
	if block, ok := f.Block[blockNumber]; ok {
		return block
	} else {
		return nil
	}
}

// Returns a block by it's color and name, or nil of not found
func (f *BlockFactory) ByName(color BlockColor, name string) *Block {
	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		b := f.Get(i)
		if strings.EqualFold(b.Name, name) && b.Color == color {
			return b
		}
	}
	return nil
}

// Returns an array with all blocks
func (f *BlockFactory) GetAll() []*Block {
	a := make([]*Block, 0)
	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		if block := f.Get(i); block != nil {
			a = append(a, block)
		}
	}
	return a
}

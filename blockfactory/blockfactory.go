/*
*******************************************************************

This file contains the singleton BlockFactory, which is used
to ensure that the blocks of the game are only created once
and can be access thread-safely

*******************************************************************
*/
package blockfactory

import (
	"math/rand"
	"strings"
	"sync"
	"time"
	"ubongo/base/array3d"
	"ubongo/block"
	"ubongo/blockset"
	"ubongo/extmath"
)

// used to create a thread-safe singleton instance of BlockFactory
var onceBlockFactorySingleton sync.Once

// the singleton
var blockFactoryInstance *BlockFactory

type BlockFactory struct {
	Block          map[int]*block.Block // access block by number
	MinBlockNumber int
	MaxBlockNumber int
	BlockByVolume  map[int]*blockset.Blockset // access blocks by volume (3, 4 or 5)

	// accessors for blocks by name, for convenience
	Yellow_hello     *block.Block
	Yellow_bighook   *block.Block
	Yellow_smallhook *block.Block
	Yellow_gate      *block.Block
	Blue_bighook     *block.Block
	Blue_flash       *block.Block
	Blue_lighter     *block.Block
	Blue_v           *block.Block
	Red_stool        *block.Block
	Red_smallhook    *block.Block
	Red_bighook      *block.Block
	Red_flash        *block.Block
	Green_flash      *block.Block
	Green_bighook    *block.Block
	Green_T          *block.Block
	Green_L          *block.Block
}

func GetBlockFactory() *BlockFactory {
	onceBlockFactorySingleton.Do(func() {
		f := new(BlockFactory)
		f.MinBlockNumber = 1
		f.MaxBlockNumber = 16
		f.Block = map[int]*block.Block{
			1:  newBlock1(),
			2:  newBlock2(),
			3:  newBlock3(),
			4:  newBlock4(),
			5:  newBlock5(),
			6:  newBlock6(),
			7:  newBlock7(),
			8:  newBlock8(),
			9:  newBlock9(),
			10: newBlock10(),
			11: newBlock11(),
			12: newBlock12(),
			13: newBlock13(),
			14: newBlock14(),
			15: newBlock15(),
			16: newBlock16()}

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
		f.BlockByVolume = make(map[int]*blockset.Blockset)
		for _, b := range f.Block {
			if _, ok := f.BlockByVolume[b.Volume]; !ok {
				f.BlockByVolume[b.Volume] = blockset.NewBlockset()
			}
			f.BlockByVolume[b.Volume].Add(b)
		}

		blockFactoryInstance = f
	})
	return blockFactoryInstance
}

// Returns the block with the given number, or nil of not found
func (f *BlockFactory) ByNumber(blockNumber int) *block.Block {
	// get block from cache if possible
	if block, ok := f.Block[blockNumber]; ok {
		return block
	} else {
		return nil
	}
}

// Returns a block by it's color and name, or nil of not found
func (f *BlockFactory) ByName(color block.BlockColor, name string) *block.Block {
	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		b := f.ByNumber(i)
		if strings.EqualFold(b.Name, name) && b.Color == color {
			return b
		}
	}
	return nil
}

// Returns an array with all blocks
func (f *BlockFactory) GetAll() []*block.Block {
	a := make([]*block.Block, 0)
	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		if block := f.ByNumber(i); block != nil {
			a = append(a, block)
		}
	}
	return a
}

// GenerateBlocksets returns resultCount blocksets randomly generated
// from all blocks such that:
//   - the sum of the blockvolumes equals volume
//   - the number of blocks in each set is equal to blockCount
//   - every block type is only used once
//
// NOTE: the returned slice can be smaller than resultCount, as no duplicate
// results are generated and the algormithms stops if it fails to generate new
// results
func (bf *BlockFactory) GenerateBlocksets(volume, blockCount, resultCount int) []*blockset.Blockset {
	// maximum number of tries to create a random blockset that is not already in
	// the result
	const maxTry = 10

	// create all possible partitions that define how many blocks of a specific volume
	// are used to fill the volume
	// limit the 3-volume
	max3 := bf.BlockByVolume[3].Count
	max4 := bf.BlockByVolume[4].Count
	max5 := bf.BlockByVolume[5].Count
	partitions := extmath.CreateParitions(volume, []int{3, 4, 5}, map[int]int{3: max3, 4: max4, 5: max5}, blockCount)
	partCount := len(partitions)

	// abort if there are no partitions fulfilling the given criteria
	if partCount == 0 {
		return []*blockset.Blockset{}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate resultCount results as requested
	results := make([]*blockset.Blockset, resultCount)
	for i := 0; i < resultCount; i++ {
		for try := 0; try < maxTry; try++ {
			// initialize new result
			curResult := blockset.NewBlockset()

			// randomly choose a partition
			partition := partitions[r.Intn(partCount)]

			// cycle through the volume-keys of the chosen partition
			for vol, count := range partition {
				// randomly choose count blocks of the given volume
				for j := 0; j < count; j++ {
					randIdx := -1
					for {
						randIdx = r.Intn(bf.BlockByVolume[vol].Count)
						if !curResult.Contains(bf.BlockByVolume[vol].At(randIdx).Number) {
							break
						}
					}
					bl := bf.BlockByVolume[vol].At(randIdx)
					curResult.Add(bl)
				}
			}
			if !blockset.ContainsBlockset(results, curResult) {
				results[i] = curResult
				break
			}
			// abort if we couldn't create a new result after maxTries
			if try == maxTry-1 {
				return results[:i]
			}
		}
	}
	return results
}

/*******************************************************************************
 * Block-creator functions for the 16 blocks of the game
 *******************************************************************************/
func newBlock1() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {1, 0}, {0, 0}}})
	return &block.Block{
		Number: 1,
		Name:   "hello",
		Color:  block.Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock2() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 0}, {1, 0}, {1, 1}}, {{1, 0}, {0, 0}, {0, 0}}})
	return &block.Block{
		Number: 2,
		Name:   "big hook",
		Color:  block.Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock3() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 0}, {0, 0}}, {{1, 1}, {0, 1}}})
	return &block.Block{
		Number: 3,
		Name:   "small hook",
		Color:  block.Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock4() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}, {1, 1}}})
	return &block.Block{
		Number: 4,
		Name:   "gate",
		Color:  block.Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock5() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {0, 0}, {0, 0}}, {{1, 0}, {1, 0}, {1, 0}}})
	return &block.Block{
		Number: 5,
		Name:   "big hook",
		Color:  block.Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock6() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 0}, {1, 1}, {0, 1}}, {{1, 0}, {0, 0}, {0, 0}}})
	return &block.Block{
		Number: 6,
		Name:   "flash",
		Color:  block.Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock7() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1}, {1}, {1}}, {{0}, {1}, {1}}})
	return &block.Block{
		Number: 7,
		Name:   "lighter",
		Color:  block.Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock8() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{0, 1}, {1, 1}}})
	return &block.Block{
		Number: 8,
		Name:   "v",
		Color:  block.Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock9() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}}, {{1, 0}, {1, 0}}})
	return &block.Block{
		Number: 9,
		Name:   "stool",
		Color:  block.Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock10() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}}, {{0, 0}, {1, 0}}})
	return &block.Block{
		Number: 10,
		Name:   "small hook",
		Color:  block.Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock11() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {0, 0}, {1, 0}}})
	return &block.Block{
		Number: 11,
		Name:   "big hook",
		Color:  block.Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock12() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1}, {1}, {0}}, {{0}, {1}, {1}}})
	return &block.Block{
		Number: 12,
		Name:   "flash",
		Color:  block.Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock13() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}, {0, 0}}, {{0, 0}, {1, 0}, {1, 0}}})
	return &block.Block{
		Number: 13,
		Name:   "flash",
		Color:  block.Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock14() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1, 0}, {1, 0}, {1, 0}}, {{1, 1}, {0, 0}, {0, 0}}})
	return &block.Block{
		Number: 14,
		Name:   "big hook",
		Color:  block.Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock15() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1}, {1}, {1}}, {{0}, {1}, {0}}})
	return &block.Block{
		Number: 15,
		Name:   "T",
		Color:  block.Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock16() *block.Block {
	baseShape := array3d.NewFromData([][][]int8{{{1}, {1}, {1}}, {{1}, {0}, {0}}})
	return &block.Block{
		Number: 16,
		Name:   "L",
		Color:  block.Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

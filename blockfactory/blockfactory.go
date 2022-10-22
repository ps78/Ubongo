// Package blockfactory contains the F (BlockFactory) singleton type and related methods
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

// onceBlockFactorySingleton is used to create a thread-safe singleton instance of BlockFactory
var onceBlockFactorySingleton sync.Once

// blockFactoryInstance the singleton instance of the BlockFactory
var blockFactoryInstance *F

// F represents factory that allows accessing the game's various blocks
type F struct {
	// block is a map where each of the game's 16 block types can be accessed by their number 1..16
	block map[int]*block.B

	// MinBlockNumber can be used to iterate over Block, it is equal to 1
	MinBlockNumber int

	// MaxBlockNumber can be used to iterate over Block, it is equal to 16
	MaxBlockNumber int

	// byVolume is an map to access blocks based on their volume (which is the key)
	byVolume map[int]*blockset.S

	// The following are accessors for the 16 blocks by human readable name, for convenience

	// Block number 1
	Yellow_hello *block.B
	// Block number 2
	Yellow_bighook *block.B
	// Block number 3
	Yellow_smallhook *block.B
	// Block number 4
	Yellow_gate *block.B
	// Block number 5
	Blue_bighook *block.B
	// Block number 6
	Blue_flash *block.B
	// Block number 7
	Blue_lighter *block.B
	// Block number 8
	Blue_v *block.B
	// Block number 9
	Red_stool *block.B
	// Block number 10
	Red_smallhook *block.B
	// Block number 11
	Red_bighook *block.B
	// Block number 12
	Red_flash *block.B
	// Block number 13
	Green_flash *block.B
	// Block number 14
	Green_bighook *block.B
	// Block number 15
	Green_T *block.B
	// Block number 16
	Green_L *block.B
}

// Get returns the singleton instance of the BlockFactory
func Get() *F {
	onceBlockFactorySingleton.Do(func() {
		f := new(F)
		f.MinBlockNumber = 1
		f.MaxBlockNumber = 16
		f.block = map[int]*block.B{
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

		f.Yellow_hello = f.block[1]
		f.Yellow_bighook = f.block[2]
		f.Yellow_smallhook = f.block[3]
		f.Yellow_gate = f.block[4]
		f.Blue_bighook = f.block[5]
		f.Blue_flash = f.block[6]
		f.Blue_lighter = f.block[7]
		f.Blue_v = f.block[8]
		f.Red_stool = f.block[9]
		f.Red_smallhook = f.block[10]
		f.Red_bighook = f.block[11]
		f.Red_flash = f.block[12]
		f.Green_flash = f.block[13]
		f.Green_bighook = f.block[14]
		f.Green_T = f.block[15]
		f.Green_L = f.block[16]

		// add the blocks to the volume-map
		f.byVolume = make(map[int]*blockset.S)
		for _, b := range f.block {
			if _, ok := f.byVolume[b.Volume]; !ok {
				f.byVolume[b.Volume] = blockset.New()
			}
			f.byVolume[b.Volume].Add(b)
		}

		blockFactoryInstance = f
	})
	return blockFactoryInstance
}

// ByNumber returns the block with the given number, or nil of not found
func (f *F) ByNumber(blockNumber int) *block.B {
	if f == nil {
		return nil
	}
	// get block from cache if possible
	if block, ok := f.block[blockNumber]; ok {
		return block
	} else {
		return nil
	}
}

// ByName returns a block by it's color and name, or nil of not found
func (f *F) ByName(color block.BlockColor, name string) *block.B {
	if f == nil {
		return nil
	}
	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		b := f.ByNumber(i)
		if strings.EqualFold(b.Name, name) && b.Color == color {
			return b
		}
	}
	return nil
}

// ByVolume returns a blockset containing all blocks with the given volume
// returns nil if no such block exists
func (f *F) ByVolume(volume int) *blockset.S {
	if f == nil {
		return blockset.New()
	}
	if bs, ok := f.byVolume[volume]; ok {
		return bs
	} else {
		return blockset.New()
	}
}

// GetAll returns the set with all blocks
func (f *F) GetAll() *blockset.S {
	a := blockset.New()
	if f != nil {
		for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
			a.Add(f.ByNumber(i))
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
func (bf *F) GenerateBlocksets(volume, blockCount, resultCount int) []*blockset.S {
	if bf == nil {
		return []*blockset.S{}
	}

	// maximum number of tries to create a random blockset that is not already in
	// the result
	const maxTry = 10

	// create all possible partitions that define how many blocks of a specific volume
	// are used to fill the volume
	// limit the 3-volume
	max3 := bf.byVolume[3].Count
	max4 := bf.byVolume[4].Count
	max5 := bf.byVolume[5].Count
	partitions := extmath.CreateParitions(volume, []int{3, 4, 5}, map[int]int{3: max3, 4: max4, 5: max5}, blockCount)
	partCount := len(partitions)

	// abort if there are no partitions fulfilling the given criteria
	if partCount == 0 {
		return []*blockset.S{}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate resultCount results as requested
	results := make([]*blockset.S, resultCount)
	for i := 0; i < resultCount; i++ {
		for try := 0; try < maxTry; try++ {
			// initialize new result
			curResult := blockset.New()

			// randomly choose a partition
			partition := partitions[r.Intn(partCount)]

			// cycle through the volume-keys of the chosen partition
			for vol, count := range partition {
				// randomly choose count blocks of the given volume
				for j := 0; j < count; j++ {
					randIdx := -1
					for {
						randIdx = r.Intn(bf.byVolume[vol].Count)
						if !curResult.Contains(bf.byVolume[vol].Get(randIdx).Number) {
							break
						}
					}
					bl := bf.byVolume[vol].Get(randIdx)
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

// *********************************************************** //
// * Block-creator functions for the 16 blocks of the game   * //
// *********************************************************** //

func newBlock1() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {1, 0}, {0, 0}}})
	return &block.B{
		Number: 1,
		Name:   "hello",
		Color:  block.Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock2() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 0}, {1, 0}, {1, 1}}, {{1, 0}, {0, 0}, {0, 0}}})
	return &block.B{
		Number: 2,
		Name:   "big hook",
		Color:  block.Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock3() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 0}, {0, 0}}, {{1, 1}, {0, 1}}})
	return &block.B{
		Number: 3,
		Name:   "small hook",
		Color:  block.Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock4() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}, {1, 1}}})
	return &block.B{
		Number: 4,
		Name:   "gate",
		Color:  block.Yellow,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock5() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {0, 0}, {0, 0}}, {{1, 0}, {1, 0}, {1, 0}}})
	return &block.B{
		Number: 5,
		Name:   "big hook",
		Color:  block.Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock6() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 0}, {1, 1}, {0, 1}}, {{1, 0}, {0, 0}, {0, 0}}})
	return &block.B{
		Number: 6,
		Name:   "flash",
		Color:  block.Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock7() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1}, {1}, {1}}, {{0}, {1}, {1}}})
	return &block.B{
		Number: 7,
		Name:   "lighter",
		Color:  block.Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock8() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{0, 1}, {1, 1}}})
	return &block.B{
		Number: 8,
		Name:   "v",
		Color:  block.Blue,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock9() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}}, {{1, 0}, {1, 0}}})
	return &block.B{
		Number: 9,
		Name:   "stool",
		Color:  block.Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock10() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}}, {{0, 0}, {1, 0}}})
	return &block.B{
		Number: 10,
		Name:   "small hook",
		Color:  block.Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock11() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}, {1, 0}}, {{0, 0}, {0, 0}, {1, 0}}})
	return &block.B{
		Number: 11,
		Name:   "big hook",
		Color:  block.Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock12() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1}, {1}, {0}}, {{0}, {1}, {1}}})
	return &block.B{
		Number: 12,
		Name:   "flash",
		Color:  block.Red,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock13() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 1}, {1, 0}, {0, 0}}, {{0, 0}, {1, 0}, {1, 0}}})
	return &block.B{
		Number: 13,
		Name:   "flash",
		Color:  block.Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock14() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1, 0}, {1, 0}, {1, 0}}, {{1, 1}, {0, 0}, {0, 0}}})
	return &block.B{
		Number: 14,
		Name:   "big hook",
		Color:  block.Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock15() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1}, {1}, {1}}, {{0}, {1}, {0}}})
	return &block.B{
		Number: 15,
		Name:   "T",
		Color:  block.Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

func newBlock16() *block.B {
	baseShape := array3d.NewFromData([][][]int8{{{1}, {1}, {1}}, {{1}, {0}, {0}}})
	return &block.B{
		Number: 16,
		Name:   "L",
		Color:  block.Green,
		Shapes: baseShape.CreateRotations(),
		Volume: baseShape.Count(1)}
}

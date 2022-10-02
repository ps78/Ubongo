package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Blockset is a (unordered) set of blocks
type Blockset struct {
	items []*Block
	Count int
}

// Sorts the items in bs by the block number
func (bs *Blockset) normalize() {
	sort.Slice(bs.items, func(i, j int) bool {
		return bs.items[i].Number < bs.items[j].Number
	})
}

// Creates a blockset from a list/slice of blocks
func NewBlockset(block ...*Block) *Blockset {
	bs := Blockset{}
	bs.Count = len(block)
	bs.items = make([]*Block, len(block))
	copy(bs.items, block)
	bs.normalize()
	return &bs
}

// Returns the block reference for the given index
func (bs *Blockset) At(idx int) *Block {
	if idx < 0 || idx >= bs.Count {
		return nil
	} else {
		return bs.items[idx]
	}
}

// Returns a string representation of a blockset
func (bs *Blockset) String() string {
	s := "["
	for i := 0; i < bs.Count; i++ {
		if i != 0 {
			s += ", "
		}
		s += fmt.Sprintf("%s %s", bs.At(i).Color, bs.At(i).Name)
	}
	return s + "]"
}

// Returns a copy of the blockset as slice. The slice is ordered
// in contrast to the blockset
func (bs *Blockset) AsSlice() []*Block {
	s := make([]*Block, bs.Count)
	copy(s, bs.items)
	return s
}

// Adds block to the blockset
func (bs *Blockset) Add(blocks ...*Block) {
	if blocks == nil {
		return
	}
	for _, b := range blocks {
		if b != nil {
			bs.items = append(bs.items, b)
		}
	}
	bs.Count = len(bs.items)
	bs.normalize()
}

// Removes all blocks with the given blockNumber from the blockset
func (bs *Blockset) Remove(blockNumber int) {
	if bs.Contains(blockNumber) {
		newItems := make([]*Block, 0, bs.Count-1)
		for oldIdx := range bs.items {
			if bs.items[oldIdx].Number != blockNumber {
				newItems = append(newItems, bs.items[oldIdx])
			}
		}
		bs.items = newItems
		bs.Count = len(newItems)
	}
}

func (bs *Blockset) RemoveAt(idx int) {
	if idx < 0 || idx >= bs.Count {
		return
	} else if idx == bs.Count-1 {
		bs.items = bs.items[:idx]
	} else {
		bs.items = append(bs.items[:idx], bs.items[idx+1:]...)
	}
	bs.Count = len(bs.items)
}

// Compares two Blocksets. Considers blocks with the same number as identical
func (bs *Blockset) IsEqual(other *Blockset) bool {
	if bs == nil || other == nil || bs.Count != other.Count {
		return false
	}
	for i := 0; i < bs.Count; i++ {
		if bs.items[i].Number != other.items[i].Number {
			return false
		}
	}
	return true
}

// Returns true of the set contains a block with the given number
func (bs *Blockset) Contains(blockNumber int) bool {
	for _, b := range bs.items {
		if b.Number == blockNumber {
			return true
		}
	}
	return false
}

// Creates a clone of the blockset
func (orig *Blockset) Clone() *Blockset {
	clone := Blockset{}
	clone.Count = orig.Count
	clone.items = make([]*Block, orig.Count)
	copy(clone.items, orig.items)
	return &clone
}

// Returns the total volume of the blockset
func (bs *Blockset) Volume() int {
	sum := 0
	for _, bl := range bs.items {
		sum += bl.Volume
	}
	return sum
}

func ContainsBlockset(sets []*Blockset, bs *Blockset) bool {
	if sets == nil || bs == nil {
		return false
	}

	for _, s := range sets {
		if bs.IsEqual(s) {
			return true
		}
	}
	return false
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
func GenerateBlocksets(bf *BlockFactory, volume, blockCount, resultCount int) []*Blockset {
	// maximum number of tries to create a random blockset that is not already in
	// the result
	const maxTry = 10

	// create all possible partitions that define how many blocks of a specific volume
	// are used to fill the volume
	// limit the 3-volume
	max3 := bf.BlockByVolume[3].Count
	max4 := bf.BlockByVolume[4].Count
	max5 := bf.BlockByVolume[5].Count
	partitions := CreateParitions(volume, []int{3, 4, 5}, map[int]int{3: max3, 4: max4, 5: max5}, blockCount)
	partCount := len(partitions)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate resultCount results as requested
	results := make([]*Blockset, resultCount)
	for i := 0; i < resultCount; i++ {
		for try := 0; try < maxTry; try++ {
			// initialize new result
			curResult := NewBlockset()

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
			if !ContainsBlockset(results, curResult) {
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

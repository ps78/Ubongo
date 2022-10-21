package blockset

import (
	"fmt"
	"sort"
	"ubongo/block"
)

// Blockset is a (unordered) set of blocks
type Blockset struct {
	Items []*block.Block
	Count int
}

// Sorts the items in bs by the block number
func (bs *Blockset) normalize() {
	sort.Slice(bs.Items, func(i, j int) bool {
		return bs.Items[i].Number < bs.Items[j].Number
	})
}

// Creates a blockset from a list/slice of blocks
func NewBlockset(blocks ...*block.Block) *Blockset {
	bs := Blockset{}
	bs.Count = len(blocks)
	bs.Items = make([]*block.Block, len(blocks))
	copy(bs.Items, blocks)
	bs.normalize()
	return &bs
}

// Returns the block reference for the given index
func (bs *Blockset) At(idx int) *block.Block {
	if idx < 0 || idx >= bs.Count {
		return nil
	} else {
		return bs.Items[idx]
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
func (bs *Blockset) AsSlice() []*block.Block {
	s := make([]*block.Block, bs.Count)
	copy(s, bs.Items)
	return s
}

// Adds block to the blockset
func (bs *Blockset) Add(blocks ...*block.Block) {
	if blocks == nil {
		return
	}
	for _, b := range blocks {
		if b != nil {
			bs.Items = append(bs.Items, b)
		}
	}
	bs.Count = len(bs.Items)
	bs.normalize()
}

// Removes all blocks with the given blockNumber from the blockset
func (bs *Blockset) Remove(blockNumber int) {
	if bs.Contains(blockNumber) {
		newItems := make([]*block.Block, 0, bs.Count-1)
		for oldIdx := range bs.Items {
			if bs.Items[oldIdx].Number != blockNumber {
				newItems = append(newItems, bs.Items[oldIdx])
			}
		}
		bs.Items = newItems
		bs.Count = len(newItems)
	}
}

func (bs *Blockset) RemoveAt(idx int) {
	if idx < 0 || idx >= bs.Count {
		return
	} else if idx == bs.Count-1 {
		bs.Items = bs.Items[:idx]
	} else {
		bs.Items = append(bs.Items[:idx], bs.Items[idx+1:]...)
	}
	bs.Count = len(bs.Items)
}

// Compares two Blocksets. Considers blocks with the same number as identical
func (bs *Blockset) IsEqual(other *Blockset) bool {
	if bs == nil || other == nil || bs.Count != other.Count {
		return false
	}
	for i := 0; i < bs.Count; i++ {
		if bs.Items[i].Number != other.Items[i].Number {
			return false
		}
	}
	return true
}

// Returns true of the set contains a block with the given number
func (bs *Blockset) Contains(blockNumber int) bool {
	for _, b := range bs.Items {
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
	clone.Items = make([]*block.Block, orig.Count)
	copy(clone.Items, orig.Items)
	return &clone
}

// Returns the total volume of the blockset
func (bs *Blockset) Volume() int {
	sum := 0
	for _, bl := range bs.Items {
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

// Package blockset contains the type S (Blockset) and its methods
package blockset

import (
	"fmt"
	"sort"
	"ubongo/block"
)

// S is a (unordered) set of blocks, without duplicates
type S struct {
	items []*block.B
	Count int
}

// normalize sorts the items in bs by the block number
func (bs *S) normalize() {
	sort.Slice(bs.items, func(i, j int) bool {
		return bs.items[i].Number < bs.items[j].Number
	})
}

// New creates a blockset from a list/slice of blocks
func New(blocks ...*block.B) *S {
	bs := S{}
	bs.items = make([]*block.B, 0)
	for _, block := range blocks {
		bs.Add(block)
	}
	return &bs
}

// Get returns the block reference for the given index
// returns nil if the index is invalid
func (bs *S) Get(idx int) *block.B {
	if bs == nil {
		return nil
	}
	if idx < 0 || idx >= bs.Count {
		return nil
	} else {
		return bs.items[idx]
	}
}

// String returns a string representation of a blockset
func (bs *S) String() string {
	if bs == nil {
		return "(nil)"
	} else {
		s := "["
		for i := 0; i < bs.Count; i++ {
			if i != 0 {
				s += ", "
			}
			s += fmt.Sprintf("%s %s", bs.Get(i).Color, bs.Get(i).Name)
		}
		return s + "]"
	}
}

// AsSlice returns a copy of the blockset as slice. The slice is ordered
// by block number
func (bs *S) AsSlice() []*block.B {
	if bs == nil {
		return nil
	} else {
		s := make([]*block.B, bs.Count)
		copy(s, bs.items)
		return s
	}
}

// Add adds block to the blockset, if it doesn't already exist
func (bs *S) Add(blocks ...*block.B) {
	if bs == nil || blocks == nil {
		return
	}
	for _, b := range blocks {
		if b != nil {
			if !bs.Contains(b.Number) {
				bs.items = append(bs.items, b)
			}
		}
	}
	bs.Count = len(bs.items)
	bs.normalize()
}

// Remove removes the block with the given blockNumber from the blockset
func (bs *S) Remove(blockNumber int) {
	if bs == nil {
		return
	} else {
		if bs.Contains(blockNumber) {
			newItems := make([]*block.B, 0, bs.Count-1)
			for oldIdx := range bs.items {
				if bs.items[oldIdx].Number != blockNumber {
					newItems = append(newItems, bs.items[oldIdx])
				}
			}
			bs.items = newItems
			bs.Count = len(newItems)
		}
	}
}

// RemoveAt removes a block from the set by its index
func (bs *S) RemoveAt(idx int) {
	if bs == nil || idx < 0 || idx >= bs.Count {
		return
	} else if idx == bs.Count-1 {
		bs.items = bs.items[:idx]
	} else {
		bs.items = append(bs.items[:idx], bs.items[idx+1:]...)
	}
	bs.Count = len(bs.items)
}

// Equals compares two Blocksets. Considers blocks with the same number as identical
func (bs *S) Equals(other *S) bool {
	if bs == nil && other == nil {
		return true
	} else if bs == nil || other == nil || bs.Count != other.Count {
		return false
	} else {
		for i := 0; i < bs.Count; i++ {
			if bs.items[i].Number != other.items[i].Number {
				return false
			}
		}
		return true
	}
}

// Contains returns true of the set contains a block with the given number
func (bs *S) Contains(blockNumber int) bool {
	if bs == nil {
		return false
	} else {
		for _, b := range bs.items {
			if b.Number == blockNumber {
				return true
			}
		}
		return false
	}
}

// Clone creates a clone of the blockset, copying the block references,
// not the blocks themselves
func (orig *S) Clone() *S {
	if orig == nil {
		return nil
	} else {
		clone := S{}
		clone.Count = orig.Count
		clone.items = make([]*block.B, orig.Count)
		copy(clone.items, orig.items)
		return &clone
	}
}

// Volume returns the total volume of the blockset
// 0 for a nil reference
func (bs *S) Volume() int {
	sum := 0
	if bs != nil {
		for _, bl := range bs.items {
			sum += bl.Volume
		}
	}
	return sum
}

// ContainsBlockset checks if a slice of blocksets contains a
// specific blockset
func ContainsBlockset(sets []*S, bs *S) bool {
	if sets == nil || bs == nil {
		return false
	}

	for _, s := range sets {
		if bs.Equals(s) {
			return true
		}
	}
	return false
}

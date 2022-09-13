package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockColorString(t *testing.T) {
	f := NewBlockFactory()
	for _, b := range f.GetAll() {
		colorName := b.Color.String()
		assert.True(t, len(colorName) > 0)
	}
}

func TestBlockString(t *testing.T) {
	b := NewBlockFactory().Get(1)
	assert.True(t, len(b.String()) > 0)
}

func TestBlockFactoryGet(t *testing.T) {
	f := NewBlockFactory()
	var nilBlock *Block = nil

	assert.Equal(t, nilBlock, f.Get(f.MinBlockNumber-1))
	assert.Equal(t, nilBlock, f.Get(f.MaxBlockNumber+1))
	for n := f.MinBlockNumber; n <= f.MaxBlockNumber; n++ {
		assert.True(t, f.Get(n) != nil)
	}

	// test that repeatedely returning the same block does not create a new instance
	var a *Block = f.Get(1)
	var b *Block = f.Get(1)
	assert.True(t, a == b, "References to block are not identical after repeated BlockFactory.Get() calls")
}

// Runs all block-factory functions and tests the blocks for
// consistency
func TestBlocks(t *testing.T) {
	blocks := NewBlockFactory().GetAll()

	for _, block := range blocks {

		// all shapes must have the same volume
		expVolume := block.Volume
		expZeros := block.Shapes[0].Count(0)
		for i, s := range block.Shapes {
			actVolume := s.Count(1)
			assert.Equal(t, expVolume, actVolume, "Volume of block shape %d is wrong", i)
			assert.Equal(t, expZeros, s.Count(0), "Volume of block shpae %d has the wrong number of zeros", i)
		}

		// all shapes must have the same size of a bounding box
		baseBox := block.Shapes[0].GetBoundingBox()
		boxSum := baseBox[0] + baseBox[1] + baseBox[2]
		boxProd := baseBox[0] * baseBox[1] * baseBox[2]
		for i, s := range block.Shapes {
			// if the sum and product of the dimensions match respectively, we can
			// deduct that the bounding boxes are identical and just rotated
			box := s.GetBoundingBox()
			sum := box[0] + box[1] + box[2]
			prod := box[0] * box[1] * box[2]
			assert.True(t, sum == boxSum && prod == boxProd,
				"Shapes[%d] of Block %d has the wrong bounding box (%d,%d,%d) (expect sum/prod=%d/%d, actual sum/prod=%d/%d)",
				i, block.Number, box[0], box[1], box[2], boxSum, boxProd, sum, prod)
		}
	}
}

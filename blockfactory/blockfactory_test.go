package blockfactory_test

import (
	"testing"
	"ubongo/block"
	. "ubongo/blockfactory"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	f := Get()
	var nilBlock *block.B = nil

	assert.Equal(t, nilBlock, f.ByNumber(f.MinBlockNumber-1))
	assert.Equal(t, nilBlock, f.ByNumber(f.MaxBlockNumber+1))
	for n := f.MinBlockNumber; n <= f.MaxBlockNumber; n++ {
		assert.True(t, f.ByNumber(n) != nil)
	}

	// test that repeatedely returning the same block does not create a new instance
	var a *block.B = f.ByNumber(1)
	var b *block.B = f.ByNumber(1)
	assert.True(t, a == b, "References to block are not identical after repeated BlockFactory.Get() calls")
}

func TestByNumber(t *testing.T) {
	f := Get()

	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		b := f.ByNumber(i)
		assert.NotNil(t, b)
		assert.Equal(t, i, b.Number)
	}

	var nilFactory *F = nil
	assert.Nil(t, nilFactory.ByNumber(1))
}

func TestByName(t *testing.T) {
	f := Get()

	for _, name := range []string{"hello", "big hook", "small hook", "gate"} {
		b := f.ByName(block.Yellow, name)
		assert.Equal(t, block.Yellow, b.Color)
		assert.Equal(t, name, b.Name)
	}
	for _, name := range []string{"big hook", "flash", "lighter", "v"} {
		b := f.ByName(block.Blue, name)
		assert.Equal(t, block.Blue, b.Color)
		assert.Equal(t, name, b.Name)
	}
	for _, name := range []string{"stool", "small hook", "big hook", "flash"} {
		b := f.ByName(block.Red, name)
		assert.Equal(t, block.Red, b.Color)
		assert.Equal(t, name, b.Name)
	}
	for _, name := range []string{"flash", "big hook", "T", "L"} {
		b := f.ByName(block.Green, name)
		assert.Equal(t, block.Green, b.Color)
		assert.Equal(t, name, b.Name)
	}

	assert.Nil(t, f.ByName(block.Blue, "Superman"))

	var nilFactory *F = nil
	assert.Nil(t, nilFactory.ByName(block.Blue, "lighter"))
}

func TestByVolume(t *testing.T) {
	f := Get()

	assert.Equal(t, 0, f.ByVolume(2).Count)
	assert.Equal(t, 1, f.ByVolume(3).Count)
	assert.Equal(t, 5, f.ByVolume(4).Count)
	assert.Equal(t, 10, f.ByVolume(5).Count)
	assert.Equal(t, 0, f.ByVolume(6).Count)

	var nilFactory *F = nil
	assert.Equal(t, 0, nilFactory.ByVolume(5).Count)
}

func TestGetAll(t *testing.T) {
	f := Get()

	blocks := f.GetAll()
	assert.Equal(t, 16, len(blocks))

	// check the blocks are unique
	blockMap := make(map[string]*block.B)
	for _, b := range blocks {
		blockMap[b.Color.String()+b.Name] = b
	}
	assert.Equal(t, 16, len(blockMap))
}

func TestGenerateBlocksets(t *testing.T) {
	f := Get()
	vol := 21
	blockCount := 5
	blocksets := f.GenerateBlocksets(vol, blockCount, 10000)

	// this test should return less than the requested number of results
	assert.LessOrEqual(t, len(blocksets), 10000)
	for _, s := range blocksets {
		assert.Equal(t, blockCount, s.Count)
		assert.Equal(t, vol, s.Volume())
	}

	// this test should return exactly 3 results
	blocksets = f.GenerateBlocksets(vol, blockCount, 3)
	assert.LessOrEqual(t, len(blocksets), 3)

	var nilFactory *F = nil
	assert.Equal(t, 0, len(nilFactory.GenerateBlocksets(3, 4, 99)))
}

func TestGenerateBlocksetsEmpty(t *testing.T) {
	f := Get()
	vol := 18
	blockCount := 5
	blocksets := f.GenerateBlocksets(vol, blockCount, 100)

	assert.Equal(t, 0, len(blocksets))
}

// Runs all block-factory functions and tests the blocks for
// consistency
func TestBlocks(t *testing.T) {
	blocks := Get().GetAll()

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

package blockfactory_test

import (
	"testing"
	"ubongo/block"
	. "ubongo/blockfactory"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	f := GetBlockFactory()
	var nilBlock *block.Block = nil

	assert.Equal(t, nilBlock, f.ByNumber(f.MinBlockNumber-1))
	assert.Equal(t, nilBlock, f.ByNumber(f.MaxBlockNumber+1))
	for n := f.MinBlockNumber; n <= f.MaxBlockNumber; n++ {
		assert.True(t, f.ByNumber(n) != nil)
	}

	// test that repeatedely returning the same block does not create a new instance
	var a *block.Block = f.ByNumber(1)
	var b *block.Block = f.ByNumber(1)
	assert.True(t, a == b, "References to block are not identical after repeated BlockFactory.Get() calls")
}

func TestByNumber(t *testing.T) {
	f := GetBlockFactory()

	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		b := f.ByNumber(i)
		assert.NotNil(t, b)
		assert.Equal(t, i, b.Number)
	}
}

func TestByName(t *testing.T) {
	f := GetBlockFactory()

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
}

func TestGetAll(t *testing.T) {
	f := GetBlockFactory()

	blocks := f.GetAll()
	assert.Equal(t, 16, len(blocks))

	// check the blocks are unique
	blockMap := make(map[string]*block.Block)
	for _, b := range blocks {
		blockMap[b.Color.String()+b.Name] = b
	}
	assert.Equal(t, 16, len(blockMap))
}

func TestGenerateBlocksets(t *testing.T) {
	f := GetBlockFactory()
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
}

func TestGenerateBlocksetsEmpty(t *testing.T) {
	f := GetBlockFactory()
	vol := 18
	blockCount := 5
	blocksets := f.GenerateBlocksets(vol, blockCount, 100)

	assert.Equal(t, 0, len(blocksets))
}

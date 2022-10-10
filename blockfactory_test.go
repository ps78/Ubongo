package main

import (
	"image"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockFactoryGet(t *testing.T) {
	f := GetBlockFactory()
	var nilBlock *Block = nil

	assert.Equal(t, nilBlock, f.ByNumber(f.MinBlockNumber-1))
	assert.Equal(t, nilBlock, f.ByNumber(f.MaxBlockNumber+1))
	for n := f.MinBlockNumber; n <= f.MaxBlockNumber; n++ {
		assert.True(t, f.ByNumber(n) != nil)
	}

	// test that repeatedely returning the same block does not create a new instance
	var a *Block = f.ByNumber(1)
	var b *Block = f.ByNumber(1)
	assert.True(t, a == b, "References to block are not identical after repeated BlockFactory.Get() calls")
}

func TestBlockFactoryByNumber(t *testing.T) {
	f := GetBlockFactory()

	for i := f.MinBlockNumber; i <= f.MaxBlockNumber; i++ {
		b := f.ByNumber(i)
		assert.NotNil(t, b)
		assert.Equal(t, i, b.Number)
	}
}

func TestBlockFactoryByName(t *testing.T) {
	f := GetBlockFactory()

	for _, name := range []string{"hello", "big hook", "small hook", "gate"} {
		b := f.ByName(Yellow, name)
		assert.Equal(t, Yellow, b.Color)
		assert.Equal(t, name, b.Name)
	}
	for _, name := range []string{"big hook", "flash", "lighter", "v"} {
		b := f.ByName(Blue, name)
		assert.Equal(t, Blue, b.Color)
		assert.Equal(t, name, b.Name)
	}
	for _, name := range []string{"stool", "small hook", "big hook", "flash"} {
		b := f.ByName(Red, name)
		assert.Equal(t, Red, b.Color)
		assert.Equal(t, name, b.Name)
	}
	for _, name := range []string{"flash", "big hook", "T", "L"} {
		b := f.ByName(Green, name)
		assert.Equal(t, Green, b.Color)
		assert.Equal(t, name, b.Name)
	}

	assert.Nil(t, f.ByName(Blue, "Superman"))
}

func TestBlockFactoryGetAll(t *testing.T) {
	f := GetBlockFactory()

	blocks := f.GetAll()
	assert.Equal(t, 16, len(blocks))

	// check the blocks are unique
	blockMap := make(map[string]*Block)
	for _, b := range blocks {
		blockMap[b.Color.String()+b.Name] = b
	}
	assert.Equal(t, 16, len(blockMap))
}

func TestBlockFactoryRenderAll(t *testing.T) {
	f := GetBlockFactory()
	dir, _ := os.MkdirTemp("./", "testing*")
	defer os.RemoveAll(dir)

	width := 500
	height := 400
	files := f.RenderAll(dir, width, height)
	assert.Equal(t, 16, len(files))
	for _, file := range files {
		infile, err := os.Open(file)
		assert.Nil(t, err)
		defer infile.Close()
		img, _, errPng := image.Decode(infile)
		assert.Nil(t, errPng)
		assert.Equal(t, width, img.Bounds().Dx())
		assert.Equal(t, height, img.Bounds().Dy())
	}
}

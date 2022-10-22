package block_test

import (
	"strings"
	"testing"

	"image/color"
	. "ubongo/block"
	"ubongo/blockfactory"

	"github.com/stretchr/testify/assert"
)

func TestBlockColorString(t *testing.T) {
	f := blockfactory.Get()
	for _, b := range f.GetAll().AsSlice() {
		colorName := strings.ToLower(b.Color.String())
		assert.NotEqual(t, "unknown", colorName)
	}
	assert.Equal(t, "unknown", strings.ToLower(BlockColor(99).String()))
}

func TestToRGBA(t *testing.T) {
	white := color.RGBA{255, 255, 255, 0}
	black := color.RGBA{0, 0, 0, 0}

	// we simply check that the standard block colors are not black or white
	assert.NotEqual(t, white, Red.ToRGBA())
	assert.NotEqual(t, black, Red.ToRGBA())
	assert.NotEqual(t, white, Green.ToRGBA())
	assert.NotEqual(t, black, Green.ToRGBA())
	assert.NotEqual(t, white, Blue.ToRGBA())
	assert.NotEqual(t, black, Blue.ToRGBA())
	assert.NotEqual(t, white, Yellow.ToRGBA())
	assert.NotEqual(t, black, Yellow.ToRGBA())

	// unkonwn block colors should map to white
	assert.Equal(t, white, BlockColor(99).ToRGBA())
}

func TestBlockString(t *testing.T) {
	b := blockfactory.Get().ByNumber(1)
	assert.True(t, len(b.String()) > 0)

	var nilBlock *B = nil
	assert.Equal(t, "(nil)", nilBlock.String())
}

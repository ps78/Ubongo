package graphics_test

import (
	"image"
	"os"
	"testing"
	"ubongo/blockfactory"
	"ubongo/card"
	"ubongo/cardfactory"
	"ubongo/game"
	. "ubongo/graphics"

	"github.com/stretchr/testify/assert"
)

func TestGameSolutionCreateImage(t *testing.T) {
	cf := cardfactory.GetCardFactory()
	p := cf.Get(card.Difficult, 3).Problems[7]
	g := game.NewGame(p)
	solutions := g.Solve()

	width := 400
	height := 300

	img := CreateImage(solutions[0], width, height, 0, 0, 0, 0.1)

	assert.NotNil(t, img)
	assert.Equal(t, width, img.Bounds().Dx())
	assert.Equal(t, height, img.Bounds().Dy())
}

func TestRenderAll(t *testing.T) {
	f := blockfactory.GetBlockFactory()
	dir, _ := os.MkdirTemp("./", "testing*")
	defer os.RemoveAll(dir)

	width := 500
	height := 400
	minBlackRatio := 0.8
	maxBlackRatio := 0.98

	// read back the images that were created
	files := RenderAll(f.GetAll(), dir, width, height)
	assert.Equal(t, 16, len(files))
	for _, file := range files {
		infile, err := os.Open(file)
		assert.Nil(t, err)
		defer infile.Close()
		img, _, errPng := image.Decode(infile)

		assert.Nil(t, errPng)
		assert.Equal(t, width, img.Bounds().Dx())
		assert.Equal(t, height, img.Bounds().Dy())

		// the ratio of black pixels should be in the given range
		blackRatio := getPixelRatio(img, 0, 0, 0)
		assert.True(t, blackRatio >= minBlackRatio && blackRatio <= maxBlackRatio)
	}
}

// Returns the ratio of pixels that have the given color.
// Value between 0 and 1
func getPixelRatio(img image.Image, red, green, blue uint32) float64 {
	h := img.Bounds().Dy()
	w := img.Bounds().Dx()
	counter := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r == red && g == green && b == blue {
				counter++
			}
		}
	}
	return float64(counter) / float64(w*h)
}

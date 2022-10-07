package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameSolutionCreateImage(t *testing.T) {
	cf := GetCardFactory()
	p := cf.Get(Difficult, 3).Problems[7]
	g := NewGame(p)
	solutions := g.Solve()

	width := 400
	height := 300

	img := solutions[0].CreateImage(width, height, 0, 0, 0, 0.1)

	assert.NotNil(t, img)
	assert.Equal(t, width, img.Bounds().Dx())
	assert.Equal(t, height, img.Bounds().Dy())
}

func TestSaveAsPng(t *testing.T) {
	/*
		f, _ := os.CreateTemp("", "*")
		f.Name()

		SaveAsPng()
	*/
}

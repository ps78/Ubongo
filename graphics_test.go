package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSolutionImage(t *testing.T) {
	cf := GetCardFactory()
	p := cf.Get(Difficult, 3).Problems[7]
	g := NewGame(p)
	solutions := g.Solve()

	img := GetSolutionImage(solutions[0], 400, 300, 0, 0, 0)

	assert.NotNil(t, img)
	assert.Equal(t, 400, img.Bounds().Dx())
	assert.Equal(t, 300, img.Bounds().Dy())
}

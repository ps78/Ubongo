package main

import (
	"testing"
)

func TestTryAddBlock(t *testing.T) {
	area := Array2d{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}}
	g := NewGame(area, 2)
	shape := CreateBlock08().Shapes[0]
	pos := Vector{0, 0, 0}

	if !g.TryAddBlock(shape, pos) {
		t.Errorf(("TryAddBlock returned no success where it should"))
	}

	exp := [][][]int8{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}}
	if !Equal3DArray(exp, g.Volume) {
		t.Errorf("The resulting volume after TryAddBlock is not as expected")
	}
}

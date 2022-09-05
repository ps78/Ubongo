package main

import (
	"testing"
)

func arrayContainsVector(vectorArray []Vector, v Vector) bool {
	if vectorArray == nil {
		return false
	}
	for _, vec := range vectorArray {
		if vec == v {
			return true
		}
	}
	return false
}

// Runs all block-factory functions and tests the blocks for
// consistency
func TestBlocks(t *testing.T) {
	blockFactory := []BlockFactoryFunc{MakeBlock08}

	for _, f := range blockFactory {
		block := f()

		// all shapes must have the same volume
		expVolume := block.Volume
		for i, s := range block.Shapes {
			actVolume := GetBlockVolume(s)
			if actVolume != expVolume {
				t.Errorf("Shapes[%d] of Block %d has the wrong volume (%d instead of %d)",
					i, block.Number, actVolume, expVolume)
			}
		}

		// all shapes must have the same size of a bounding box
		baseBox := GetBoundingBoxFromBlockShape(block.Shapes[0])
		boxSum := baseBox[0] + baseBox[1] + baseBox[2]
		boxProd := baseBox[0] * baseBox[1] * baseBox[2]
		for i, s := range block.Shapes {
			box := GetBoundingBoxFromBlockShape(s)
			sum := box[0] + box[1] + box[2]
			prod := box[0] * box[1] * box[2]
			if sum != boxSum || prod != boxProd {
				t.Errorf("Shapes[%d] of Block %d has the wrong bounding box (%d,%d,%d) (expect sum/prod=%d/%d, actual sum/prod=%d/%d)",
					i, block.Number, box[0], box[1], box[2], boxSum, boxProd, sum, prod)
			}
		}
	}
}

func TestGetShiftVectors(t *testing.T) {
	outer := BoundingBox{4, 3, 2}
	inner := BoundingBox{3, 2, 1}
	shifts := outer.GetShiftVectors(inner)
	if len(shifts) != 8 {
		t.Errorf("GetShiftVectors return %d results, expected %d", len(shifts), 8)
	}

	expected := []Vector{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1}}

	for _, v := range expected {
		if !arrayContainsVector(shifts, v) {
			t.Errorf("Vector %s is missing in output of GetShiftVector()", v)
		}
	}
}

func TestGetShiftVectorsEmpty(t *testing.T) {
	outer := BoundingBox{4, 3, 2}
	inner := BoundingBox{5, 2, 1}
	shifts := outer.GetShiftVectors(inner)
	if len(shifts) != 0 {
		t.Errorf("GetShiftVectors did not return an empty slice")
	}
}

func TestTryAdd(t *testing.T) {
	p := MakeProblem("B12", 1, ProblemShape{{1, 1, 0, 0}, {0, 1, 1, 0}, {1, 1, 1, 1}}, []*Block{})
	vol := p.CreateVolume()
	block := MakeBlock08()

	success, newVol := TryAdd(vol, block.Shapes[0], Vector{0, 0, 0})
	if !success {
		t.Errorf(("TryAdd returned no success"))
	}
	t.Logf("%v", newVol)
}

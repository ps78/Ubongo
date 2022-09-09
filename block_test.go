package main

import (
	"testing"
)

// Runs all block-factory functions and tests the blocks for
// consistency
func TestBlocks(t *testing.T) {
	blockFactory := []BlockFactoryFunc{CreateBlock08}

	for _, f := range blockFactory {
		block := f()

		// all shapes must have the same volume
		expVolume := block.Volume
		for i, s := range block.Shapes {
			actVolume := CountValues3D(s, 1)
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
			// if the sum and product of the dimensions match respectively, we can
			// deduct that the bounding boxes are identical and just rotated
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

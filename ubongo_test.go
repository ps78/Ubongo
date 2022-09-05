package main

import "testing"

// Runs all block-factory functions and tests the blocks for
// consistency
func TestBlocks(t *testing.T) {
	blockFactory := []BlockFactoryFunc{MakeBlock08}

	for _, f := range blockFactory {
		block := f()

		// all shapes must have the same volume
		expVolume := block.NumCubes
		for i, s := range block.Shapes {
			actVolume := GetBlockVolume(s)
			if actVolume != expVolume {
				t.Errorf("Shapes[%d] of Block %d has the wrong volume (%d instead of %d)",
					i, block.Number, actVolume, expVolume)
			}
		}
	}
}

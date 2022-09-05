package main

import "fmt"

// BoundingBox describes a box with 3 dimensions expressed in unit-squares
type BoundingBox [3]int

// Vector represents a 3-dimensional int-vector
type Vector [3]int

func (v Vector) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v[0], v[1], v[2])
}

// GetShiftVectors returns all possible placements of the inner bounding box
// inside the outer bounding box.
// Returns an empty slice if inner does not fit into outer at all
func (outer BoundingBox) GetShiftVectors(inner BoundingBox) []Vector {
	delta := []int{outer[0] - inner[0], outer[1] - inner[1], outer[2] - inner[2]}
	n := (delta[0] + 1) * (delta[1] + 1) * (delta[2] + 1)

	// return empty vector if there is not fit
	if n <= 0 {
		return make([]Vector, 0)
	} else {
		shifts := make([]Vector, n)
		i := 0
		for x := 0; x <= delta[0]; x++ {
			for y := 0; y <= delta[1]; y++ {
				for z := 0; z <= delta[2]; z++ {
					shifts[i] = [3]int{x, y, z}
					i++
				}
			}
		}
		return shifts
	}
}

func TryAdd(vol *ProblemVolume, block BlockShape, shift Vector) (bool, *ProblemVolume) {
	result := CopyVolume(vol)
	box := GetBoundingBoxFromBlockShape(block)
	for x := shift[0]; x < box[0]+shift[0]; x++ {
		for y := shift[1]; y < box[1]+shift[1]; y++ {
			for z := shift[2]; z < box[2]+shift[2]; z++ {
				if block[x-shift[0]][y-shift[1]][z-shift[2]] == 1 {
					// if space is part of volume and empty -> ok
					if (*vol)[x][y][z] == 1 {
						(*result)[x][y][z] = 2 // mark space as occupied
					} else {
						return false, nil // otherwise abort
					}
				}
			}
		}
	}
	return true, result
}

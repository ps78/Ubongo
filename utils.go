package main

import (
	"fmt"
)

// Vector represents a 3-dimensional int-vector
type Vector [3]int

type Vectorf [3]float64

// Returns a String representation of a vector
func (v Vector) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v[0], v[1], v[2])
}

func (v Vectorf) String() string {
	return fmt.Sprintf("(%f,%f,%f)", v[0], v[1], v[2])
}

func (a Vector) Float64() Vectorf {
	return Vectorf{float64(a[0]), float64(a[1]), float64(a[2])}
}

func (a Vectorf) Add(b Vectorf) Vectorf {
	return Vectorf{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func (a Vectorf) Sub(b Vectorf) Vectorf {
	return Vectorf{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a Vectorf) Div(b float64) Vectorf {
	return Vectorf{a[0] / b, a[1] / b, a[2] / b}
}

func (a Vectorf) Mult(b float64) Vectorf {
	return Vectorf{a[0] * b, a[1] * b, a[2] * b}
}

// GetShiftVectors returns all possible placements of the inner bounding box
// inside the outer bounding box.
// Returns an empty slice if inner does not fit into outer at all
func (outerBoundingBox Vector) GetShiftVectors(innerBoundingBox Vector) []Vector {
	delta := Vector{
		outerBoundingBox[0] - innerBoundingBox[0],
		outerBoundingBox[1] - innerBoundingBox[1],
		outerBoundingBox[2] - innerBoundingBox[2]}

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

// Looks for a in lst (comparing the values)
func FindArray3d(lst []*Array3d, a *Array3d) (bool, int) {
	if lst != nil && a != nil {
		for i, arr := range lst {
			if a.IsEqual(arr) {
				return true, i
			}
		}
	}
	return false, -1
}

// CreateParitions returns all possible partition given the parameters:
//
//	n         : The number to partition
//	parts     : slice of allowed integers to use as parts of the partitions
//	maxCounts : map defining the maximum multiplier of each part in the partition
//	partLen   : the length of the partition, i.e. the total number of elements
func CreateParitions(n int, parts []int, maxCounts map[int]int, partLen int) [](map[int]int) {
	solutions := make([](map[int]int), 0)

	createPartitionsRecurisve(n, &parts, 0, &maxCounts, partLen, nil, &solutions)

	return solutions
}

// Recursive function called by CreatePartitions. Don't use directly
func createPartitionsRecurisve(nRemainder int, parts *[]int, partIdx int, maxCounts *map[int]int, partLen int, curSolution *map[int]int, solutions *[](map[int]int)) {
	if curSolution == nil {
		s := make(map[int]int)
		curSolution = &s
	}

	part := (*parts)[partIdx]

	// if the next partition number is too big, jump to next partition
	if part > nRemainder && partIdx+1 < len(*parts) {
		createPartitionsRecurisve(nRemainder, parts, partIdx+1, maxCounts, partLen, curSolution, solutions)
	} else {
		maxCount := nRemainder / part
		if maxCount > (*maxCounts)[part] {
			maxCount = (*maxCounts)[part]
		}
		for count := 0; count <= maxCount; count++ {
			(*curSolution)[part] = count

			newRemainder := nRemainder - part*count

			// remainder is zero, we found a solution
			if newRemainder == 0 {
				// check the lenght of the partition
				c := 0
				for _, v := range *curSolution {
					c += v
				}
				if c <= partLen {
					newSolution := make(map[int]int)
					for k, v := range *curSolution {
						newSolution[k] = v
					}
					*solutions = append(*solutions, newSolution)
				}

				// continue recursion
			} else if newRemainder > 0 && partIdx+1 < len(*parts) {
				createPartitionsRecurisve(newRemainder, parts, partIdx+1, maxCounts, partLen, curSolution, solutions)
			}
			(*curSolution)[part] = 0
		}
	}
}

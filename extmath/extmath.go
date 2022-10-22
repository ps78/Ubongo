// Package extmath contains extended math functions (e.g. partitioning of integers)
package extmath

// CreateParitions returns all possible partition given the parameters:
//	n         : The number to partition
//	parts     : slice of allowed integers to use as parts of the partitions
//	maxCounts : map defining the maximum multiplier of each part in the partition
//	partLen   : the length of the partition, i.e. the total number of elements
func CreateParitions(n int, parts []int, maxCounts map[int]int, partLen int) [](map[int]int) {
	solutions := make([](map[int]int), 0)

	createPartitionsRecurisve(n, &parts, 0, &maxCounts, partLen, nil, &solutions)

	return solutions
}

// createPartitionsRecursive is a recursive function called by CreatePartitions. Don't use directly
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
				if c == partLen {
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

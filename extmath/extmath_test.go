package extmath_test

import (
	"testing"
	. "ubongo/extmath"

	"github.com/stretchr/testify/assert"
)

func TestCreatePartitions(t *testing.T) {
	n := 21
	parts := []int{5, 4, 3}
	partLen := 5
	maxCounts := map[int]int{3: 1, 4: 5, 5: 10}
	partitions := CreateParitions(n, parts, maxCounts, partLen)

	assert.Equal(t, 2, len(partitions))
	for _, part := range partitions {
		sum := 0
		count := 0
		for k, v := range part {
			sum += k * v
			count += v
		}
		assert.Equal(t, n, sum)
		assert.Equal(t, partLen, count)
	}
}

func TestCreatePartitionsNoResult(t *testing.T) {
	n := 18
	parts := []int{5, 4, 3}
	partLen := 5
	maxCounts := map[int]int{3: 1, 4: 5, 5: 10}
	partitions := CreateParitions(n, parts, maxCounts, partLen)

	assert.Equal(t, 0, len(partitions))
}

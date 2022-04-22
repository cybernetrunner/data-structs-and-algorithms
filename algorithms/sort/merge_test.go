package sort_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"data-structs-and-algorithms/algorithms/sort"
)

func Test_MergeSort(t *testing.T) {
	nums := []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}
	res := []int{2, 4, 6, 11, 13, 15, 18, 19, 21, 71}
	ress := sort.MergeSort(nums)

	assert.ElementsMatch(t, ress, res)
}
package minimum_swaps_to_sort_an_array_test

import (
	"testing"

	sut "github.com/smoothprogrammer/minmax/geeksforgeeks/minimum-swaps-to-sort-an-array"
	"github.com/smoothprogrammer/testr"
)

func TestMinimumSwapsToSortAnArray(t *testing.T) {
	tests := map[string]struct {
		arr   []int
		swaps int
	}{
		"zero":       {nil, 0},
		"one":        {[]int{5}, 0},
		"ascending":  {[]int{1, 2, 3, 4, 5}, 0},
		"descending": {[]int{5, 4, 3, 2, 1}, 2},
		"one cycle":  {[]int{5, 1, 2, 3, 4}, 4},
		"many cycle": {[]int{2, 5, 6, 3, 1, 4, 8, 7}, 5},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			swaps := sut.Solution(tc.arr)
			assert.Equal(swaps, tc.swaps)
		})
	}
}

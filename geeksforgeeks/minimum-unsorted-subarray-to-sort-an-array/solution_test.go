package minimum_unsorted_subarray_to_sort_an_array_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/minimum-unsorted-subarray-to-sort-an-array"
	"github.com/minizilla/testr"
)

func TestMinimumUnsortedSubarrayToSortAnArray(t *testing.T) {
	tests := map[string]struct {
		arr  []int
		s, e int
	}{
		"zero":       {nil, 0, 0},
		"one":        {[]int{0}, 0, 0},
		"ascending":  {[]int{1, 2, 3, 4, 5}, 0, 0},
		"descending": {[]int{5, 4, 3, 2, 1}, 0, 4},
		"x":          {[]int{1, 2, 7, 3, 4, 9, 8, 10, 11}, 2, 6},
		"y":          {[]int{1, 3, 7, 2, 4, 6, 10, 8, 9, 11}, 1, 8},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			s, e := sut.Solution(tc.arr)
			assert.Equal(s, tc.s)
			assert.Equal(e, tc.e)
		})
	}
}

package minimum_and_maximum_array_elements_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/minimum-and-maximum-array-elements"
	"github.com/minizilla/testr"
)

func TestMinimumAndMaximumArrayElements(t *testing.T) {
	tests := map[string]struct {
		arr      []int
		min, max int
	}{
		"zero":       {nil, 0, 0},
		"one":        {[]int{5}, 5, 5},
		"ascending":  {[]int{5, 7, 8, 10, 12}, 5, 12},
		"descending": {[]int{12, 10, 8, 7, 5}, 5, 12},
		"random":     {[]int{7, 5, 12, 8, 10}, 5, 12},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			min, max := sut.Solution(tc.arr)
			assert.Equal(min, tc.min)
			assert.Equal(max, tc.max)
		})
	}
}

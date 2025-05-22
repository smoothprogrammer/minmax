package chocolate_distribution_problem_test

import (
	"testing"

	sut "github.com/smoothprogrammer/minmax/geeksforgeeks/chocolate-distribution-problem"
	"github.com/smoothprogrammer/testr"
)

func TestChocolateDistributionProblem(t *testing.T) {
	tests := map[string]struct {
		arr     []int
		m       int
		mindiff int
	}{
		"zero":       {nil, 0, 0},
		"one":        {[]int{1}, 1, 0},
		"ascending":  {[]int{1, 2, 3, 4, 5}, 3, 2},
		"descending": {[]int{5, 4, 3, 2, 1}, 3, 2},
		"random":     {[]int{25, 26, 23, 31, 11, 29, 31, 34}, 5, 6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			mindiff := sut.Solution(tc.arr, tc.m)
			assert.Equal(mindiff, tc.mindiff)
		})
	}
}

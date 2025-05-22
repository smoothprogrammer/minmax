package rotate_matrix_90_degree_without_extra_space_test

import (
	"testing"

	sut "github.com/smoothprogrammer/minmax/geeksforgeeks/rotate-matrix-90-degree-without-extra-space"
	"github.com/smoothprogrammer/testr"
)

func TestRotateMatrix90DegreeWithoutExtraSpace(t *testing.T) {
	tests := map[string]struct {
		in, out [][]int
	}{
		"zero": {nil, nil},
		"1x1": {
			[][]int{
				{1},
			},
			[][]int{
				{1},
			},
		},
		"2x2": {
			[][]int{
				{1, 2},
				{3, 4},
			},
			[][]int{
				{2, 4},
				{1, 3},
			},
		},
		"3x3": {
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			},
		},
		"4x4": {
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			[][]int{
				{4, 8, 12, 16},
				{3, 7, 11, 15},
				{2, 6, 10, 14},
				{1, 5, 9, 13},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			out := sut.Solution(tc.in)
			assert.Equal(out, tc.out)
		})
	}
}

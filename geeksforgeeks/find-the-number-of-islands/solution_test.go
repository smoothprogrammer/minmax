package find_the_number_of_islands_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/find-the-number-of-islands"
	"github.com/minizilla/testr"
)

func TestFindTheNumberOfIslands(t *testing.T) {
	tests := map[string]struct {
		matrix  [][]int
		islands int
	}{
		"zero": {nil, 0},
		"one": {
			[][]int{
				{0, 1},
				{1, 0},
			}, 1,
		},
		"two": {
			[][]int{
				{0, 1, 0, 0, 0},
				{1, 0, 0, 1, 1},
			}, 2,
		},
		"three": {
			[][]int{
				{1, 0, 1, 1},
				{0, 1, 0, 1},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{1, 0, 0, 0},
			}, 3,
		},
		"four": {
			[][]int{
				{1, 0, 1, 0, 1, 0, 0, 1, 0},
				{0, 1, 0, 1, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 1, 0},
				{1, 1, 0, 0, 0, 1, 0, 1, 0},
			}, 4,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			islands := sut.Solution(tc.matrix)
			assert.Equal(islands, tc.islands)
		})
	}
}

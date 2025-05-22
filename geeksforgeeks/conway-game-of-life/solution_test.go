package conway_game_of_life_test

import (
	"testing"

	sut "github.com/smoothprogrammer/minmax/geeksforgeeks/conway-game-of-life"
	"github.com/smoothprogrammer/testr"
)

func TestConwayGameOfLife(t *testing.T) {
	tests := map[string]struct {
		in, out [][]int
	}{
		"zero": {nil, nil},
		"live<2": {
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
		"live=2||3": {
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			[][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
		"live>3": {
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 1, 0, 0},
				{0, 1, 1, 0, 0},
				{0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0},
			},
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
		"dead=3": {
			[][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
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

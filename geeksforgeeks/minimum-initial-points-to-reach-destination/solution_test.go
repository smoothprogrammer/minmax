package minimum_initial_points_to_reach_destination_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/minimum-initial-points-to-reach-destination"
	"github.com/minizilla/testr"
)

func TestMinimumInitialPointsToReachDestination(t *testing.T) {
	tests := map[string]struct {
		matrix [][]int
		points int
	}{
		"zero": {nil, 0},
		"one": {
			[][]int{
				{1},
			}, 1,
		},
		"two": {
			[][]int{
				{-1, -1},
				{1, -2},
			}, 3,
		},
		"negative": {
			[][]int{
				{-1, -1, -10, -10},
				{-1, -1, -10, -10},
				{-1, -1, -10, -10},
				{-1, -1, -10, -10},
				{-1, -1, -10, -10},
				{1, -1, -10, -10},
			}, 26,
		},
		"positive": {
			[][]int{
				{1, 1, 10, 10},
				{1, 1, 10, 10},
				{1, 1, 10, 10},
				{1, 1, 10, 10},
				{1, 1, 10, 10},
				{1, 1, 10, 10},
			}, 1,
		},
		"random": {
			[][]int{
				{-1, -10, -1, 10},
				{-10, -1, 0, 10},
				{1, -1, -10, 10},
			}, 13,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			points := sut.Solution(tc.matrix)
			assert.Equal(points, tc.points)
		})
	}
}

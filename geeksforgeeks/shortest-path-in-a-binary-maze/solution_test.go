package shortest_path_in_a_binary_maze_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/shortest-path-in-a-binary-maze"
	"github.com/minizilla/testr"
)

func TestShortestPathInABinaryMaze(t *testing.T) {
	tests := map[string]struct {
		maze         [][]int
		src, dst     sut.Point
		shortestPath int
	}{
		"nil": {nil, sut.Point{0, 0}, sut.Point{0, 0}, 0},
		"invalid point": {
			[][]int{
				{0, 0},
			},
			sut.Point{-1, 0},
			sut.Point{0, 2},
			0,
		},
		"zero src": {
			[][]int{
				{0, 1},
			},
			sut.Point{0, 0},
			sut.Point{0, 1},
			0,
		},
		"zero dst": {
			[][]int{
				{1, 0},
			},
			sut.Point{0, 0},
			sut.Point{0, 1},
			0,
		},
		"unconnected point": {
			[][]int{
				{1, 0, 0, 0},
				{1, 1, 0, 0},
				{1, 0, 1, 1},
			},
			sut.Point{0, 0},
			sut.Point{2, 3},
			0,
		},
		"connected point": {
			[][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 0, 0, 1},
				{1, 1, 1, 1},
			},
			sut.Point{0, 0},
			sut.Point{4, 2},
			6,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			shortestPath := sut.Solution(tc.maze, tc.src, tc.dst)
			assert.Equal(shortestPath, tc.shortestPath)
		})
	}
}

package tic_tac_toe_validity_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/tic-tac-toe-validity"
	"github.com/minizilla/testr"
)

func TestTicTacToeValidity(t *testing.T) {
	tests := map[string]struct {
		board   [][]rune
		isValid bool
	}{
		"nil": {nil, false},
		"empty": {[][]rune{
			{' ', ' ', ' '},
			{' ', ' ', ' '},
			{' ', ' ', ' '},
		}, true},
		"O first": {[][]rune{
			{'X', 'O', 'X'},
			{'X', 'O', 'O'},
			{'O', 'X', 'O'},
		}, false},
		"double win": {[][]rune{
			{'O', 'O', 'X'},
			{'O', 'X', 'X'},
			{'O', 'X', 'X'},
		}, false},
		"double win diagonal": {[][]rune{
			{'X', 'O', 'X'},
			{'O', 'X', 'O'},
			{'X', 'O', 'X'},
		}, false},
		"double X win": {[][]rune{
			{'O', 'O', 'X'},
			{'O', 'O', 'X'},
			{'X', 'X', 'X'},
		}, false},
		"O win invalid": {[][]rune{
			{'O', 'X', ' '},
			{'O', 'X', 'X'},
			{'O', ' ', 'X'},
		}, false},
		"X win invalid": {[][]rune{
			{'X', ' ', ' '},
			{'X', 'O', 'O'},
			{'X', ' ', 'O'},
		}, false},
		"draw": {[][]rune{
			{'X', 'O', 'X'},
			{'X', 'X', 'O'},
			{'O', 'X', 'O'},
		}, true},
		"X win horizontal": {[][]rune{
			{'X', 'X', 'X'},
			{' ', 'O', 'O'},
			{' ', ' ', ' '},
		}, true},
		"O win horizontal": {[][]rune{
			{'O', 'O', 'O'},
			{' ', 'X', 'X'},
			{' ', ' ', 'X'},
		}, true},
		"X win vertical": {[][]rune{
			{'X', ' ', ' '},
			{'X', 'O', 'O'},
			{'X', ' ', ' '},
		}, true},
		"O win vertical": {[][]rune{
			{'O', ' ', ' '},
			{'O', 'X', 'X'},
			{'O', ' ', 'X'},
		}, true},
		"X win diagonal": {[][]rune{
			{'X', ' ', ' '},
			{'O', 'X', 'O'},
			{' ', ' ', 'X'},
		}, true},
		"O win diagonal": {[][]rune{
			{'O', ' ', ' '},
			{'X', 'O', 'X'},
			{'X', ' ', 'O'},
		}, true},
		"X win anti diagonal": {[][]rune{
			{' ', ' ', 'X'},
			{'O', 'X', 'O'},
			{'X', ' ', ' '},
		}, true},
		"O win anti diagonal": {[][]rune{
			{' ', ' ', 'O'},
			{'X', 'O', 'X'},
			{'O', 'X', ' '},
		}, true},
		"5x5": {[][]rune{
			{'X', ' ', 'O', ' ', ' '},
			{' ', 'X', 'O', ' ', ' '},
			{'X', ' ', 'O', ' ', ' '},
			{' ', 'X', 'O', ' ', ' '},
			{'X', ' ', 'O', ' ', ' '},
		}, true},
		"7x7": {[][]rune{
			{' ', ' ', 'O', ' ', ' ', 'X', ' '},
			{' ', ' ', 'O', ' ', ' ', ' ', 'X'},
			{' ', ' ', 'O', ' ', ' ', 'X', ' '},
			{' ', ' ', 'O', ' ', ' ', ' ', 'X'},
			{' ', ' ', 'O', ' ', ' ', 'X', ' '},
			{' ', ' ', 'O', ' ', ' ', ' ', 'X'},
			{' ', ' ', 'O', ' ', ' ', 'X', ' '},
		}, true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			isValid := sut.Solution(tc.board)
			assert.Equal(isValid, tc.isValid)
		})
	}
}

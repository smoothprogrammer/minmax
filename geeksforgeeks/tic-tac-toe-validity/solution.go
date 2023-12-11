package tic_tac_toe_validity

const (
	emptyMark = ' '
	xMark     = 'X'
	oMark     = 'O'
)

func Solution(board [][]rune) bool {
	n := len(board)
	if n == 0 {
		return false
	}

	nx, no, wx, wo := countWinMarks(board)

	if nx < no || nx > no+1 {
		return false
	}

	if wx+wo > 1 {
		return false
	}

	if wx == 1 {
		return nx == no+1
	}

	if wo == 1 {
		return no == nx
	}

	return true
}

type Counter struct {
	X, O       int
	WinX, WinO int
}

func (counter *Counter) add(mark rune) {
	switch mark {
	case xMark:
		counter.X++
	case oMark:
		counter.O++
	}
}

func countWinMarks(board [][]rune) (nx, no, wx, wo int) {
	n := len(board)
	var diagonal, antiDiagonal Counter
	for i := 0; i < n; i++ {
		var row Counter
		var col Counter
		for j := 0; j < n; j++ {
			row.add(board[i][j])
			col.add(board[j][i])
		}

		nx += row.X
		no += row.O

		if row.X == n {
			wx++
		}
		if row.O == n {
			wo++
		}
		if col.X == n {
			wx++
		}
		if col.O == n {
			wo++
		}

		diagonal.add(board[i][i])
		antiDiagonal.add(board[i][n-1-i])
	}

	if diagonal.X == n {
		wx++
	}
	if diagonal.O == n {
		wo++
	}
	if antiDiagonal.X == n {
		wx++
	}
	if antiDiagonal.O == n {
		wo++
	}

	return
}

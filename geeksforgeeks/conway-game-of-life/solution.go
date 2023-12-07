package conway_game_of_life

func Solution(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	out := make([][]int, 0, len(matrix))
	for i, rows := range matrix {
		outRows := make([]int, 0, len(rows))
		for j, cell := range rows {
			neighbors := calculateNeighbors(matrix, i, j)
			switch {
			case underpopulation(matrix[i][j], neighbors):
				cell = 0
			case overpoulation(matrix[i][j], neighbors):
				cell = 0
			case reproduction(matrix[i][j], neighbors):
				cell = 1
			}
			outRows = append(outRows, cell)
		}
		out = append(out, outRows)
	}

	return out
}

func calculateNeighbors(matrix [][]int, r, c int) int {
	var out int
	row, col := len(matrix), len(matrix[r])
	for i := r - 1; i <= r+1; i++ {
		if i < 0 || i >= row {
			continue
		}
		for j := c - 1; j <= c+1; j++ {
			if j < 0 || j >= col || (i == r && j == c) {
				continue
			}
			out += matrix[i][j]
		}
	}
	return out
}

func underpopulation(cell, neighbors int) bool { return cell == 1 && neighbors < 2 }
func overpoulation(cell, neighbors int) bool   { return cell == 1 && neighbors > 3 }
func reproduction(cell, neighbors int) bool    { return cell == 0 && neighbors == 3 }

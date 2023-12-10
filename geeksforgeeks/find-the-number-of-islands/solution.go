package find_the_number_of_islands

func Solution(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	var islands int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 1 {
				islands++
				dfs(matrix, rows, cols, r, c)
			}
		}
	}

	return islands
}

func dfs(matrix [][]int, rows, cols, r, c int) {
	if r < 0 || r >= rows ||
		c < 0 || c >= cols ||
		matrix[r][c] == 0 {
		return
	}

	matrix[r][c] = 0
	dfs(matrix, rows, cols, r-1, c-1)
	dfs(matrix, rows, cols, r-1, c)
	dfs(matrix, rows, cols, r-1, c+1)
	dfs(matrix, rows, cols, r, c-1)
	dfs(matrix, rows, cols, r, c+1)
	dfs(matrix, rows, cols, r+1, c-1)
	dfs(matrix, rows, cols, r+1, c)
	dfs(matrix, rows, cols, r+1, c+1)
}

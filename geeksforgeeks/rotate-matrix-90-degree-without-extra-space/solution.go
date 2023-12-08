package rotate_matrix_90_degree_without_extra_space

func Solution(m [][]int) [][]int {
	n := len(m)
	if n <= 1 {
		return m
	}

	for i := 0; i < n/2+n%2; i++ {
		for j := i; j < n-1-i; j++ {
			ar, br, cr, dr := j, n-1-i, n-1-j, i
			ac, bc, cc, dc := i, j, n-1-i, n-1-j
			m[ar][ac], m[br][bc], m[cr][cc], m[dr][dc] = m[dr][dc], m[ar][ac], m[br][bc], m[cr][cc]
		}
	}

	return m
}

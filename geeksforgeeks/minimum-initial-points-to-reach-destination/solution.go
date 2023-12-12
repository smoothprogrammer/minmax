package minimum_initial_points_to_reach_destination

func Solution(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])

	var minPoints int
	initialCell := cell{points: matrix[0][0]}
	initialCell.setLowestPoints()

	queue := []cell{initialCell}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.row == m-1 && curr.col == n-1 {
			currMinPoints := curr.minPoints()
			// fmt.Println(curr.lowestPoints)
			if minPoints == 0 {
				minPoints = currMinPoints
			}
			if currMinPoints < minPoints {
				minPoints = currMinPoints
			}
			continue
		}

		neighbors := []cell{
			{row: curr.row, col: curr.col + 1},
			{row: curr.row + 1, col: curr.col},
		}

		for _, neighbor := range neighbors {
			if neighbor.row >= m || neighbor.col >= n {
				continue
			}

			neighbor.points = curr.points + matrix[neighbor.row][neighbor.col]
			neighbor.lowestPoints = curr.lowestPoints
			neighbor.setLowestPoints()
			queue = append(queue, neighbor)
		}
	}

	return minPoints
}

type cell struct {
	row, col             int
	lowestPoints, points int
}

func (c *cell) setLowestPoints() {
	if c.points < c.lowestPoints {
		c.lowestPoints = c.points
	}
}

func (c *cell) minPoints() int {
	if c.lowestPoints >= 0 {
		return 1
	}
	return 1 - c.lowestPoints
}

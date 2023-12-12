package shortest_path_in_a_binary_maze

type Point struct {
	Row, Col int
}

func Solution(maze [][]int, src, dst Point) int {
	m := len(maze)
	if m == 0 {
		return 0
	}
	n := len(maze[0])

	if src.Row < 0 || src.Row >= m || src.Col < 0 || src.Col >= n ||
		dst.Row < 0 || dst.Row >= m || dst.Col < 0 || dst.Col >= n {
		return 0
	}

	if maze[src.Row][src.Col] == 0 || maze[dst.Row][dst.Col] == 0 {
		return 0
	}

	return bfs(maze, src, dst)
}

type cell struct {
	Point
	distance int
}

func bfs(maze [][]int, src, dst Point) int {
	m, n := len(maze), len(maze[0])
	queue := []cell{{src, 0}}
	var shortestPath int

	for len(queue) > 0 {
		curr := queue[0]
		if curr.Point == dst {
			shortestPath = curr.distance
			break
		}

		queue = queue[1:]
		maze[curr.Row][curr.Col] = 0

		neighbors := []Point{
			{curr.Row - 1, curr.Col},
			{curr.Row, curr.Col + 1},
			{curr.Row + 1, curr.Col},
			{curr.Row, curr.Col - 1},
		}

		for _, neighbor := range neighbors {
			if neighbor.Row < 0 || neighbor.Row >= m ||
				neighbor.Col < 0 || neighbor.Col >= n ||
				maze[neighbor.Row][neighbor.Col] == 0 {
				continue
			}

			queue = append(queue, cell{neighbor, curr.distance + 1})
		}
	}

	return shortestPath
}

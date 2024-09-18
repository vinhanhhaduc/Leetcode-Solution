type Point struct {
	x, y int
}
var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func check(x, y, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func bfs(forest [][]int, start, target Point) int {
	m, n := len(forest), len(forest[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := list.New()
	queue.PushBack(start)
	visited[start.x][start.y] = true

	steps := 0
	for queue.Len() > 0 {
		queueSize := queue.Len()
		for i := 0; i < queueSize; i++ {
			curr := queue.Remove(queue.Front()).(Point)
			if curr == target {
				return steps
			}

			for _, dir := range directions {
				newX, newY := curr.x+dir.x, curr.y+dir.y
				if check(newX, newY, m, n) && forest[newX][newY] != 0 && !visited[newX][newY] {
					queue.PushBack(Point{newX, newY})
					visited[newX][newY] = true
				}
			}
		}
		steps++
	}
	return -1
}

func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])
	var trees []Point
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				trees = append(trees, Point{i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		return forest[trees[i].x][trees[i].y] < forest[trees[j].x][trees[j].y]
	})
	start := Point{0, 0}
	totalSteps := 0

	for _, tree := range trees {
		steps := bfs(forest, start, tree)
		if steps == -1 {
			return -1
		}
		totalSteps += steps
		start = tree
	}
	return totalSteps
}
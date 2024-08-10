func dfs(x, y, n int, visited [][]bool) {
	if x < 0 || y < 0 || x >= 3 * n || y >= 3 * n || visited[x][y] {
		return
	}
	visited[x][y] = true
	dfs(x + 1, y, n, visited)
	dfs(x - 1, y, n, visited)
	dfs(x, y + 1, n, visited)
	dfs(x, y - 1, n, visited)
}

func regionsBySlashes(grid []string) int {
	n := len(grid)
	visited := make([][]bool, 3 * n)
	for i := range visited {
		visited[i] = make([]bool, 3 * n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '/' {
				visited[3 * i][3 * j + 2] = true
				visited[3 * i + 1][3 * j + 1] = true
				visited[3 * i + 2][3 * j] = true
			} else if grid[i][j] == '\\' {
				visited[3 * i][3 * j] = true
				visited[3 * i + 1][3 * j + 1] = true
				visited[3 * i + 2][3 * j + 2] = true
			}
		}
	}

	count := 0
	for i := 0; i < 3 * n; i++ {
		for j := 0; j < 3 * n; j++ {
			if !visited[i][j] {
				dfs(i, j, n, visited)
				count++
			}
		}
	}
	return count
}
func dfs(grid [][]int, x, y, count int) int {
    m, n := len(grid), len(grid[0])
    if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == -1 {
        return 0
    }
    if grid[x][y] == 2 {
        if count == -1 {
            return 1
        }
        return 0
    }
    grid[x][y] = -1
    res := 0
    res += dfs(grid, x+1, y, count-1)
    res += dfs(grid, x-1, y, count-1)
    res += dfs(grid, x, y+1, count-1)
    res += dfs(grid, x, y-1, count-1)
    grid[x][y] = 0

    return res
}

func uniquePathsIII(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    startX, startY := -1, -1
    cnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                startX, startY = i, j
            } else if grid[i][j] == 0 {
                cnt++
            }
        }
    }
    return dfs(grid, startX, startY, cnt)
}

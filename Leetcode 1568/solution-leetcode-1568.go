func dfs(grid [][]int, r, c int, visited [][]bool) {
    if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 || visited[r][c] {
        return
    }

    visited[r][c] = true

    dfs(grid, r+1, c, visited)
    dfs(grid, r-1, c, visited)
    dfs(grid, r, c+1, visited)
    dfs(grid, r, c-1, visited)
}

func check(grid [][]int) bool {
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 1 {
                dfs(grid, r, c, visited)
                for i := 0; i < rows; i++ {
                    for j := 0; j < cols; j++ {
                        if grid[i][j] == 1 && !visited[i][j] {
                            return false
                        }
                    }
                }
                return true
            }
        }
    }
    return false
}

func minDays(grid [][]int) int {
    if !check(grid) {
        return 0
    }
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[0]); c++ {
            if grid[r][c] == 1 {
                grid[r][c] = 0
                if !check(grid) {
                    return 1
                }
                grid[r][c] = 1
            }
        }
    }
    return 2
}
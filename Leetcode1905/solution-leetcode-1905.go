func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    m, n := len(grid1), len(grid1[0])
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid2[i][j] == 1 {
                if dfs(grid1, grid2, i, j, m, n) {
                    count++
                }
            }
        }
    }

    return count
}
func dfs(grid1, grid2 [][]int, x, y, m, n int) bool {
    if x < 0 || x >= m || y < 0 || y >= n || grid2[x][y] == 0 {
        return true
    }
    grid2[x][y] = 0
    isSubIsland := grid1[x][y] == 1
    isSubIsland = dfs(grid1, grid2, x+1, y, m, n) && isSubIsland
    isSubIsland = dfs(grid1, grid2, x-1, y, m, n) && isSubIsland
    isSubIsland = dfs(grid1, grid2, x, y+1, m, n) && isSubIsland
    isSubIsland = dfs(grid1, grid2, x, y-1, m, n) && isSubIsland

    return isSubIsland
}
func maxMoves(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    result := 0
    
    dp := make([]int, m) 

    for j := 1; j < n; j++ {
        prev := 0    
        check := false
        
        for i := 0; i < m; i++ {
            currm := -1
            currp := dp[i]
            if i-1 >= 0 && prev != -1 && grid[i][j] > grid[i-1][j-1] {
                currm = max(currm, prev+1)
            }
            if dp[i] != -1 && grid[i][j] > grid[i][j-1] {
                currm = max(currm, dp[i]+1)
            }
            if i+1 < m && dp[i+1] != -1 && grid[i][j] > grid[i+1][j-1] {
                currm = max(currm, dp[i+1]+1)
            }
            
            dp[i] = currm
            check = check || (dp[i] != -1)
            prev = currp
        }
        
        if !check {
            break
        }
        result = j
    }
    
    return result
}

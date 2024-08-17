func maxPoints(points [][]int) int64 {
    m, n := len(points), len(points[0])
    dp := make([]int64, n)
    for j := 0; j < n; j++ {
        dp[j] = int64(points[0][j])
    }
    for i := 1; i < m; i++ {
        lr := make([]int64, n)
        lr[0] = dp[0]
        for j := 1; j < n; j++ {
            lr[j] = max(lr[j-1] - 1, dp[j])
        }
        rl := make([]int64, n)
        rl[n-1] = dp[n-1]
        for j := n-2; j >= 0; j-- {
            rl[j] = max(rl[j+1] - 1, dp[j])
        }
        for j := 0; j < n; j++ {
            dp[j] = int64(points[i][j]) + max(lr[j], rl[j])
        }
    }
    var res int64 = dp[0]
    for j := 1; j < n; j++ {
        res = max(res, dp[j])
    }
    
    return res
}
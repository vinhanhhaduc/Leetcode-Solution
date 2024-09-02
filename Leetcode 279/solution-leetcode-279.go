func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = i
    }

    for i := 1; i <= n; i++ {
        j := 1
        for j*j <= i {
            dp[i] = min(dp[i], dp[i - j*j] + 1)
            j++
        }
    }

    return dp[n]
}
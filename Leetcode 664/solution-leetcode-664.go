func strangePrinter(s string) int {
    n := len(s)
    if n == 0 {
        return 0
    }

    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][i] = 1
    }

    for length := 2; length <= n; length++ {
        for i := 0; i <= n - length; i++ {
            j := i + length - 1
            dp[i][j] = length
            if s[i] == s[j] {
                dp[i][j] = dp[i][j-1]
            } else {
                for k := i; k < j; k++ {
                    dp[i][j] = min(dp[i][j], dp[i][k] + dp[k+1][j])
                }
            }
        }
    }

    return dp[0][n-1]
}
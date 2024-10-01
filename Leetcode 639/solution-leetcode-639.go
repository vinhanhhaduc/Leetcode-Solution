const mod = 1_000_000_007

func numDecodings(s string) int {
    n := len(s)
    dp := make([]int, n+1)
    dp[0] = 1
    
    for i := 1; i <= n; i++ {
        if s[i-1] == '*' {
            dp[i] += 9 * dp[i-1]
        } else if s[i-1] != '0' {
            dp[i] += dp[i-1]
        }
        if i > 1 {
            if s[i-2] == '1' {
                if s[i-1] == '*' {
                    dp[i] += 9 * dp[i-2]
                } else {
                    dp[i] += dp[i-2] 
                }
            } else if s[i-2] == '2' {
                if s[i-1] == '*' {
                    dp[i] += 6 * dp[i-2]
                } else if s[i-1] <= '6' {
                    dp[i] += dp[i-2]
                }
            } else if s[i-2] == '*' {
                if s[i-1] == '*' {
                    dp[i] += 15 * dp[i-2]
                } else if s[i-1] <= '6' {
                    dp[i] += 2 * dp[i-2]
                } else {
                    dp[i] += dp[i-2]
                }
            }
        }
        
        dp[i] %= mod
    }
    
    return dp[n]
}
func minHeightShelves(books [][]int, shelfWidth int) int {
    n := len(books)
    dp := make([]int, n + 1)
    for i := range dp {
        dp[i] = 1 << 31 - 1
    }
    dp[0] = 0

    for i := 1; i <= n; i++ {
        w := 0
        h := 0
        for j := i; j > 0; j-- {
            w += books[j - 1][0]
            if w > shelfWidth {
                break
            }
            if h < books[j - 1][1] {
                h = books[j - 1][1]
            }
            if dp[i] > dp[j - 1] + h {
                dp[i] = dp[j - 1] + h
            }
        }
    }

    return dp[n]
}
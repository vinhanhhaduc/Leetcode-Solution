func maxEnvelopes(envelopes [][]int) int {
    sort.Slice(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            return envelopes[i][1] > envelopes[j][1]
        }
        return envelopes[i][0] < envelopes[j][0]
    })
    dp := []int{}

    for _, envelope := range envelopes {
        height := envelope[1]
        i := sort.Search(len(dp), func(i int) bool { return dp[i] >= height })
        if i == len(dp) {
            dp = append(dp, height)
        } else {
            dp[i] = height
        }
    }

    return len(dp)
}
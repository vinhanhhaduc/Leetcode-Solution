func largestCombination(candidates []int) int {
    res := 0

    for bit := 0; bit < 30; bit++ {
        count := 0
        for _, num := range candidates {
            if (num & (1 << bit)) != 0 {
                count++
            }
        }
        if count > res {
            res = count
        }
    }

    return res
}
func minSwapsCouples(row []int) int {
    n := len(row)
    pos := make(map[int]int)
    for i := 0; i < n; i++ {
        pos[row[i]] = i
    }

    swaps := 0
    for i := 0; i < n; i += 2 {
        first := row[i]
        second := first ^ 1
        if row[i+1] != second {
            swaps++
            partnerIndex := pos[second]
            row[i+1], row[partnerIndex] = row[partnerIndex], row[i+1]
            pos[row[partnerIndex]] = partnerIndex
            pos[row[i+1]] = i + 1
        }
    }

    return swaps

}
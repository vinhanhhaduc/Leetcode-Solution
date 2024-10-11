func arrayRankTransform(arr []int) []int {
    n := len(arr)
    if n == 0 {
        return []int{}
    }
    sortedArr := make([]int, n)
    copy(sortedArr, arr)
    sort.Ints(sortedArr)
    rankMap := make(map[int]int)
    rank := 1
    for _, num := range sortedArr {
        if _, exists := rankMap[num]; !exists {
            rankMap[num] = rank
            rank++
        }
    }
    res := make([]int, n)
    for i := 0; i < n; i++ {
        res[i] = rankMap[arr[i]]
    }

    return res
}
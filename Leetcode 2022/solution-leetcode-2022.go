func construct2DArray(original []int, m int, n int) [][]int {
     if len(original) != m * n {
        return [][]int{}
    }
    res := make([][]int, m)
    for i := range res {
        res[i] = make([]int, n)
    }
    for i := 0; i < len(original); i++ {
        row := i / n
        col := i % n
        res[row][col] = original[i]
    }
    return res
}
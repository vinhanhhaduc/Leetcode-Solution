func backtrack(row, n int, cols, d1, d2 []bool, count *int) {
        if row == n {
            *count++
            return
        }
        for col := 0; col < n; col++ {
            idx1 := row - col + n   
            idx2 := row + col
            if !cols[col] && !d1[idx1] && !d2[idx2] {
                cols[col], d1[idx1], d2[idx2] = true, true, true
                backtrack(row + 1, n, cols, d1, d2, count)
                cols[col], d1[idx1], d2[idx2] = false, false, false
            }
        }
    }

func totalNQueens(n int) int {
    cols := make([]bool, n)
    d1 := make([]bool, 2*n)
    d2 := make([]bool, 2*n)
    count := 0
    backtrack(0, n, cols, d1, d2, &count)

    return count
}

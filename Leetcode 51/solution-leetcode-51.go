func solveNQueens(n int) [][]string {
    res := [][]string{}
    board := make([][]byte, n)
    for i := range board {
        board[i] = make([]byte, n)
        for j := range board[i] {
            board[i][j] = '.'
        }
    }
    
    columns := make(map[int]bool)
    main := make(map[int]bool)
    anti := make(map[int]bool)
    
    backtrack(0, n, board, &res, columns, main, anti)
    
    return res
}

func backtrack(row int, n int, board [][]byte, res *[][]string, columns, main, anti map[int]bool) {
    if row == n {
        solution := make([]string, n)
        for i := range board {
            solution[i] = string(board[i])
        }
        *res = append(*res, solution)
        return
    }

    for col := 0; col < n; col++ {
        if columns[col] || main[row - col] || anti[row + col] {
            continue
        }
        board[row][col] = 'Q'
        columns[col] = true
        main[row - col] = true
        anti[row + col] = true
        backtrack(row + 1, n, board, res, columns, main, anti)

        board[row][col] = '.'
        delete(columns, col)
        delete(main, row - col)
        delete(anti, row + col)
    }
}
func rotateTheBox(box [][]byte) [][]byte {
    rows := len(box)
    cols := len(box[0])
    
    res := make([][]byte, cols)
    for i := range res {
        res[i] = make([]byte, rows)
        for j := range res[i] {
            res[i][j] = '.'
        }
    }
    
    for r := 0; r < rows; r++ {
        i := cols - 1
        for c := cols - 1; c >= 0; c-- {
            if box[r][c] == '#' {
                res[i][rows - r - 1] = '#'
                i--
            } else if box[r][c] == '*' {
                res[c][rows - r - 1] = '*'
                i = c - 1
            }
        }
    }
    return res
}
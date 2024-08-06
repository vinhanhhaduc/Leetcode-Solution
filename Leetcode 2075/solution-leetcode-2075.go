func decodeCiphertext(encodedText string, rows int) string {
	n := len(encodedText)
    cols := n / rows
    matrix := make([][]rune, rows)
    index := 0
    for i := 0; i < rows; i++ {
        matrix[i] = []rune(encodedText[index:index+cols])
        index += cols
    }
    result := []rune{}
    for diag := 0; diag < cols; diag++ {
        r, c := 0, diag
        for r < rows && c < cols {
            if matrix[r][c] != ' ' {
                result = append(result, matrix[r][c])
            }
            r++
            c++
        }
    }

    return string(result)
}
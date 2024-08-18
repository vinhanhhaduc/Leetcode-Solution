func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }

    res := []int{}
    top, bottom := 0, len(matrix)-1
    left, right := 0, len(matrix[0])-1

    for top <= bottom && left <= right {
        for i := left; i <= right; i++ {
            res = append(res, matrix[top][i])
        }
        top++
        for i := top; i <= bottom; i++ {
            res = append(res, matrix[i][right])
        }
        right--
        if top <= bottom {
            for i := right; i >= left; i-- {
                res = append(res, matrix[bottom][i])
            }
            bottom--
        }
        if left <= right {
            for i := bottom; i >= top; i-- {
                res = append(res, matrix[i][left])
            }
            left++
        }
    }

    return res
}
func restoreMatrix(rowSum []int, colSum []int) [][]int {
    res := make([][]int, len(rowSum))
    for i := range res {
        res[i] = make([]int, len(colSum))
    }
    for i := 0;i < len(res); i++{
        for j := 0;j < len(res[0]); j++{
            res[i][j] = int(math.Min(float64(rowSum[i]), float64(colSum[j])));
            rowSum[i] -= res[i][j];
            colSum[j] -= res[i][j];
        }
    }
    return res;
}
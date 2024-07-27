func trimMean(arr []int) float64 {
    sort.Ints(arr)
    sum, f := 0, len(arr) / 20
    for i := f; i < len(arr) - f; i++ {
        sum += arr[i]
    }
    return float64(sum) / float64(len(arr) - f * 2)
}
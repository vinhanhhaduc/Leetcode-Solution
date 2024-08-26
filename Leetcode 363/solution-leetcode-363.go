func maxSumSubmatrix(matrix [][]int, k int) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }

    m, n := len(matrix), len(matrix[0])
    maxSum := math.MinInt32

    for r1 := 0; r1 < m; r1++ {
        sums := make([]int, n)

        for r2 := r1; r2 < m; r2++ {
            // Update sums to include the new row r2
            for col := 0; col < n; col++ {
                sums[col] += matrix[r2][col]
            }
            maxSum = max(maxSum, helper2(sums, k))
        }
    }

    return maxSum
}
func helper1(arr []int, target int) []int {
    i := sort.Search(len(arr), func(i int) bool { 
        return arr[i] >= target 
    })
    arr = append(arr[:i], append([]int{target}, arr[i:]...)...)
    return arr
}
func helper2(nums []int, k int) int {
    sum := 0
    maxSum := math.MinInt32
    prefixSums := []int{0}

    for _, num := range nums {
        sum += num
        i := sort.Search(len(prefixSums), func(j int) bool { 
            return sum - prefixSums[j] <= k 
        })

        if i < len(prefixSums) {
            maxSum = max(maxSum, sum-prefixSums[i])
        }

        prefixSums = helper1(prefixSums, sum)
    }

    return maxSum
}
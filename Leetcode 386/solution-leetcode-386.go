func dfs(curr int, n int ,result *[]int) {
        if curr > n {
            return
        }
        *result = append(*result, curr)
        for i := 0; i <= 9; i++ {
            if curr * 10 + i > n {
                break
            }
            dfs(curr * 10 + i, n, result)
        }
    }
func lexicalOrder(n int) []int {
    var result []int
    
    for i := 1; i <= 9; i++ {
        dfs(i, n, &result)
    }
    return result
}
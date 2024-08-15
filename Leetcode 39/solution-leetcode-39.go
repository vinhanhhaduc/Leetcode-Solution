func backtrack(candidates []int, target int, start int, combination []int, res *[][]int) {
    if target == 0 {
        *res = append(*res, append([]int{}, combination...))
        return
    }
    if target < 0 {
        return
    }

    for i := start; i < len(candidates); i++ {
        combination = append(combination, candidates[i])
        backtrack(candidates, target-candidates[i], i, combination, res)
        combination = combination[:len(combination)-1]
    }
}
func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    var combination []int
    
    sort.Ints(candidates)

    backtrack(candidates, target, 0, combination, &res)
    
    return res
}
func backtrack(candidates []int, start int, target int, combination []int, result *[][]int) {
    if target == 0 {
        *result = append(*result, append([]int{}, combination...))
        return
    }

    for i := start; i < len(candidates); i++ {
        if i > start && candidates[i] == candidates[i-1] {
            continue
        }

        if candidates[i] > target {
            break
        }

        backtrack(candidates, i+1, target-candidates[i], append(combination, candidates[i]), result)
    }
}

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    result := [][]int{}
    backtrack(candidates, 0, target, []int{}, &result)
    return result
}
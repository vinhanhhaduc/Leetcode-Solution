func judgePoint24(cards []int) bool {
    nums := make([]float64, len(cards))
    for i := range cards {
        nums[i] = float64(cards[i])
    }
    return solve(nums)
}

func solve(nums []float64) bool {
    if len(nums) == 1 {
        return math.Abs(nums[0] - 24) < 1e-6
    }
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if i != j {
                next := []float64{}
                for k := 0; k < len(nums); k++ {
                    if k != i && k != j {
                        next = append(next, nums[k])
                    }
                }
                for _, result := range helper(nums[i], nums[j]) {
                    next = append(next, result)
                    if solve(next) {
                        return true
                    }
                    next = next[:len(next)-1]
                }
            }
        }
    }
    return false
}
func helper(a, b float64) []float64 {
    results := []float64{a + b, a - b, a * b}
    if b != 0 {
        results = append(results, a / b)
    }
    return results
}

func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    if n == 2 {
        return int(math.Max(float64(nums[0]), float64(nums[1])))
    }
    return int(math.Max(float64(robRange(nums, 0, n-2)), float64(robRange(nums, 1, n-1))))
}
func robRange(nums []int, start int, end int) int {
    prev2, prev1 := 0, 0
    
    for i := start; i <= end; i++ {
        current := int(math.Max(float64(prev1), float64(nums[i]+prev2)))
        prev2 = prev1
        prev1 = current
    }
    
    return prev1
}
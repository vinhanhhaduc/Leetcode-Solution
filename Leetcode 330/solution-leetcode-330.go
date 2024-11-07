func minPatches(nums []int, n int) int {
    p := 0
    m := 1
    i := 0
    
    for m <= n {
        if i < len(nums) && nums[i] <= m {
            m += nums[i]
            i++
        } else {
            m += m
            p++
        }
    }
    
    return p
}
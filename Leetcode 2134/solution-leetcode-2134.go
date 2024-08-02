func minSwaps(nums []int) int {
    n := len(nums)
    t := 0
    for _, num := range nums {
        if num == 1 {
            t++
        }
    }
    
    if t == 0 || t == n {
        return 0
    }
    curr := 0
    for i := 0; i < t; i++ {
        if nums[i] == 1 {
            curr++
        }
    }
    
    maxo := curr
    for i := t; i < n + t; i++ {
        if nums[i % n] == 1 {
            curr++
        }
        if nums[(i - t) % n] == 1 {
            curr--
        }
        if curr > maxo {
            maxo = curr
        }
    }
    return t - maxo
}
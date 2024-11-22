func maximumSubarraySum(nums []int, k int) int64 {
    res := int64(0)
    count := make(map[int]int)
    curSum := int64(0)
    
    l := 0
    for r := 0; r < len(nums); r++ {
        curSum += int64(nums[r])
        count[nums[r]]++
        
        if r - l + 1 > k {
            count[nums[l]]--
            if count[nums[l]] == 0 {
                delete(count, nums[l])
            }
            curSum -= int64(nums[l])
            l++
        }
        
        if len(count) == r - l + 1 && r - l + 1 == k {
            if curSum > res {
                res = curSum
            }
        }
    }
    
    return res
}
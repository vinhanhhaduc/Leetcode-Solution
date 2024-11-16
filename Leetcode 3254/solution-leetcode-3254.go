func resultsArray(nums []int, k int) []int {
    res := make([]int, 0)
    l := 0
    consecCnt := 1
    
    for r := 0; r < len(nums); r++ {
        if r > 0 && nums[r-1]+1 == nums[r] {
            consecCnt++
        }
        
        if r-l+1 > k {
            if nums[l]+1 == nums[l+1] {
                consecCnt--
            }
            l++
        }
        
        if r-l+1 == k {
            if consecCnt == k {
                res = append(res, nums[r])
            } else {
                res = append(res, -1)
            }
        }
    }
    
    return res
}
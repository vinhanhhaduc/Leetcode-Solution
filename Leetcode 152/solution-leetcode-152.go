func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    maxNums := nums[0]
    currMax, currMin := nums[0], nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] < 0 {
            currMax, currMin = currMin, currMax
        }
        
        currMax = max(nums[i], currMax * nums[i])
        currMin = min(nums[i], currMin * nums[i])
        
        maxNums = max(maxNums, currMax)
    }
    
    return maxNums
}
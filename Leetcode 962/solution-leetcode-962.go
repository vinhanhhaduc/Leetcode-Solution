func maxWidthRamp(nums []int) int {
    n := len(nums)
    stack := []int{}
    for i := 0; i < n; i++ {
        if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
            stack = append(stack, i)
        }
    }
    
    maxWidth := 0
    for j := n - 1; j >= 0; j-- {
        for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
            i := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if j - i > maxWidth {
                maxWidth = j - i
            }
        }
    }
    
    return maxWidth
}
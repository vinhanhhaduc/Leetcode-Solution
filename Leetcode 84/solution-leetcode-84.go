func largestRectangleArea(heights []int) int {
     stack := []int{}
    res := 0
    n := len(heights)
    for i := 0; i < n; i++ {
        for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }
            res = max(res, h*width)
        }
        stack = append(stack, i)
    }
    for len(stack) > 0 {
        h := heights[stack[len(stack)-1]]
        stack = stack[:len(stack)-1]
        width := n
        if len(stack) > 0 {
            width = n - stack[len(stack)-1] - 1
        }
        res = max(res, h*width)
    }
    
    return res
}
func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0, len(nums)-k+1)
    dq := make([]int, 0, k+1)
    for i, v := range nums {
        if i > k-1 && dq[0] <= i-k {
            dq = dq[1:]
        }
        for len(dq) > 0 && v > nums[dq[len(dq)-1]] {
            dq = dq[:len(dq)-1]
        }
        dq = append(dq, i)
        if i >= k-1 {
            res = append(res, nums[dq[0]])
        }
    }
    return res
}
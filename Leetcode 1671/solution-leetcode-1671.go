func minimumMountainRemovals(nums []int) int {
    n := len(nums)
    if n < 3 {
        return 0
    }
    a := make([]int, n)
    for i := range a {
        a[i] = 1
    }
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                a[i] = max(a[i], a[j] + 1)
            }
        }
    }
    b := make([]int, n)
    for i := range b {
        b[i] = 1
    }
    for i := n - 2; i >= 0; i-- {
        for j := n - 1; j > i; j-- {
            if nums[i] > nums[j] {
                b[i] = max(b[i], b[j] + 1)
            }
        }
    }
    res  := 0
    for i := 1; i < n - 1; i++ {
        if a[i] > 1 && b[i] > 1 {
            res = max(res, a[i] + b[i] - 1)
        }
    }
    return n - res
}
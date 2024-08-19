func getPermutation(n int, k int) string {
    a := make([]int, n)
    nums := make([]int, n)
    a[0] = 1
    for i := 1; i < n; i++ {
        a[i] = a[i-1] * i
    }
    for i := 0; i < n; i++ {
        nums[i] = i + 1
    }
    k--
    res := ""
    for i := n; i > 0; i-- {
        idx := k / a[i-1]
        res += string('0' + nums[idx])
        nums = append(nums[:idx], nums[idx+1:]...)
        k %= a[i-1]
    }
    
    return res
}
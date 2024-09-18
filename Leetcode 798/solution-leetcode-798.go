func bestRotation(nums []int) int {
    n := len(nums)
    a := make([]int, n)

    for i := 0; i < n; i++ {
        left := (i + 1) % n
        right := (i - nums[i] + n + 1) % n
        a[left]++
        a[right]--

        if left > right {
            a[0]++
        }
    }

    maxScore := 0
    res := 0
    cur := 0
    for i := 0; i < n; i++ {
        cur += a[i]
        if cur > maxScore {
            maxScore = cur
            res = i
        }
    }

    return res
}
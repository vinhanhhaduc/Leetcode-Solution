func minimizedMaximum(n int, quantities []int) int {
    helper := func(x int) bool {
        s := 0
        for _, q := range quantities {
            s += (q + x - 1) / x
            if s > n {
                return false
            }
        }
        return s <= n
    }
    left, right := 1, 0
    for _, q := range quantities {
        if q > right {
            right = q
        }
    }
    for left < right {
        mid := (left + right) / 2
        if helper(mid) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return left
}
func countRangeSum(nums []int, lower int, upper int) int {
    n := len(nums)
    prefix := make([]int, n+1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    return helper(prefix, 0, n+1, lower, upper)
}

func helper(prefix []int, start, end, lower, upper int) int {
    if end - start <= 1 {
        return 0
    }
    
    mid := (start + end) / 2
    count := helper(prefix, start, mid, lower, upper) + 
             helper(prefix, mid, end, lower, upper)
    
    j, k, t := mid, mid, mid
    cache := make([]int, end-start)
    r := 0
    
    for i := start; i < mid; i++ {
        for k < end && prefix[k] - prefix[i] < lower {
            k++
        }
        for j < end && prefix[j] - prefix[i] <= upper {
            j++
        }
        count += j - k
        for t < end && prefix[t] < prefix[i] {
            cache[r] = prefix[t]
            t++
            r++
        }
        cache[r] = prefix[i]
        r++
    }
    
    copy(prefix[start:start + r], cache[:r])
    
    return count
}

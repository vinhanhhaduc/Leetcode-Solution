func longestSquareStreak(nums []int) int {
    sort.Ints(nums)
    a := make(map[int]bool)
    for _, num := range nums {
        a[num] = true
    }
    
    res := 0
    
    for _, num := range nums {
        length := 0
        current := num
        for a[current] {
            length++
            current *= current
        }

        if length >= 2 {
            res = max(res, length)
        }
    }
    
    if res < 2 {
        return -1
    }
    return res
}
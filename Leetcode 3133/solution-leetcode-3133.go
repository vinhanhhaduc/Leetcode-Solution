func minEnd(n int, x int) int64 {
    res := int64(x)
    r := int64(n - 1)
    pos := int64(1)
    
    for r != 0 {
        if (int64(x) & pos) == 0 {
            res |= (r & 1) * pos
            r >>= 1
        }
        pos <<= 1
    }
    
    return res
}
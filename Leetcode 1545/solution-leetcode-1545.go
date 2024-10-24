func findKthBit(n int, k int) byte {
    if n == 1 {
        return '0'
    }
    length := (1 << n) - 1  
    mid := length / 2 + 1   
    if k == mid {
        return '1'
    }
    if k < mid {
        return findKthBit(n - 1, k)
    }

    m := length - k + 1
    bit := findKthBit(n - 1, m)
    if bit == '0' {
        return '1'
    }
    return '0'
}
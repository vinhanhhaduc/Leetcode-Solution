func canArrange(arr []int, k int) bool {
    a := make([]int, k)
    for _, num := range arr {
        r := ((num % k) + k) % k 
        a[r]++
    }
    if a[0]%2 != 0 {
        return false
    }
    for i := 1; i <= k/2; i++ {
        if a[i] != a[k-i] {
            return false
        }
    }
    if k%2 == 0 && a[k/2]%2 != 0 {
        return false
    }

    return true
}
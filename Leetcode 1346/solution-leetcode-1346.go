func checkIfExist(arr []int) bool {
    n := len(arr)
    sort.Ints(arr) 
    for i := 0; i < n; i++ {
        l, r := 0, n-1
        x := 2 * arr[i]
        for l <= r {
            mid := (l + r) / 2
            if arr[mid] == x && mid != i { 
                return true
            } else if arr[mid] < x {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return false
}
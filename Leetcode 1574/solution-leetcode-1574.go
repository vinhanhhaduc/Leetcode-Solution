func findLengthOfShortestSubarray(arr []int) int {
    n := len(arr)
    if n <= 1 {
        return 0
    }
    left := 0
    for left < n-1 && arr[left] <= arr[left+1] {
        left++
    }
    if left == n-1 {
        return 0
    }
    right := n - 1
    for right > 0 && arr[right-1] <= arr[right] {
        right--
    }
    result := min(n-left-1, right)
    i, j := 0, right
    for i <= left && j < n {
        if arr[i] <= arr[j] {
            result = min(result, j-i-1)
            i++
        } else {
            j++
        }
    }

    return result
}
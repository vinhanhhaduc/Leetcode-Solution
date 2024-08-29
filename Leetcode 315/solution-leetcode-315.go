func mergeSort(nums []int, indices []int, counts []int, left int, right int) {
    if left >= right {
        return
    }

    mid := (left + right) / 2
    mergeSort(nums, indices, counts, left, mid)
    mergeSort(nums, indices, counts, mid+1, right)

    temp := make([]int, right-left+1)
    i, j, k := left, mid+1, 0
    rightCount := 0
    for i <= mid && j <= right {
        if nums[indices[j]] < nums[indices[i]] {
            temp[k] = indices[j]
            rightCount++
            j++
        } else {
            counts[indices[i]] += rightCount
            temp[k] = indices[i]
            i++
        }
        k++
    }
    for i <= mid {
        counts[indices[i]] += rightCount
        temp[k] = indices[i]
        i++
        k++
    }

    for j <= right {
        temp[k] = indices[j]
        j++
        k++
    }
    for i := left; i <= right; i++ {
        indices[i] = temp[i-left]
    }
}
func countSmaller(nums []int) []int {
    n := len(nums)
    counts := make([]int, n)
    indices := make([]int, n)
    for i := 0; i < n; i++ {
        indices[i] = i
    }

    mergeSort(nums, indices, counts, 0, n-1)
    return counts
}
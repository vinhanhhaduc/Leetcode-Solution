func searchRange(nums []int, target int) []int {
    n := len(nums)
    l, r := 0, n - 1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    li := l
    l, r = 0, n - 1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] <= target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    ri := l - 1

    if li >= n || nums[li] != target {
        return []int{-1, -1}
    } else {
        return []int{li, ri}
    }
}
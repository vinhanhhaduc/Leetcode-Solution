func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    if indexDiff <= 0 || valueDiff < 0 {
        return false
    }
    a := make(map[int]int)
    asize := valueDiff + 1

    for i := 0; i < len(nums); i++ {
        ai := helper(nums[i], asize)
        if _, exists := a[ai]; exists {
            return true
        }
        if num, exists := a[ai-1]; exists && abs(nums[i]-num) < asize {
            return true
        }
        if num, exists := a[ai+1]; exists && abs(nums[i]-num) < asize {
            return true
        }

        a[ai] = nums[i]
        if i >= indexDiff {
            delete(a, helper(nums[i-indexDiff], asize))
        }
    }

    return false
}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func helper(num int, size int) int {
    if num >= 0 {
        return num / size
    }
    return (num + 1) / size - 1
}
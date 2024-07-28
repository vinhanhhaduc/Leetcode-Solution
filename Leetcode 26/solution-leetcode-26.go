func removeDuplicates(nums []int) int {
    first := nums[0]
    l := 1
    for i := 0; i < len(nums); i++ {
        if nums[i] != first {
            nums[l] = nums[i]
            l++
        }
        first = nums[i]
    }
    return l
}
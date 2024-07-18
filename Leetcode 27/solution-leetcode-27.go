func removeElement(nums []int, val int) int {
    i := 0
    for _, j := range nums {
        if j != val {
            nums[i] = j
            i++
        }
    }
    return i
}
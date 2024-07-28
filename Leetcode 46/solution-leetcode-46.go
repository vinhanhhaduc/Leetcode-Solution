func backtrack(nums []int, first int, res *[][]int) {
    if first == len(nums) {
        perm := make([]int, len(nums))
        copy(perm, nums)
        *res = append(*res, perm)
        return
    }
    for i := first; i < len(nums); i++ {
        nums[first], nums[i] = nums[i], nums[first]
        backtrack(nums, first + 1, res)
        nums[first], nums[i] = nums[i], nums[first]
    }
}
func permute(nums []int) [][]int {
    var res [][]int
    backtrack(nums, 0, &res)
    return res
}
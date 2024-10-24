func countMaxOrSubsets(nums []int) int {
    maxOr := 0
    for _, num := range nums {
        maxOr |= num
    }

    var helper func(int, int) int
    helper = func(index, currentOr int) int {
        if index == len(nums) {
            if currentOr == maxOr {
                return 1
            }
            return 0
        }

        in := helper(index + 1, currentOr | nums[index])
        e := helper(index + 1, currentOr)
        return in + e
    }
    return helper(0, 0)
}
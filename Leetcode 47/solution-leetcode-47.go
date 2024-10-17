func permuteUnique(nums []int) [][]int {
    var result [][]int
    sort.Ints(nums)
    used := make([]bool, len(nums)) 
    var current []int

    var backtrack func()
    backtrack = func() {
        if len(current) == len(nums) {
            temp := make([]int, len(current))
            copy(temp, current)
            result = append(result, temp)
            return
        }

        for i := 0; i < len(nums); i++ {
            if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
                continue
            }
            current = append(current, nums[i])
            used[i] = true
            backtrack()
            current = current[:len(current)-1]
            used[i] = false
        }
    }
    backtrack()

    return result
}
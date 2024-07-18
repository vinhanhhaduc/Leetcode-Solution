/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) int {
    cnt := 0
    const MAX_DISTANCE = 10
    dfs(root, distance, &cnt, MAX_DISTANCE)
    return cnt
}
func dfs(node *TreeNode, distance int, count *int, MAX_DISTANCE int) []int {
    if node == nil {
        return make([]int, MAX_DISTANCE+1)
    }

    if node.Left == nil && node.Right == nil {
        res := make([]int, MAX_DISTANCE+1)
        res[1] = 1
        return res
    }

    left := dfs(node.Left, distance, count, MAX_DISTANCE)
    right := dfs(node.Right, distance, count, MAX_DISTANCE)

    for i := 1; i <= distance; i++ {
        for j := 1; j <= distance - i; j++ {
            *count += left[i] * right[j]
        }
    }

    res := make([]int, MAX_DISTANCE+1)
    for i := 1; i < MAX_DISTANCE; i++ {
        res[i + 1] = left[i] + right[i]
    }

    return res
}
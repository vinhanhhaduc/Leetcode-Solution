/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
     var result [][]int
    helper(root, targetSum, []int{}, &result)
    return result
}
func helper(node *TreeNode, targetSum int, currentPath []int, result *[][]int) {
    if node == nil {
        return
    }
    currentPath = append(currentPath, node.Val)
    if node.Left == nil && node.Right == nil && targetSum == node.Val {
        pathCopy := make([]int, len(currentPath))
        copy(pathCopy, currentPath)
        *result = append(*result, pathCopy)
    } else {
        helper(node.Left, targetSum - node.Val, currentPath, result)
        helper(node.Right, targetSum - node.Val, currentPath, result)
    }
    currentPath = currentPath[:len(currentPath)-1]
}
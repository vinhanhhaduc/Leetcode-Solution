/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    helper(root, &res)
    return res
}
func helper(node *TreeNode, res *[]int) {
    if node == nil {
        return
    }
    helper(node.Left, res)
    helper(node.Right, res)
    *res = append(*res, node.Val)
}
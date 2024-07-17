/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    toDeleteMap := make(map[int]bool)
    for _, val := range to_delete {
        toDeleteMap[val] = true
    }
    var forest []*TreeNode
    dfs(root, true, toDeleteMap, &forest)
    return forest
}
func dfs(node *TreeNode, isRoot bool, a map[int]bool, f *[]*TreeNode) *TreeNode {
    if node == nil {
        return nil
    }
    
    shouldDelete := a[node.Val]
    
    if isRoot && !shouldDelete {
        *f = append(*f, node)
    }
    
    node.Left = dfs(node.Left, shouldDelete, a, f)
    node.Right = dfs(node.Right, shouldDelete, a, f)
    
    if shouldDelete {
        return nil
    }
    return node
}
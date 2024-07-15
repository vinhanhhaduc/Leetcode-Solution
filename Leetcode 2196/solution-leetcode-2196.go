/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    nodes := make(map[int]*TreeNode)
    children := make(map[int]bool)

    for _, desc := range descriptions {
        pv, cv, isLeft := desc[0], desc[1], desc[2]

        pn, exists := nodes[pv]
        if !exists {
            pn = &TreeNode{Val: pv}
            nodes[pv] = pn
        }

        cn, exists := nodes[cv]
        if !exists {
            cn = &TreeNode{Val: cv}
            nodes[cv] = cn
        }

        if isLeft == 1 {
            pn.Left = cn
        } else {
            pn.Right = cn
        }
        children[cv] = true
    }

    var root *TreeNode
    for i, j := range nodes {
        if !children[i] {
            root = j
            break
        }
    }

    return root
}
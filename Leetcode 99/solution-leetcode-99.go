/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var first, second, prev *TreeNode

func recoverTree(root *TreeNode) {
	first, second, prev = nil, nil, nil
	inorder(root)
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	if prev != nil && root.Val < prev.Val {
		if first == nil {
			first = prev
		}
		second = root
	}
	prev = root
	inorder(root.Right)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderLen := len(postorder)
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[postorderLen-1]}
	postorder = postorder[:postorderLen-1]
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(inorder[:pos], postorder[:len(inorder[:pos])])
			root.Right = buildTree(inorder[pos+1:], postorder[len(inorder[:pos]):])
		}
	}
	return root
}
func buildTree1(inorder []int, postorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildInPos2TreeDFS(postorder, 0, len(postorder)-1, 0, inPos)
}

func buildInPos2TreeDFS(post []int, postStart int, postEnd int, inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	root := &TreeNode{Val: post[postEnd]}
	rootIdx := inPos[post[postEnd]]
	leftLen := rootIdx - inStart
	root.Left = buildInPos2TreeDFS(post, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = buildInPos2TreeDFS(post, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}
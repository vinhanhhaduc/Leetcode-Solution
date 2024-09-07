/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func check(treeNode *TreeNode, listNode *ListNode) bool {
    if listNode == nil { 
        return true
    }
    if treeNode == nil { 
        return false
    }
    if treeNode.Val == listNode.Val {
        return check(treeNode.Left, listNode.Next) || check(treeNode.Right, listNode.Next)
    }
    
    return false
}

func isSubPath(head *ListNode, root *TreeNode) bool {
    if root == nil {
        return false
    }
    return check(root, head) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

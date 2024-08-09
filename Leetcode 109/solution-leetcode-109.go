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
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	mid := helper(head)
	root := &TreeNode{Val: mid.Val}
	leftHead := head
	if leftHead != mid {
		root.Left = sortedListToBST(leftHead)
	}
	root.Right = sortedListToBST(mid.Next)

	return root
}
func helper(head *ListNode) *ListNode {
	prev := &ListNode{}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev.Next != nil {
		prev.Next = nil
	}
	return slow
}
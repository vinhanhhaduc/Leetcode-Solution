/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    curr := res
    carry := 0
    for l1 != nil || l2 != nil {
        var x, y int
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }
        
        sum := carry + x + y
        carry = sum / 10
        curr.Next = &ListNode{Val: sum % 10}
        curr = curr.Next
    }
    
    if carry > 0 {
        curr.Next = &ListNode{Val: carry}
    }
    
    return res.Next
}
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	for current != nil && current.Next != nil {
	    g := gcd(current.Val, current.Next.Val)
		newNode := &ListNode{Val: g}
		newNode.Next = current.Next
		current.Next = newNode
		current = newNode.Next
	}

	return head
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := &ListNode{Next: head, Val: -999999}
	cur := newHead
	last := newHead
	front := head
	for front.Next != nil {
		if front.Val == cur.Val {
			front = front.Next
			continue
		} else {
			if cur.Next != front {
				last.Next = front
				if front.Next != nil && front.Next.Val != front.Val {
					last = front
				}
				cur = front
				front = front.Next
			} else {
				last = cur
				cur = cur.Next
				front = front.Next
				
			}
		}
	}
	if front.Val == cur.Val {
		last.Next = nil
	} else {
		if cur.Next != front {
			last.Next = front
		}
	}
	return newHead.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next != nil && head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	nilNode := &ListNode{Val: 0, Next: head}
	head = nilNode

	lastVal := 0
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			lastVal = head.Next.Val
			for head.Next != nil && lastVal == head.Next.Val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return nilNode.Next
}

func deleteDuplicates4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nilNode := &ListNode{Val: 0, Next: head}
	lastIsDel := false
	head = nilNode
	pre, back := head.Next, head.Next.Next
	for head.Next != nil && head.Next.Next != nil {
		if pre.Val != back.Val && lastIsDel {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
			continue
		}

		if pre.Val == back.Val {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = true
		} else {
			head = head.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
		}
	}
	if lastIsDel && head.Next != nil {
		head.Next = nil
	}
	return nilNode.Next
}
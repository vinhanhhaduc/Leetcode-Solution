/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    current := head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}
	ps := length / k
	e := length % k
	result := make([]*ListNode, k)
	current = head
	for i := 0; i < k && current != nil; i++ {
		result[i] = current
		p := ps
		if i < e {
			p++
		}
		for j := 1; j < p; j++ {
			current = current.Next
		}
		next := current.Next
		current.Next = nil
		current = next
	}
	return result
}
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
     if root == nil {
        return nil
    }
    
    curr := root
    
    for curr != nil {
        if curr.Child != nil {
            next := curr.Next
            last := flatten(curr.Child)
            curr.Next = curr.Child
            curr.Child.Prev = curr
            for last.Next != nil {
                last = last.Next
            }
            last.Next = next
            if next != nil {
                next.Prev = last
            }

            curr.Child = nil
        }
        curr = curr.Next
    }
    
    return root
}
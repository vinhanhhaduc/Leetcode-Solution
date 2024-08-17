/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
        return nil
    }
    current := root
    for current.Left != nil {
        nextl := current.Left
        for current != nil {
            current.Left.Next = current.Right
            if current.Next != nil {
                current.Right.Next = current.Next.Left
            }
            current = current.Next
        }
        current = nextl
    }

    return root
}
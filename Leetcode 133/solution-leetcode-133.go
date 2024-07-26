/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    clone := make(map[*Node]*Node)
    return dfs(node, clone)
}
func dfs(node *Node, cloneg map[*Node]*Node) *Node {
    if n, found := cloneg[node]; found {
        return n
    }
    clone := &Node{Val: node.Val}
    cloneg[node] = clone

    for _, neighbor := range node.Neighbors {
        clone.Neighbors = append(clone.Neighbors, dfs(neighbor, cloneg))
    }
    
    return clone
}
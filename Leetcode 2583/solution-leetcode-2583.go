/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    if root == nil {
        return -1
    }
    queue := list.New()
    queue.PushBack(root)
    
    var res[]int
    for queue.Len() > 0 {
        size := queue.Len()
        total := 0
        for i := 0; i < size; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            total += node.Val
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }
        res = append(res, total)
    }
    sort.Slice(res, func(i, j int) bool {
        return res[i] > res[j]
    })
    if len(res) < k {
        return -1
    }
    return int64(res[k-1])
}
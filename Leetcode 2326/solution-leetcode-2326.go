/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func spiralMatrix(m int, n int, head *ListNode) [][]int {
    matrix := make([][]int, m)
    for i := 0; i < m; i++ {
        matrix[i] = make([]int, n)
        for j := 0; j < n; j++ {
            matrix[i][j] = -1
        }
    }
    top, bottom, left, right := 0, m-1, 0, n-1
    curr := head
    for top <= bottom && left <= right && curr != nil {
        for i := left; i <= right && curr != nil; i++ {
            matrix[top][i] = curr.Val
            curr = curr.Next
        }
        top++
        for i := top; i <= bottom && curr != nil; i++ {
            matrix[i][right] = curr.Val
            curr = curr.Next
        }
        right--
        for i := right; i >= left && curr != nil; i-- {
            matrix[bottom][i] = curr.Val
            curr = curr.Next
        }
        bottom--

        for i := bottom; i >= top && curr != nil; i-- {
            matrix[i][left] = curr.Val
            curr = curr.Next
        }
        left++
    }

    return matrix
}

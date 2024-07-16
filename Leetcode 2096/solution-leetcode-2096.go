/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func find(root *TreeNode, value int, res []byte) ([]byte, bool) {
	if root == nil {
		return nil, false
	}

	if root.Val == value {
		return res, true
	}

	if leftPath, found := find(root.Left, value, append(res, 'L')); found {
		return leftPath, true
	}

	if rightPath, found := find(root.Right, value, append(res, 'R')); found {
		return rightPath, true
	}

	return nil, false
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	s, _ := find(root, startValue, []byte{})
	d, _ := find(root, destValue, []byte{})
	
	i := 0
	for i < len(s) && i < len(d) && s[i] == d[i] {
		i++
	}
	s = s[i:]
	d = d[i:]
	u := strings.Repeat("U", len(s))
	t := string(d)

	return u + t
}
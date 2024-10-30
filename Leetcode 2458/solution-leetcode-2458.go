func treeQueries(root *TreeNode, queries []int) []int {
    heights := map[int]int{}
    helper2(root, heights)
    maxRemoved := map[int]int{}
    helper(root, 0, 0, maxRemoved, heights)
    res := []int{}
    for _, q := range queries {
        res = append(res, maxRemoved[q])
    }
    return res
}

func helper(root *TreeNode, lvl, mx int, r, h map[int]int) {
    if root == nil {
        return
    }
    r[root.Val] = mx
    mxh := max(mx, lvl)
    if root.Right != nil {
        mxh = max(mx, h[root.Right.Val] + lvl)
    }
    helper(root.Left, lvl+1, mxh, r, h)
    mxh = max(mx, lvl)
    if root.Left != nil {
        mxh = max(mx, h[root.Left.Val] + lvl)
    }
    helper(root.Right, lvl+1, mxh, r, h)
}

func helper2(root *TreeNode, h map[int]int) int {
    if root == nil {
        return 0
    }
    l := helper2(root.Left, h)
    r := helper2(root.Right, h)
    h[root.Val] = max(l, r) + 1
    return h[root.Val]
}
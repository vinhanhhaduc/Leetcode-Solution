type Codec struct{}
func Constructor() Codec {
	return Codec{}
}
func serializeHelper(node *TreeNode, sb *[]string) {
	if node == nil {
		*sb = append(*sb, "#")
		return
	}
	*sb = append(*sb, strconv.Itoa(node.Val))
	serializeHelper(node.Left, sb)
	serializeHelper(node.Right, sb)
}
func (this *Codec) serialize(root *TreeNode) string {
	var sb []string
	serializeHelper(root, &sb)
	return strings.Join(sb, ",")
}
func deserializeHelper(nodes *[]string, index *int) *TreeNode {
	if (*nodes)[*index] == "#" {
		*index++
		return nil
	}
	val, _ := strconv.Atoi((*nodes)[*index])
	node := &TreeNode{Val: val}
	*index++
	node.Left = deserializeHelper(nodes, index)
	node.Right = deserializeHelper(nodes, index)
	return node
}
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	index := 0
	return deserializeHelper(&nodes, &index)
}
package serializeandreconstructtree

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

func Encode(root *NTreeNode) *Node {
	if root == nil {
		return nil
	}
	head := &Node{Val: root.Val}
	head.Left = en(root.Children)
	return head
}
func en(childrens []*NTreeNode) *Node {
	var head, cur *Node
	for _, child := range childrens {
		node := &Node{Val: child.Val}
		if head == nil {
			head = node
		} else {
			cur.Right = node
		}
		cur = node
		cur.Left = en(child.Children)
	}
	return head
}
func Decode(root *Node) *NTreeNode {
	if root == nil {
		return nil
	}
	return &NTreeNode{Val: root.Val, Children: de(root.Left)}
}
func de(root *Node) (ans []*NTreeNode) {
	for root != nil {
		cur := &NTreeNode{Val: root.Val, Children: de(root.Left)}
		ans = append(ans, cur)
		root = root.Right
	}
	return
}

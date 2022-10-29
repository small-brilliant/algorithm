package main

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

func NewNode(v int) *Node {
	return &Node{Val: v}
}
func flatten(root *Node) {
	dfs(root)
}
func dfs(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.Left == nil && head.Right == nil {
		return head
	}
	if head.Left != nil {
		node := head.Right

		head.Right = head.Left

		lastNode := dfs(head.Left)

		head.Left = nil
		if lastNode != nil {
			lastNode.Right = node
		}
		return dfs(node)
	}
	return dfs(head.Right)
}
func main() {
	head := NewNode(1)
	head.Left = NewNode(2)
	head.Left.Left = NewNode(3)
	head.Left.Right = NewNode(4)
	head.Left.Left.Left = NewNode(5)
	flatten(head)
}

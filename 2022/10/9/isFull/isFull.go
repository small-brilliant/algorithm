package isfull

// 判断一棵树是不是满二叉树
type Node struct {
	Val   int
	Right *Node
	Left  *Node
}
type Info struct {
	nodes  int
	height int
}

func NewInfo(nodes, height int) *Info {
	return &Info{nodes, height}
}
func IsFull(head *Node) bool {
	if head == nil {
		return true
	}
	info := process(head)
	return 1<<info.height-1 == info.nodes
}
func process(head *Node) *Info {
	if head == nil {
		return NewInfo(0, 0)
	}
	left, right := process(head.Left), process(head.Right)
	nodes := left.nodes + right.nodes + 1
	height := max(left.height, right.height) + 1
	return NewInfo(nodes, height)
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

package maxdistance

// 找到一棵树里面的两个节点之间的最大距离
type Node struct {
	Val   int
	Right *Node
	Left  *Node
}
type Info struct {
	maxdistance int
	height      int
}

func NewInfo(maxdistance, height int) *Info {
	return &Info{maxdistance, height}
}
func MaxDistance2(head *Node) int {
	return process(head).maxdistance
}
func process(head *Node) *Info {
	if head == nil {
		return NewInfo(0, 0)
	}
	left, right := process(head.Left), process(head.Right)
	maxdistance := 0
	p1 := left.maxdistance
	p2 := right.maxdistance
	p3 := left.height + right.height + 1
	height := max(left.height, right.height) + 1
	maxdistance = max(max(p1, p2), p3)
	return NewInfo(maxdistance, height)
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

package isbalanced

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}
type Info struct {
	Isbalan bool
	height  int
}

func NewInfo(Isbalan bool, height int) *Info {
	return &Info{
		Isbalan,
		height,
	}
}
func IsBalanced(head *Node) bool {
	return process(head).Isbalan
}

func process(head *Node) *Info {
	if head == nil {
		return NewInfo(false, 0)
	}
	Isbalan, height := false, 0
	left, right := process(head.Left), process(head.Right)
	if !left.Isbalan || !right.Isbalan || abs(left.height-right.height) > 1 {
		Isbalan = false
	}
	height = max(left.height, right.height) + 1
	return NewInfo(Isbalan, height)
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

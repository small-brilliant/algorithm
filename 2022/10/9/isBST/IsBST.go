package isbst

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}
type Info struct {
	isBST bool
	maxV  int
	minV  int
}

func NewInfo(isBST bool, maxV int, minV int) *Info {
	return &Info{isBST, maxV, minV}
}

func IsBST(head *Node) bool {
	if head == nil {
		return true
	}
	return process(head).isBST
}
func process(head *Node) *Info {
	if head == nil {
		return nil
	}
	isBST := true
	left, right := process(head.Left), process(head.Right)
	var maxV, minV int
	if left != nil {
		maxV = left.maxV
		minV = left.minV
	}
	if right != nil {
		maxV = max(right.maxV, maxV)
		minV = min(right.minV, minV)
	}
	if right != nil && (head.Val >= right.minV || !right.isBST) {
		isBST = false
	}

	if left != nil && (head.Val <= left.maxV || !left.isBST) {
		isBST = false
	}
	return NewInfo(isBST, maxV, minV)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

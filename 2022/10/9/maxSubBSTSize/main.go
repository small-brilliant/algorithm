package main

import "fmt"

// 找到一棵树中节点数最多的子树，这个字数满足搜索二叉树
type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

func NewNode(v int) *Node {
	return &Node{Val: v}
}

type Info struct {
	nodes      int
	isbalanced bool
	max        int
	min        int
}

func NewInfo(nodes, max, min int, isbalanced bool) *Info {
	return &Info{nodes, isbalanced, max, min}
}

func MaxSubBSTSize(head *Node) int {
	return process(head).nodes
}
func process(head *Node) *Info {
	if head == nil {
		return nil
	}
	nodes := 0
	isbalanced := true
	l, r := process(head.Left), process(head.Right)
	maxV := head.Val
	minV := head.Val
	if l != nil {
		maxV = l.max
		minV = l.min
	}
	if r != nil {
		maxV = max(maxV, r.max)
		minV = min(minV, r.min)
	}
	if l != nil && (!l.isbalanced || l.max >= head.Val) {
		isbalanced = false
	}
	if r != nil && (!r.isbalanced || r.min <= head.Val) {
		isbalanced = false
	}
	if isbalanced {
		if r != nil {
			nodes += r.nodes
		}
		if l != nil {
			nodes += l.nodes
		}
		nodes += 1
	} else {
		nodes = max(l.nodes, r.nodes)
	}
	return NewInfo(nodes, maxV, minV, isbalanced)
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

func main() {
	head := NewNode(5)
	// head.Left = NewNode(4)
	// head.Left.Left = NewNode(1)
	// head.Left.Right = NewNode(6)
	// head.Right = NewNode(8)
	// head.Right.Left = NewNode(7)
	// head.Right.Right = NewNode(9)
	// head.Right.Right.Right = NewNode(10)
	fmt.Println(MaxSubBSTSize(head))
}

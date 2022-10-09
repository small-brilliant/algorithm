package iscbt

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

func IsCBT(head *Node) bool {
	var flag bool
	queue := make([]*Node, 0)
	if head == nil {
		return true
	}
	queue = append(queue, head)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil && node.Right != nil {
			return false
		}
		if flag && node.Left != nil || node.Right != nil {
			return false
		}
		if node.Left == nil || node.Right == nil {
			flag = true
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return true
}

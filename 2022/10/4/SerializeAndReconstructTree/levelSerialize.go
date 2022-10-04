package serializeandreconstructtree

import (
	"strconv"
)

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

func LeveSerial(head *Node) (ans []string) {
	if head == nil {
		ans = append(ans, "#")
		return
	}
	ans = append(ans, strconv.Itoa(head.Val))
	queue := make([]*Node, 0)
	queue = append(queue, head)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
			ans = append(ans, strconv.Itoa(node.Left.Val))
		} else {
			ans = append(ans, "#")
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			ans = append(ans, strconv.Itoa(node.Right.Val))
		} else {
			ans = append(ans, "#")
		}
	}
	return
}
func BuildByLevelQueue(str []string) *Node {
	if len(str) == 0 {
		return nil
	}
	head := generateNode(str[0])
	str = str[1:]
	queue := make([]*Node, 0)
	queue = append(queue, head)
	var node *Node
	for len(queue) != 0 {
		node = queue[0]
		queue = queue[1:]
		node.Left = generateNode(str[0])
		str = str[1:]
		node.Right = generateNode(str[0])
		str = str[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return head
}

func generateNode(v string) *Node {
	if v == "#" {
		return nil
	}
	i, _ := strconv.Atoi(v)
	return &Node{Val: i}
}

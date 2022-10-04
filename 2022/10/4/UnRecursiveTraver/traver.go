package main

import "fmt"

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

// 先序遍历的非递归
func pre(head *Node) {
	fmt.Println("pre-order: ")
	if head != nil {
		stack := make([]Node, 0)
		stack = append(stack, *head)
		for len(stack) != 0 {
			*head = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(head.Val)
			if head.Right != nil {
				stack = append(stack, *head.Right)
			}
			if head.Left != nil {
				stack = append(stack, *head.Left)
			}
		}
	}
}

// 后续遍历的非递归
func pos1(head *Node) {
	fmt.Println("pos-order: ")
	if head != nil {
		stack1 := make([]Node, 0)
		stack2 := make([]Node, 0)
		stack1 = append(stack1, *head)
		for len(stack1) != 0 {
			*head = stack1[len(stack1)-1]
			stack2 = append(stack2, *head)
			if head.Left != nil {
				stack1 = append(stack1, *head.Left)
			}
			if head.Right != nil {
				stack1 = append(stack1, *head.Right)
			}
		}
		for len(stack2) != 0 {
			node := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			fmt.Println(node.Val)
		}

	}
}

//  中序遍历的非递归
func in(cur *Node) {
	if cur != nil {
		stack := make([]*Node, 0)
		for len(stack) != 0 {
			if cur != nil {
				stack = append(stack, cur)
				cur = cur.Left
			} else {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Println(cur.Val)
				cur = cur.Right
			}
		}

	}
}

package main

import (
	"algorithm/utils"
	"fmt"
)

type Node struct {
	pass  int
	end   int
	nexts [26]*Node
}

func Construct() Node {
	return Node{}
}

func (n *Node) Insert(word string) {
	if word == "" {
		return
	}
	node := n
	node.pass++
	path := 0
	for i := 0; i < len(word); i++ {
		path = int(word[i] - 'a')
		if node.nexts[path] == nil {
			node.nexts[path] = new(Node)
		}
		node = node.nexts[path]
		node.pass++
	}
	node.end++
}
func (n *Node) Search(word string) *Node {
	if word == "" {
		return nil
	}
	node := n
	path := 0
	for i := 0; i < len(word); i++ {
		path = int(word[i] - 'a')
		if node.nexts[path] != nil {
			node = node.nexts[path]
		} else {
			return nil
		}
	}
	return node
}
func (n *Node) Delete(word string) bool {
	if word == "" {
		return false
	}
	node := n
	path := 0
	for i := 0; i < len(word); i++ {
		path = int(word[i] - 'a')
		if node.nexts[path] != nil {
			node.pass--
			node = node.nexts[path]
		} else {
			return false
		}
	}
	node.end--
	return true
}

func (n *Node) prefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	node := n
	path := 0
	for i := 0; i < len(pre); i++ {
		path = int(pre[i] - 'a')
		if node.nexts[path] != nil {
			node = node.nexts[path]
		}
	}
	return node.pass
}
func main() {
	maxLen := 100
	node := Construct()
	for i := 0; i < 1000; i++ {
		node.Insert(utils.RandomArrString(maxLen))
	}
	node.Delete("agc")
	fmt.Println(node.Search("abcde"))
}

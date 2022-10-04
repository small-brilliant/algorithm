package serializeandreconstructtree

import "strconv"

func PreSerial(head *Node) (ans []string) {
	pres(head, ans)
	return
}
func pres(head *Node, ans []string) {
	if head == nil {
		_ = append(ans, "#")
		return
	}
	ans = append(ans, strconv.Itoa(head.Val))
	pres(head.Left, ans)
	pres(head.Right, ans)
}

func BuildByPreQueue(str []string) *Node {
	if len(str) == 0 {
		return nil
	}
	return preb(str)
}
func preb(str []string) *Node {
	v := str[0]
	str = str[1:]
	if v == "#" {
		return nil
	}
	num, _ := strconv.Atoi(v)
	head := &Node{Val: num}
	head.Left = preb(str)
	head.Right = preb(str)
	return head
}

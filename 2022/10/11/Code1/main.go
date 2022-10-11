package main

import "fmt"

func code1(a [][]int) int {
	N := len(a)
	if N&1 == 1 {
		return 0
	}
	n := N >> 1
	return process(a, 0, n)
}

// rest 表示A地区的剩余
func process(income [][]int, index, rest int) int {
	if index == len(income) {
		return 0
	}
	if len(income)-index == rest {
		return income[index][0] + process(income, index+1, rest-1)
	}
	if rest == 0 {
		return income[index][1] + process(income, index+1, rest)
	}
	p1 := income[index][1] + process(income, index+1, rest)
	p2 := income[index][0] + process(income, index+1, rest-1)
	return max(p1, p2)
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	a := [][]int{{2, 4}, {9, 2}, {1, 2}, {3, 1}}
	fmt.Println(code1(a))
}

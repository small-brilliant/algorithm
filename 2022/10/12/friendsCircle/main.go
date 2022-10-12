package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Frends(arr [][]int) int {
	parent, rank := make([]int, 1000001), make([]int, 1000001)
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union func(i, j int)
	union = func(i, j int) {
		x, y := find(i), find(j)
		if x == y {
			return
		}
		if rank[x] > rank[y] {
			x, y = y, x
		}
		parent[x] = y
		rank[y] += rank[x]
	}
	for i := 0; i < len(parent); i++ {
		parent[i] = i
		rank[i] = 1
	}
	for _, a := range arr {
		union(a[0], a[1])
	}
	ans := 0
	for i := 0; i < len(rank); i++ {
		ans = max(ans, rank[i])
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	T, _ := strconv.Atoi(in.Text())
	for i := 0; i < T; i++ {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())
		arr := make([][]int, n)
		for j := 0; j < n; j++ {
			in.Scan()
			data := strings.Split(in.Text(), " ")
			x, _ := strconv.Atoi(data[0])
			y, _ := strconv.Atoi(data[0])
			arr[j] = []int{x, y}
		}
		fmt.Println(Frends(arr))
	}
}

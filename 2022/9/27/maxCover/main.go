package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type t []int

func (t *t) Push(v interface{}) {
	*t = append(*t, v.(int))
}
func (t *t) Pop() (v interface{}) {
	v = (*t)[:t.Len()-1]
	*t = (*t)[:t.Len()-1]
	return
}
func (t *t) Less(i, j int) bool {
	return (*t)[i] < (*t)[j]
}
func (t *t) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}
func (t *t) Len() int {
	return len(*t)
}

func maxCover(a [][]int) int {
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	h := new(t)
	heap.Init(h)
	ans := 1
	for _, v := range a {
		for h.Len() > 0 && (*h)[0] <= v[0] {
			h.Pop()
		}
		h.Push(v[1])
		if ans <= h.Len() {
			ans = h.Len()
		}
	}
	return ans
}
func main() {
	ans := maxCover([][]int{{1, 2}, {2, 3}, {3, 11}, {11, 122}, {122, 133}, {133, 143}})
	fmt.Println(ans)
}

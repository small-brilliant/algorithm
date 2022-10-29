package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func merge(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	h := &topK{}
	heap.Init(h)
	heap.Push(h, intervals[0][1])
	ans := 0
	for _, interv := range intervals[1:] {
		if interv[0] > (*h)[0] {
			ans = max(ans, h.Len())
			for h.Len() > 0 && interv[0] > (*h)[0] {
				heap.Pop(h)
			}
		}
		heap.Push(h, interv[1])
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type topK []int

func (t *topK) Less(i, j int) bool { return (*t)[i] < (*t)[j] }
func (t *topK) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}
func (t *topK) Len() int {
	return len(*t)
}
func (t *topK) Pop() interface{} {
	a := (*t)[t.Len()-1]
	*t = (*t)[:t.Len()-1]
	return a
}
func (t *topK) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func main() {

	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {2, 7}, {8, 10}, {15, 18}}))
}

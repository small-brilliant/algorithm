package main

import (
	"algorithm/utils"
	"container/heap"
	"fmt"
)

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
func topkOfHeap(a []int, k int) []int {
	h := &topK{}
	for _, v := range a {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return *h
}

func main() {
	maxValue, maxLen := 20, 10
	fmt.Println("===测试开始===")
	for i := 0; i < 1; i++ {
		nums := utils.RandomArray(maxValue, maxLen)
		ans := topkOfHeap(nums, 3)
		for _, num := range nums {
			fmt.Print(num)
			fmt.Print(",")
		}
		fmt.Println()
		fmt.Println(ans)
	}
	fmt.Println("===测试结束===")
}

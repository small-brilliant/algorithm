package main

import (
	"algorithm/utils"
	"fmt"
)

// 小顶堆
func heapify(a []int, root, end int) {
	for {
		// 从0开始，左孩子节点索引
		child := root*2 + 1
		if child > end {
			return
		}
		// 比较左孩子和右孩子的大小
		if child < end && a[child] >= a[child+1] {
			child++
		}
		if a[root] < a[child] {
			return
		}
		a[root], a[child] = a[child], a[root]
		// 继续往下沉
		root = child
	}
}

func topkOfHeap2(a []int, k int) []int {
	end := len(a) - 1
	for i := end / 2; i >= 0; i-- {
		heapify(a, i, end)
	}
	for i := end; i >= 0; i-- {
		a[i], a[0] = a[0], a[i]
		end--
		heapify(a, 0, end)
	}
	return a[:k]
}

func main() {
	maxValue, maxLen := 20, 10
	fmt.Println("===测试开始===")
	for i := 0; i < 1; i++ {
		nums := utils.RandomArray(maxValue, maxLen)
		ans := topkOfHeap2(nums, 3)
		for _, num := range nums {
			fmt.Print(num)
			fmt.Print(",")
		}
		fmt.Println()
		fmt.Println(ans)
	}
	fmt.Println("===测试结束===")
}

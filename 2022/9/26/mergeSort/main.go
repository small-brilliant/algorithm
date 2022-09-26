package main

import (
	"algorithm/utils"
	"fmt"
)

func mergeSort(a []int) []int {
	process(a, 0, len(a)-1)
	return a
}
func process(a []int, L, R int) {
	if L >= R {
		return
	}
	mid := L + (R-L)/2
	process(a, L, mid)
	process(a, mid+1, R)
	merge(a, L, mid, R)
}
func merge(a []int, L, mid, R int) {
	ans := make([]int, R-L+1)
	p1, p2 := L, mid+1
	index := 0
	for ; p1 <= mid && p2 <= R; index++ {
		if a[p1] <= a[p2] {
			ans[index] = a[p1]
			p1++
		} else {
			ans[index] = a[p2]
			p2++
		}
	}
	for p1 <= mid {
		ans[index] = a[p1]
		p1++
		index++
	}
	for p2 <= R {
		ans[index] = a[p2]
		p2++
		index++
	}
	for i := L; i <= R; i++ {
		a[i] = ans[i-L]
	}
}

func check(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}
func main() {
	maxValue, maxLen := 200, 2000
	fmt.Println("===测试开始===")
	for i := 0; i < 10000; i++ {
		nums := utils.RandomArray(maxValue, maxLen)
		ans := mergeSort(nums)
		if !check(ans) {
			for _, num := range nums {
				fmt.Print(num)
				fmt.Print(",")
			}
			fmt.Println()
			fmt.Println(ans)
			break
		}
	}
	fmt.Println("===测试结束===")
}

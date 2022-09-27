package main

import (
	"algorithm/utils"
	"fmt"
)

func quickSort(a []int) []int {
	fz(a, 0, len(a)-1)
	return a
}
func fz(a []int, L, R int) {
	if L >= R {
		return
	}
	l, r := L, R
	p := a[L]
	for l < r {
		for l < r && p <= a[r] {
			r--
		}
		a[l] = a[r]
		for l < r && p >= a[l] {
			l++
		}
		a[r] = a[l]
	}
	a[l] = p
	fz(a, L, l-1)
	fz(a, l+1, R)
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
	maxValue, maxLen := 200, 200
	fmt.Println("===测试开始===")
	for i := 0; i < 10000; i++ {
		nums := utils.RandomArray(maxValue, maxLen)
		ans := quickSort(nums)
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

package main

import (
	"algorithm/utils"
	"fmt"
	"math/rand"
)

func helanFlag(a []int, x int) []int {
	l, r := 0, len(a)-1
	i := 0
	for i <= r && l < r {
		if a[i] > x {
			a[i], a[r] = a[r], a[i]
			r--
			continue
		} else if a[i] < x {
			a[i], a[l] = a[l], a[i]
			l++
			i++
		} else {
			i++
		}
	}
	return a
}
func check(a []int, x int) bool {
	index, n := 0, len(a)
	for index < n && a[index] < x {
		index++
	}
	for index < n && a[index] == x {
		index++
	}
	for index < n && a[index] > x {
		index++
	}
	return index == n
}
func main() {
	maxValue, maxLen := 200, 100
	fmt.Println("===测试开始===")
	for i := 0; i < 10000; i++ {
		nums := utils.RandomArray(maxLen, maxValue)
		x := int(rand.Float64() * float64(maxValue))
		ans := helanFlag(nums, x)
		if !check(ans, x) {
			for _, num := range nums {
				fmt.Print(num)
				fmt.Print(",")
			}
			fmt.Println()
			fmt.Println(x)
			fmt.Println(ans)
			break
		}
	}
	fmt.Println("===测试结束===")
}

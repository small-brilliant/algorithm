package main

import (
	"algorithm/utils"
	"fmt"
	"math"
)

func radixsort(a []int) []int {
	digit := 0
	max := int(^0)
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	for max != 0 {
		digit++
		max /= 10
	}
	ans := make([]int, len(a))
	for d := 1; d <= digit; d++ {
		count := make([]int, 10)
		for _, v := range a {
			idx := getDigit(v, d)
			count[idx]++
		}
		// 表示当前位数的最后位置
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}
		for j := len(a) - 1; j >= 0; j-- {
			idx := getDigit(a[j], d)
			ans[count[idx]-1] = a[j]
			count[idx]--
		}
		for i := 0; i < len(a); i++ {
			a[i] = ans[i]
		}
	}
	return a
}
func getDigit(num, digit int) int {
	if digit == 1 {
		return num % 10
	}
	num /= int(math.Pow10(digit - 1))
	return num % 10
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
	maxValue, maxLen := 20000, 10000
	fmt.Println("===测试开始===")
	for i := 0; i < 20000; i++ {
		nums := utils.RandomArray(maxLen, maxValue)
		ans := radixsort(nums)
		if !check(ans) {
			for _, num := range nums {
				fmt.Print(num)
				fmt.Print(",")
			}
			fmt.Println()
			fmt.Println(ans)
		}
	}
	fmt.Println("===测试结束===")
}

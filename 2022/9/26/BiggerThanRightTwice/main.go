package main

import (
	"algorithm/utils"
	"fmt"
)
// i位置左边比a[i]*2还大的个数。
func biggerThanRightTwice(a []int) int {
	b := make([]int, len(a))
	copy(a, b)
	count := 0
	var process func(b []int, L, R int)
	process = func(b []int, L, R int) {
		if L >= R {
			return
		}
		M := L + (R-L)/2
		process(b, L, M)
		process(b, M+1, R)
		count += merge(b, L, M, R)
	}
	process(b, 0, len(b)-1)
	return count

}

func merge(a []int, L, M, R int) int {
	ans := make([]int, R-L+1)
	p1, p2 := L, M+1
	i, count := 0, 0
	for ; p1 <= M && p2 <= R; i++ {
		if a[p1] <= a[p2] {
			ans[i] = a[p1]
			p1++
		} else {
			if a[p2]*2 < a[p1] {
				count += (M - p1) + 1
			}
			ans[i] = a[p2]
			p2++
		}
	}
	for p1 <= M {
		ans[i] = a[p1]
		p1++
		i++
	}
	for p2 <= R {
		ans[i] = a[p2]
		p2++
		i++
	}
	for i := L; i <= R; i++ {
		a[i] = ans[i-L]
	}
	return count
}
func check(a []int) int {
	n, count := len(a), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j]*2 {
				count++
			}
		}
	}
	return count
}
func main() {
	maxValue, maxLen := 1000, 1000
	fmt.Println("===测试开始===")
	for i := 0; i < 10000; i++ {
		nums := utils.RandomArray(maxLen, maxValue)
		ans := biggerThanRightTwice(nums)
		if check(nums) != ans {
			for _, num := range nums {
				fmt.Print(num)
				fmt.Print(",")
			}
			fmt.Println()
			fmt.Println(check(nums))
			fmt.Println(ans)
			break
		}
	}
	fmt.Println("===测试结束===")
}

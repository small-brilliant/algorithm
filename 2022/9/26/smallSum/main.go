package main

import (
	"algorithm/utils"
	"fmt"
)

// [1,3,4,2,5]
// 1左边比1小的数：没有
// 3左边比3小的数：1
// 4左边比4小的数：1,3
// 2左边比2小的数：1
// 5左边比5小的数：1,3,4,2
// 所以小和为1+1+3+1+1+3+4+2=16
func smallsum(a []int) int {
	b := make([]int, len(a))
	copy(a, b)
	sum := 0
	var process func(b []int, L, R int)
	process = func(b []int, L, R int) {
		if L >= R {
			return
		}
		M := L + (R-L)/2
		process(b, L, M)
		process(b, M+1, R)
		sum += merge(b, L, M, R)
	}
	process(b, 0, len(b)-1)
	return sum

}

func merge(a []int, L, M, R int) int {
	ans := make([]int, R-L+1)
	p1, p2 := L, M+1
	i, sum := 0, 0
	for ; p1 <= M && p2 <= R; i++ {
		if a[p1] <= a[p2] {
			sum += (R - p2) * a[p1]
			ans[i] = a[p1]
			p1++
		} else {
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
	return sum
}

func check(a []int) int {
	sum := 0
	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				sum += a[j]
			}
		}
	}
	return sum
}

func main() {
	maxValue, maxLen := 100, 100
	fmt.Println("===测试开始===")
	for i := 0; i < 100; i++ {
		nums := utils.RandomArray(maxLen, maxValue)
		ans := smallsum(nums)
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

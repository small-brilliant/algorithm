package main

import (
	"fmt"
	"math/rand"
)

// 等概率生成1-5的随机数
func RandofOneToFive() int {
	return rand.Intn(5) + 1
}

// 通过 RandofOneToFive 获得等概率生成0-1的随机数
func a() int {
	ans := RandofOneToFive()
	for ans == 3 {
		ans = RandofOneToFive()
	}
	if ans < 3 {
		return 1
	} else {
		return 0
	}
}

// 生成0-6的随机数
func b() int {
	ans := 7
	for ans == 7 {
		ans = (a() << 2) + (a() << 1) + (a() << 0)
	}
	return ans
}
func main() {
	for i := 0; i < 100; i++ {
		fmt.Print(b() + 1)
	}
}

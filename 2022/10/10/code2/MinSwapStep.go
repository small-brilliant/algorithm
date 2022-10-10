package main

import (
	"fmt"
	"math/rand"
)

func MinSwapStep(a string) int {
	if len(a) == 0 {
		return 0
	}
	index := 0
	swapCount := 0
	for i := 0; i < len(a); i++ {
		ch := a[i]
		if ch == 'G' {
			swapCount += i - index
			index++
		}
	}
	return swapCount
}
func gernerateS(len int) string {
	b := make([]byte, len)
	for i := 0; i < len; i++ {
		p := rand.Float32()
		if p < 0.5 {
			b[i] = 'G'
		} else {
			b[i] = 'B'
		}
	}
	return string(b)
}
func main() {
	maxL := 0
	s := gernerateS(maxL)
	ans := MinSwapStep(s)

	// if ans != check(s) {
	// 	fmt.Println("Oops")
	// }
	fmt.Println(s)
	fmt.Println(ans)
}

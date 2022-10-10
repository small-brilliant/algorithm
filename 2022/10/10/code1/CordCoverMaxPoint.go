package main

import (
	"algorithm/utils"
	"fmt"
)

func CordCoverMaxPoint(a []int, k int) int {
	n := len(a)
	l, r, ans := 0, 0, 0
	for r < n {
		for a[r]-a[l] > k {
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	maxLen, maxV, k := 1000, 1000, 800
	arr := utils.RandomArray(maxLen, maxV)
	ans := CordCoverMaxPoint(arr, k)
	if ans != check1(arr, k) {
		fmt.Println("Oops")
	}
}
func check1(a []int, k int) int {
	ans := 0
	for i := 0; i < len(a); i++ {
		end := a[i] + k
		j := i
		for j < len(a) && a[j] <= end {
			j++
		}
		ans = max(ans, j-i)
	}
	return ans
}

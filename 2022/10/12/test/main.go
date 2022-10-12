package main

import "fmt"

func main() {
	var T, n, x, y int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&n)
		arr := make([][]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&x, &y)
			t := []int{x, y}
			arr[j] = t
		}
		fmt.Println(arr)
	}
}

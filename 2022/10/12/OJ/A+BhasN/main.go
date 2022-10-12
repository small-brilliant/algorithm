package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		fmt.Println(x + y)
	}
}

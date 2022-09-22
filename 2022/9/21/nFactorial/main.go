package main

import "fmt"

func nfactorial(n int) {
	sum := 0
	fa := 1
	for i := 1; i <= n; i++ {
		fa *= i
		sum += fa
	}
	fmt.Print(sum)
}
func main() {
	nfactorial(3)
}

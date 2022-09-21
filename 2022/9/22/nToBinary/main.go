package main

import "fmt"

func nToBinary(n int) {
	for i := 31; i >= 0; i-- {
		if n>>i&1 == 1 {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
}
func main() {
	nToBinary(1)
}

package main

import "fmt"

func printAllFolds(N int) {
	process(1, N, true)
}
func process(i, N int, down bool) {
	if i > N {
		return
	}
	process(i+1, N, true)
	if down {
		fmt.Print("凹，")
	} else {
		fmt.Print("凸，")
	}
	process(i+1, N, false)
}
func main() {
	printAllFolds(4)
}

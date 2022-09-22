package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if arr[j+1] < arr[j] {
				p := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = p
			}
		}
	}
}
func main() {
	arr := []int{7, 3, 1, 67, 2, 3, 6, 8, 9, 123, 678, 0}
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)
}

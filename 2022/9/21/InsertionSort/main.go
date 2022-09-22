package main

import "fmt"

func InsertionSort(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		for j := i; j >= 1; j-- {
			if arr[j] < arr[j-1] {
				t := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = t
			}
		}
	}
}
func main() {
	arr := []int{7, 3, 1, 1, 1, 1, 67, 2, 3, 6, 8, 9, 123, 678, 0}
	fmt.Println(arr)
	InsertionSort(arr)
	fmt.Println(arr)
}

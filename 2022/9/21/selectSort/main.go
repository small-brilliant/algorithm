package main

import "fmt"

func selectSort(arr []int) {
	var p int
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		p = arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = p
	}
}
func main() {
	arr := []int{7, 3, 1, 67, 2, 3, 6, 8, 9, 123, 678, 0}
	fmt.Println(arr)
	selectSort(arr)
	fmt.Println(arr)
}

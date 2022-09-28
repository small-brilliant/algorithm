package main

import (
	"fmt"
)

// 泛型函数
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v", v)
	}
}

// 泛型切片
type vector[T any] []T

// 泛型map
type M[K string, V any] map[K]V

// 泛型的Chan
type C[T any] chan T

func main() {
	// printSlice([]int64{55, 44, 33, 22, 11})
	// printSlice([]string{"在吗", "2"})

	// v := vector[int]{56, 1232}
	// printSlice(v)
	// v1 := vector[string]{"烤鸡", "烤鸭"}
	// printSlice(v1)

	// m1 := M[string, int]{"key": 1}
	// m2 := M[string, string]{"key": "key"}
	// fmt.Println(m1, m2)
	c1 := make(C[int], 10)
	c2 := make(C[string], 10)
	c1 <- 1
	c1 <- 2
	c2 <- "你好"
	c2 <- "吖"
	fmt.Println(<-c1, <-c2)
}

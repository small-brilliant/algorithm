package main

import (
	"fmt"
	"time"
)

func PrintA(a int) {
	fmt.Println("A", a)
}
func PrintB(a int) {
	fmt.Println("B", a)
}
func main() {
	ch := make(chan int, 100)
	for i := 0; i < 100; i++ {
		ch <- i
		select {
		case t := <-ch:
			go PrintA(t)
		case t := <-ch:
			go PrintB(t)
		}
	}
	time.Sleep(5 * time.Second)
}

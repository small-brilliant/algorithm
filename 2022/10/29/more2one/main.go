package main

import (
	"fmt"
	"time"
)

var c = make(chan int, 3)

func main() {
	go consumeA()
	go consumeB()
	time.Sleep(1 * time.Second)
	for i := 0; i < 100; i++ {
		c <- i
		time.Sleep(1 * time.Millisecond)
	}
}
func consumeA() {
	for {
		fmt.Println("A", <-c)
	}
}
func consumeB() {
	for {
		fmt.Println("B", <-c)
		time.Sleep(1 * time.Millisecond)
	}
}

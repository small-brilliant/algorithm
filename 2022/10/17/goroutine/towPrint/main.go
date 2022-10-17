package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	ch := make(chan struct{})
	go print1(ch)
	go print2(ch)
	wg.Wait()
}
func print1(ch chan struct{}) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- struct{}{}
		if i&1 == 0 {
			fmt.Println("goroutine1:", i)
		}
	}
}
func print2(ch chan struct{}) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-ch
		if i&1 == 1 {
			fmt.Println("goroutine2:", i)
		}
	}
}

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan struct{})
	wg.Add(2)
	go print1(ch)
	go print2(ch)
	wg.Wait()
}
func print1(ch chan struct{}) {
	for i := 1; i <= 100; i++ {
		if i&1 == 0 {
			ch <- struct{}{}
			fmt.Println("线程1", i)
		}
	}
	defer wg.Done()
}
func print2(ch chan struct{}) {
	for i := 1; i <= 100; i++ {
		if i&1 == 1 {
			<-ch
			fmt.Println("线程2", i)
		}
	}
	defer wg.Done()
}

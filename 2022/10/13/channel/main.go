package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// ch1 := make(chan struct{}, 1)
	// ch2 := make(chan struct{}, 1)
	// ch3 := make(chan struct{}, 1)
	// ch1 <- struct{}{}
	// wg.Add(3)
	// start := time.Now().Unix()
	// go print("goroutine1", ch1, ch2)
	// go print("goroutine2", ch2, ch3)
	// go print("goroutine3", ch3, ch1)
	// wg.Wait()
	// end := time.Now().Unix()
	// fmt.Printf("duration:%d", end-start)
	wg.Add(1)
	ch := make(chan struct{})
	go print2("goroutine1", ch)
	ch <- struct{}{}
	wg.Wait()

}
func print2(goroutine string, inch chan struct{}) {
	select {
	case <-inch:
		fmt.Printf("%s\n", goroutine)
		time.Sleep(1 * time.Second)
	}
	wg.Done()

}
func print(goroutine string, inch chan struct{}, outch chan struct{}) {
	select {
	case <-inch:
		fmt.Printf("%s\n", goroutine)
		time.Sleep(1 * time.Second)
		outch <- struct{}{}
	}
	wg.Done()

}

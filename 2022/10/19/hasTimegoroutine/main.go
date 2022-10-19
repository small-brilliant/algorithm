package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan struct{})
var ch2 = make(chan struct{})

func process1() {
	fmt.Println("process1 working")
	time.Sleep(1 * time.Second)
	ch1 <- struct{}{}
}
func process2() {
	<-ch1
	fmt.Println("process2 working")
	time.Sleep(3 * time.Second)
	ch2 <- struct{}{}
}
func main() {
	tic := time.NewTicker(2 * time.Second)
	go func() {
		go process1()
		select {
		case <-tic.C:
			fmt.Println("f1 timeout")
		case <-ch1:
			fmt.Println("f1 done")
		}
	}()
	go func() {
		go process2()
		select {
		case <-tic.C:
			fmt.Println("f2 timeout")
		case <-ch2:
			fmt.Println("f2 done")
		}
	}()
	time.Sleep(5 * time.Second)
}

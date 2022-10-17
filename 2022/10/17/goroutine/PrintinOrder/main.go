package main

import (
	"os"
	"runtime/trace"
	"sync"
)

var wg sync.WaitGroup

func main() {
	f, _ := os.Create("myTrace.dat")
	defer f.Close()
	_ = trace.Start(f)
	defer trace.Stop()
	wg.Add(3)
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go first(ch1)
	go second(ch1, ch2)
	go third(ch2)
	wg.Wait()
}
func first(ch chan struct{}) {
	defer wg.Done()
	println("first")
	ch <- struct{}{}
}
func second(ch1 chan struct{}, ch2 chan struct{}) {
	defer wg.Done()
	<-ch1
	println("second")
	ch2 <- struct{}{}
}
func third(ch chan struct{}) {
	defer wg.Done()
	<-ch
	println("trird")
}

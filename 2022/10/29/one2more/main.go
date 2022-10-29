package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func run(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("stop!")
			wg.Done()
		case <-time.After(time.Second):
			fmt.Println("hello")
		}
	}
}
func main() {
	stop := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go run(stop)
	}
	time.Sleep(3 * time.Second)
	close(stop)
	wg.Wait()
}

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	ch := make(chan struct{}, 2)
	for {
		select {
		case <-t.C:
			for i := 0; i < 2; i++ {
				ch <- struct{}{}
			}
			for i := 0; i <= 10000; i++ {
				go func(i int) {
					<-ch
					fmt.Println("线程", i)
				}(i)
			}
		}
	}
}

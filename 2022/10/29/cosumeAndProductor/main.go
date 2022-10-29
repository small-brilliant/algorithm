package main

import (
	"fmt"
	"time"
)

func product(ch chan int, stopCh chan struct{}) {
	for j := 0; j < 100; j++ {
		ch <- j
		time.Sleep(time.Millisecond)
	}
	close(stopCh)
}
func comsumeA(ch chan int, stopCh chan struct{}) {
	for {
		select {
		case c := <-ch:
			fmt.Println("comsumeA", c)
		case <-stopCh:
			goto end
		}
	}
end:
}
func comsumeB(ch chan int, stopCh chan struct{}) {
	for {
		select {
		case c := <-ch:
			fmt.Println("comsumeB", c)
		case <-stopCh:
			goto end
		}
	}
end:
}

func main() {

	ch := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	stopCh := make(chan struct{})

	go product(ch, stopCh)
	go comsumeA(ch1, stopCh)
	go comsumeB(ch2, stopCh)
	for {
		select {
		case c := <-ch:
			ch1 <- c
		case c := <-ch:
			ch2 <- c
		case <-stopCh:
			goto end
		}
	}
end:
}

// func PrintA(a chan int) {
// 	for {
// 		<-a
// 		fmt.Println("A", <-a)
// 	}
// }
// func PrintB(a chan int) {
// 	for {
// 		<-a
// 		fmt.Println("A", <-a)
// 	}
// }
// func product(ch1, ch2 chan int) {
// 	for i := 1; i < 100; i++ {
// 		select {
// 		case <-ch1:
// 			fmt.Println("yes")
// 			ch2 <- i
// 		case <-ch2:
// 			ch1 <- i
// 		}
// 	}
// }
// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int, 1)
// 	ch2 <- 0
// 	go PrintA(ch1)
// 	go PrintB(ch2)
// 	go product(ch1, ch2)

// 	time.Sleep(3 * time.Second)
// }

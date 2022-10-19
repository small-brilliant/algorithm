package main

import (
	"fmt"
	"time"
)

func Cat(ch1, ch2 chan struct{}) {
	<-ch1
	fmt.Println("1. Cat")
	ch2 <- struct{}{}
}
func Fish(ch1, ch2 chan struct{}) {
	<-ch1
	fmt.Println("2. Fish")
	ch2 <- struct{}{}
}
func Dog(ch1, ch2 chan struct{}) {
	<-ch1
	fmt.Println("3. Dog")
	ch2 <- struct{}{}
}
func main() {
	chCat := make(chan struct{}, 1)
	chFish := make(chan struct{})
	chDog := make(chan struct{})
	for i := 0; i < 100; i++ {
		go Cat(chCat, chFish)
		go Fish(chFish, chDog)
		go Dog(chDog, chCat)
	}
	chCat <- struct{}{}
	time.Sleep(10 * time.Second)
}

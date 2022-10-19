package main

import (
	"fmt"
	"time"
)

var word = make(chan struct{})
var num = make(chan struct{})

func printNums() {
	for i := 0; i < 10; i++ {
		<-word
		fmt.Println("1")
		num <- struct{}{}
	}
}
func printWords() {
	for i := 0; i < 10; i++ {
		<-num
		fmt.Println("A")
		word <- struct{}{}
	}
}
func main() {
	go printNums()
	go printWords()
	word <- struct{}{}
	time.Sleep(3 * time.Second)
}

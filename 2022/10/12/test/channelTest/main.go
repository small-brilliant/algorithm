package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	ch := make(chan string, 4)
	fmt.Println(ch, unsafe.Sizeof(ch))
	go sendTask(ch)
	go receiveTask(ch)
	time.Sleep(1 * time.Second)
}

func receiveTask(ch chan string) {
	for {
		task := <-ch
		fmt.Println("received", task)
	}
}

func sendTask(ch chan string) {
	taskList := []string{"this", "is", "a", "demo", "!", "!"}
	for _, task := range taskList {
		ch <- task
	}
}

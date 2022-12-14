package main

import (
	"fmt"
)

func main() {
	defer_call()
}

var x interface{} = nil

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	defer func() {
		recover()
	}()
	panic("触发异常")
}

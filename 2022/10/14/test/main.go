package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 2; i++ {
		// 当前协程等待运行，直到其他协程执行完毕后自动恢复，只有单核心时有效。
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main() {
	wg.Add(1)
	fmt.Println(runtime.GOMAXPROCS(1))
	go func() {
		say("world")
		defer wg.Done()
	}()
	say("hello")
	go func() {
		for i := 0; i < 10; i++ {
			//终止所在的线程
			fmt.Println("AAAAA")
			runtime.Goexit()
		}
		fmt.Println("C")
		defer wg.Done()
	}()
	wg.Wait()
}

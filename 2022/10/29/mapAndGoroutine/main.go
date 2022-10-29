package main

import (
	"fmt"
	"sync"
)

var maps sync.Map
var wg sync.WaitGroup

func main() {
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go caculate(i)
	}
	wg.Wait()

	maps.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
func caculate(n int) {
	defer wg.Done()
	cur := 1
	for i := 2; i <= n; i++ {
		cur += i
	}
	maps.Store(n, cur)
}

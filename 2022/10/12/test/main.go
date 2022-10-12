package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 测试
	// var T, n, x, y int
	// fmt.Scan(&T)
	// for i := 0; i < T; i++ {
	// 	fmt.Scan(&n)
	// 	arr := make([][]int, n)
	// 	for j := 0; j < n; j++ {
	// 		fmt.Scan(&x, &y)
	// 		t := []int{x, y}
	// 		arr[j] = t
	// 	}
	// 	fmt.Println(arr)
	// }

	// map的线程安全实现
	// 方法一
	var lock sync.RWMutex
	s := map[int]int{}
	for i := 0; i < 100; i++ {
		go func(i int) {
			lock.Lock()
			s[i] = i
			lock.Unlock()
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Printf("map第%d个元素值位%d", i, s[i])
			lock.RUnlock()
		}(i)
	}
	// 方法二
	var m sync.Map
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Store(i, i)
		}(i)
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			v, ok := m.Load(i)
			fmt.Printf("Load:%v,%v", v, ok)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

package generatewater

import "sync"

var wg sync.WaitGroup

func GenerateWater(n int) (ohh, hoh, hho int) {
	if n%3 != 0 {
		return 0, 0, 0
	}
	n = n / 3
	wg.Add(3)
	ox := make(chan struct{})
	hy1 := make(chan struct{})
	hy2 := make(chan struct{})
	go oxygen(n, ox)
	go hydrogen1(n, hy1)
	go hydrogen2(n, hy2)
	for i := 0; i < n; i++ {
		select {
		case <-ox:
			ohh++
			_, _ = <-hy2, <-hy1
		case <-hy1:
			hoh++
			_, _ = <-ox, <-hy2
		case <-hy2:
			hho++
			_, _ = <-hy1, <-ox
		}
	}
	wg.Wait()
	return
}
func oxygen(n int, ch chan struct{}) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
}
func hydrogen1(n int, ch chan struct{}) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
}
func hydrogen2(n int, ch chan struct{}) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
}

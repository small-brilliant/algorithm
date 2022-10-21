package main

import "fmt"

const (
	x = 100
	_
	y = 100 * x
	z = "zz"
	k
	p = iota
)

func main() {
	fmt.Println(x, y, z, k, p)
}

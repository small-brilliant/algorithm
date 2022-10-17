package main

import "fmt"

func modifyInt(i *int) {
	*i = 10
}
func main() {
	i := 9
	modifyInt(&i)
	fmt.Println(i)
}

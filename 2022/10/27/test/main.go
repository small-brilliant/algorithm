package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	s = strings.Replace(s, " ", "%20", -1)
	return s
}
func main() {
	fmt.Println(replaceSpace("we are happy"))
}
